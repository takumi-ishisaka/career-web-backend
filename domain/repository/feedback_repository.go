package repository

import (
	"database/sql"

	"github.com/nokin-all-of-career/career-web-backend/domain"
)

// FeedBackRepository : feedback repository
type FeedBackRepository interface {
	InsertFeedBack(tx *sql.Tx, feedbackID, userActionID, comment string) error
	SelectByPrimaryKey(userActionID string) (*[]domain.FeedBack, error)
	SelectByPrimaryKeys(userActionIDs []string) (*[]domain.FeedBack, error)
	SendEmailToUser(from, to string, userActionName, feedback string) error
}
