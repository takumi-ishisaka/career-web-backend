package persistence

import (
	"database/sql"
	"net/smtp"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/domain/repository"
	"github.com/nokin-all-of-career/career-web-backend/infra"
	"github.com/pkg/errors"
)

type feedbackPersistence struct {
	DB *sql.DB
}

// NewFeedBackPersistence : create feedback persistence
func NewFeedBackPersistence(DB *sql.DB) repository.FeedBackRepository {
	return &feedbackPersistence{
		DB: DB,
	}
}

func (fp feedbackPersistence) InsertFeedBack(tx *sql.Tx, feedbackID, userActionID, comment string) error {
	stmt, err := tx.Prepare("INSERT INTO feedback(feedback_id,user_action_id, comment) VALUES(?,?, ?)")
	if err != nil {
		return errors.Wrap(err, "feedbackPersistence.InsertFeedBack()")
	}
	_, err = stmt.Exec(feedbackID, userActionID, comment)
	if err != nil {
		return errors.Wrap(err, "feedbackPersistence.InsertFeedBack()")
	}

	return nil
}

func (fp feedbackPersistence) SelectByPrimaryKey(userActionID string) (*[]domain.FeedBack, error) {
	rows, err := fp.DB.Query("SELECT * FROM feedback WHERE user_action_id=?", userActionID)
	if err != nil {
		return nil, errors.Wrap(err, "feedbackPersistence.SelectFeedBackPrimaryKey()")
	}
	var feedback domain.FeedBack
	var feedbacks []domain.FeedBack
	for rows.Next() {
		err := rows.Scan(&feedback.FeedBackID, &feedback.UserActionID, &feedback.Comment, &feedback.UpdatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "feedbackPersistence.SelectFeedBackPrimaryKey()")
		}
		feedbacks = append(feedbacks, feedback)
	}

	return &feedbacks, nil
}

func (fp feedbackPersistence) SelectByPrimaryKeys(userActionIDs []string) (*[]domain.FeedBack, error) {
	query := "SELECT * FROM feedback WHERE user_action_id IN(?"
	for i := 1; i < len(userActionIDs); i++ {
		query += ",?"
	}
	query += ");"
	rows, err := fp.DB.Query(query, userActionIDs)
	var feedback domain.FeedBack
	var feedbacks []domain.FeedBack
	for rows.Next() {
		err = rows.Scan(&feedback.UserActionID, &feedback.Comment, &feedback.UpdatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "feedbackPersistence.SelectFeedBacksPrimaryKeys()")
		}
		feedbacks = append(feedbacks, feedback)
	}
	return &feedbacks, nil
}

func (fp feedbackPersistence) SendEmailToUser(from, to string, userActionName, feedback string) error {

	msg := []byte("" +
		"From: オルキャリ運営 <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: メンターからFBが届きました\r\n" +
		"\r\n" +
		userActionName + "にメンターからFBが届きました。\r\n" +
		"「内容」\r\n" +
		feedback + "\r\n" +
		"")

	err := smtp.SendMail("smtp.gmail.com:587", infra.Auth, from, []string{to}, msg)
	if err != nil {
		return errors.Wrap(err, "feedbackPersistence.SendEmailToUser()")
	}

	return nil
}
