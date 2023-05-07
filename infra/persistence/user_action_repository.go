package persistence

import (
	"database/sql"
	"net/smtp"

	"github.com/pkg/errors"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/domain/repository"
	"github.com/nokin-all-of-career/career-web-backend/infra"
)

type userActionPersistence struct {
	DB *sql.DB
}

// NewUserActionPersistence : create userAction persistence
func NewUserActionPersistence(DB *sql.DB) repository.UserActionRepository {
	return &userActionPersistence{
		DB: DB,
	}
}

// SelectByUserIDAndStatus : get user todo_action by userID and status
func (uap userActionPersistence) SelectByUserIDAndStatus(userID string, status int) ([]*domain.UserAction, error) {
	rows, err := uap.DB.Query("SELECT user_action_id,user_id, action_id, status, updated_at FROM user_action WHERE user_id=? AND status=? ORDER BY updated_at ASC", userID, status)
	if err != nil {
		return nil, errors.Wrap(err, "userActionPersisitence.SelectByUserIDAndStatus()")
	}

	userActions := []*domain.UserAction{}
	for rows.Next() {
		userAction := domain.UserAction{}
		err := rows.Scan(&userAction.UserActionID, &userAction.UserID, &userAction.ActionID, &userAction.Status, &userAction.UpdatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				return userActions, nil
			}
			return nil, errors.Wrap(err, "userActionPersistence.SelectByUserIDAndStatus()")
		}
		userActions = append(userActions, &userAction)
	}

	return userActions, nil
}

// DeleteUserAction : delete user action when to notselect from todo
func (uap userActionPersistence) DeleteUserAction(userID, actionID string) error {
	stmt, err := uap.DB.Prepare("DELETE FROM user_action WHERE user_id=? AND action_id=?")
	if err != nil {
		return errors.Wrap(err, "userActionPersistence.DeleteUserAction()")
	}

	_, err = stmt.Exec(userID, actionID)
	if err != nil {
		return errors.Wrap(err, "userActionPersisitence.DeleteUserAction()")
	}

	return nil
}

// InsertUserAction : upsert userAction to todo from notselect and to notselect from todo
func (uap userActionPersistence) InsertUserAction(userActionID, userID, actionID string, status int) error {
	stmt, err := uap.DB.Prepare("INSERT INTO user_action(user_action_id, user_id, action_id, status) VALUES(?,?,?,?)")
	if err != nil {
		return errors.Wrap(err, "userActionPersistence.InsertUserAction()")
	}

	_, err = stmt.Exec(userActionID, userID, actionID, status)
	if err != nil {
		return errors.Wrap(err, "userActionPersisitence.InsertUserAction()")
	}

	return nil
}

func (uap userActionPersistence) UpdateUserAction(userID, actionID string, status int) error {
	stmt, err := uap.DB.Prepare("UPDATE user_action SET status=? WHERE user_id=? AND action_id=?")
	if err != nil {
		return errors.Wrap(err, "userActionPersistence.UpdateUserAction()")
	}

	_, err = stmt.Exec(status, userID, actionID)
	if err != nil {
		return errors.Wrap(err, "userActionPersisitence.UpdateUserAction()")
	}

	return nil
}

func (uap userActionPersistence) SelectReflectionByPrimaryKey(userID, actionID string) (*domain.Reflection, error) {
	row := uap.DB.QueryRow("SELECT do, reflection, next_action, evaluate_value FROM user_action WHERE user_id=? AND action_id=?", userID, actionID)

	var reflection domain.Reflection
	err := row.Scan(&reflection.Do, &reflection.Reflection, &reflection.NextAction, &reflection.EvaluateValue)
	if err != nil {
		return nil, errors.Wrap(err, "userActionReflectionPersistence.SelectReflectionByPrimaryKey()")
	}

	return &reflection, nil
}

func (uap userActionPersistence) SelectByPrimaryKey(userActionID string) (*domain.UserAction, error) {
	row := uap.DB.QueryRow("SELECT * FROM user_action WHERE user_action_id=?", userActionID)

	var userAction domain.UserAction
	err := row.Scan(&userAction.UserActionID, &userAction.UserID, &userAction.ActionID, &userAction.Status, &userAction.UpdatedAt, &userAction.Do, &userAction.Reflection, &userAction.NextAction, &userAction.EvaluateValue)
	if err != nil {
		return nil, errors.Wrap(err, "userActionPersistence.SelectByPrimaryKey()")
	}

	return &userAction, nil
}

func (uap userActionPersistence) InsertReflection(userID, actionID, do, reflection, nextAction string, evaluateValue int) error {
	stmt, err := uap.DB.Prepare("UPDATE user_action SET do=?, reflection=?, next_action=?, evaluate_value=? WHERE user_id=? AND action_id=?")
	if err != nil {
		return errors.Wrap(err, "userActionPersistence.InsertReflection()")
	}

	_, err = stmt.Exec(do, reflection, nextAction, evaluateValue, userID, actionID)
	if err != nil {
		return errors.Wrap(err, "userActionPersistence.InsertReflection()")
	}

	return nil
}

// GetDoneActionCountByUserID : get user doneAction count by userID
func (uap userActionPersistence) GetDoneActionCountByUserID(userID string) (float32, error) {
	var doneCount float32
	doneCount = 0.0
	row := uap.DB.QueryRow("SELECT COUNT(*) FROM user_action WHERE user_id=? AND status=2", userID)
	err := row.Scan(&doneCount)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, errors.Wrap(err, "userActionPersistence.GetDoneActionCountByPrimaryKey")
		}
	}

	return doneCount, nil
}

// GetDoneActionCountSum : get user doneAction count sum
func (uap userActionPersistence) GetDoneActionCountSum() (int, int, error) {
	rows, err := uap.DB.Query("SELECT COUNT(*) FROM user_action WHERE status=2 GROUP BY user_id")
	if err != nil {
		return 0, 0, errors.Wrap(err, "userActionPersistence,GetDoneActionCountSum()")
	}
	countSum := 0
	users := 0
	for rows.Next() {
		count := 0
		users++
		err := rows.Scan(&count)
		if err != nil {
			if err == sql.ErrNoRows {
				return countSum, users, nil
			}
			return 0, 0, errors.Wrap(err, "userActionPersistence.GetDoneActionCountSum()")
		}
		countSum += count
	}

	return countSum, users, nil
}

func (uap userActionPersistence) GetRecommendationByActionID(actionID string) (*domain.Recommendation, error) {
	row := uap.DB.QueryRow("SELECT MAX(action_id), (SUM(evaluate_value)/COUNT(user_id)) AS recommendation_value, COUNT(user_id) AS done_user_num FROM user_action WHERE action_id=? AND status = 2;", actionID)
	recommendation, err := convertToRecommendation(actionID, row)
	if err != nil {
		return nil, errors.Wrap(err, "userActionPersistence.GetRecommendationByActionID()")
	}

	return recommendation, nil
}

func (uap userActionPersistence) GetRecommendationsByActionID(actionIDs []interface{}) ([]domain.Recommendation, error) {
	// generate query
	if len(actionIDs) == 0 {
		return nil, nil
	}
	query := "SELECT action_id, (SUM(evaluate_value)/COUNT(user_id)) AS recommendation_value, COUNT(user_id) AS done_user_num FROM user_action WHERE status = 2 AND action_id IN(?"
	for i := 1; i < len(actionIDs); i++ {
		query += ",?"
	}
	query += ") GROUP BY action_id;"
	rows, err := uap.DB.Query(query, actionIDs...)
	if err != nil {
		return nil, errors.Wrap(err, "userActionPersistence.GetRecommendationsByActionID")
	}

	var recommendations []domain.Recommendation
	count := 0
	for rows.Next() {
		var recommendation domain.Recommendation
		var nullString sql.NullString
		var nullInt sql.NullInt64
		var nullFloat sql.NullFloat64

		// check and escape null
		err := rows.Scan(&nullString, &nullFloat, &nullInt)
		if err != nil {
			return nil, errors.Wrap(err, "userActionPersistence.GetRecommendationsByActionID")
		}

		// default value of RecommendationValue is 3
		if !nullFloat.Valid {
			nullFloat.Float64 = 3
		}

		recommendation = domain.Recommendation{
			ActionID:            nullString.String,
			RecommendationValue: float32(nullFloat.Float64),
			DoneUserNum:         int(nullInt.Int64),
		}

		recommendations = append(recommendations, recommendation)
		count++
	}
	return recommendations, nil
}

func (uap userActionPersistence) GetRecommendationsMapByActionID(actionIDs []interface{}) (map[string]domain.Recommendation, error) {
	// generate query
	if len(actionIDs) == 0 {
		return nil, nil
	}
	query := "SELECT MAX(action_id),(SUM(evaluate_value)/COUNT(user_id)) AS recommendation_value,COUNT(user_id) AS done_user_num FROM user_action WHERE status = 2 AND action_id IN(?"
	for i := 1; i < len(actionIDs); i++ {
		query += ",?"
	}
	query += ");"
	rows, err := uap.DB.Query(query, actionIDs...)
	if err != nil {
		return nil, errors.Wrap(err, "userAction.SelectByPrimaryKeys()")
	}

	recommendations := map[string]domain.Recommendation{}
	var nullString sql.NullString
	var nullInt sql.NullInt64
	var nullFloat sql.NullFloat64
	count := 0
	for rows.Next() {
		var recommendation domain.Recommendation
		err := rows.Scan(&recommendation.ActionID, &recommendation.RecommendationValue, &recommendation.DoneUserNum)
		if err != nil {
			_ = rows.Scan(&nullString, &nullFloat, &nullInt)
			if nullFloat.Valid || nullInt.Valid || nullString.Valid {
				recommendation = domain.Recommendation{
					ActionID:            actionIDs[count].(string),
					RecommendationValue: 5,
					DoneUserNum:         0,
				}
				recommendations[recommendation.ActionID] = recommendation
				break
			}
			recommendations[recommendation.ActionID] = recommendation
			count++
		}
	}
	return recommendations, nil
}

func convertToRecommendation(actionID string, row *sql.Row) (*domain.Recommendation, error) {
	recommendation := domain.Recommendation{}
	var nullString sql.NullString
	var nullInt sql.NullInt64
	var nullFloat sql.NullFloat64
	err := row.Scan(&recommendation.ActionID, &recommendation.RecommendationValue, &recommendation.DoneUserNum)
	if err != nil {
		_ = row.Scan(&nullString, &nullFloat, &nullInt)
		if nullFloat.Valid || nullInt.Valid || nullString.String == "" {
			recommendation = domain.Recommendation{
				ActionID:            actionID,
				RecommendationValue: 5,
				DoneUserNum:         0,
			}
			return &recommendation, nil
		}
		return nil, errors.Wrap(err, "userActionPersistence.convertToRecommendation()")
	}

	return &recommendation, nil
}

func (uap userActionPersistence) UpdateUserActionByUserActionID(tx *sql.Tx, userActionID string, status int) error {
	stmt, err := tx.Prepare("UPDATE user_action SET status=? WHERE user_action_id=?;")
	if err != nil {
		return errors.Wrap(err, "userActionPersistence.UpdateUserActionByUserActionID()")
	}

	_, err = stmt.Exec(status, userActionID)
	if err != nil {
		return errors.Wrap(err, "userActionPersisitence.UpdateUserActionByUserActionID()")
	}

	return nil
}

func (uap userActionPersistence) SelectUserActionsOfAllUsers(status int) ([]*domain.UserAction, error) {
	rows, err := uap.DB.Query("SELECT * FROM user_action WHERE status=? ORDER BY updated_at DESC;", status)
	if err != nil {
		return nil, errors.Wrap(err, "userActionPersistence.UpdateUserActionByUserActionID()")
	}

	userActions := []*domain.UserAction{}
	for rows.Next() {
		userAction := domain.UserAction{}
		err := rows.Scan(&userAction.UserActionID, &userAction.UserID, &userAction.ActionID, &userAction.Status, &userAction.UpdatedAt, &userAction.Do, &userAction.Reflection, &userAction.NextAction, &userAction.EvaluateValue)
		if err != nil {
			if err == sql.ErrNoRows {
				return userActions, nil
			}
			return nil, errors.Wrap(err, "userActionPersistence.SelectByUserIDAndStatus()")
		}
		userActions = append(userActions, &userAction)
	}

	return userActions, nil
}

func (uap userActionPersistence) SendEmailToOperation(from, to string, userActionName, userName, do, reflection, nextAction string) error {
	msg := []byte("" +
		"From: オルキャリ運営 <" + from + ">\r\n" +
		"To: オルキャリ運営 <" + to + ">\r\n" +
		"Subject: FB依頼が届きました。\r\n" +
		"\r\n" +
		userName + "から\r\n" +
		userActionName + "のFB依頼が届きました。\r\n" +
		"【やったこと】\r\n" +
		do + "\r\n" +
		"【反省】\r\n" +
		reflection + "\r\n" +
		"【次のアクション】\r\n" +
		nextAction + "\r\n" +
		"")
	err := smtp.SendMail("smtp.gmail.com:587", infra.Auth, from, []string{to}, msg)
	if err != nil {
		return errors.Wrap(err, "feedbackPersistence.SendEmailToUser()")
	}
	return nil
}
