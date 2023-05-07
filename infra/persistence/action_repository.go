package persistence

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/domain/repository"
	"github.com/nokin-all-of-career/career-web-backend/infra"
)

type actionPersistence struct {
	DB *sql.DB
}

// NewActionPersistence : create action persistence
func NewActionPersistence(DB *sql.DB) repository.ActionRepository {
	return &actionPersistence{
		DB: DB,
	}
}

var actionMap = map[string]domain.Action{}

// InitAction : get whole actions from action tabels and put in map
func InitAction() error {
	rows, err := ReadAllActions()
	if err != nil {
		return errors.Wrap(err, "actionPersistence.InitAction()")
	}
	return InsertActionMap(rows)
}

// InsertActionMap : insert in map
func InsertActionMap(rows *sql.Rows) error {
	for rows.Next() {
		action := domain.Action{}
		err := rows.Scan(&action.ActionID, &action.CategoryID, &action.Title, &action.Content, &action.StandardTime, &action.ActionType, &action.URL, &action.After)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil
			}
			return errors.Wrap(err, "actionPersistence.InsertActionMap()")
		}
		actionMap[action.ActionID] = action
	}

	return nil
}

// ReadAllActions : get all actions
func ReadAllActions() (*sql.Rows, error) {
	rows, err := infra.DB.Query("SELECT * FROM action")
	if err != nil {
		return nil, errors.Wrap(err, "actionPersistence.ReadAllActions()")
	}

	return rows, nil
}

// SelectByActiontype : get tutorial action (return only slice of actionID)
func (ap actionPersistence) SelectByActiontype(actionType int) ([]string, error) {
	var actionIDs []string
	for _, v := range actionMap {
		if v.ActionType == actionType {
			actionIDs = append(actionIDs, v.ActionID)
		}
	}
	return actionIDs, nil
}

func (ap actionPersistence) SelectByCategoryID(categoryID string) ([]domain.Action, error) {
	actions := []domain.Action{}
	for _, v := range actionMap {
		if v.CategoryID == categoryID {
			actions = append(actions, v)
		}
	}
	return actions, nil
}

func (ap actionPersistence) SelectByPrimaryKey(actionID string) (domain.Action, error) {
	action := actionMap[actionID]
	return action, nil
}

func (ap actionPersistence) SelectByPrimaryKeys(actionIDs []interface{}) (map[string]domain.Action, error) {
	actions := map[string]domain.Action{}
	for _, actionID := range actionIDs {
		actions[actionID.(string)] = actionMap[actionID.(string)]
	}
	return actions, nil
}
