package repository

import (
	"database/sql"

	"github.com/nokin-all-of-career/career-web-backend/domain"
)

// UserActionRepository : userAction repository
type UserActionRepository interface {
	SelectByUserIDAndStatus(userID string, status int) ([]*domain.UserAction, error)
	SelectByPrimaryKey(userActionID string) (*domain.UserAction, error)
	DeleteUserAction(userID, actionID string) error
	InsertUserAction(userActionID, userID, actionID string, status int) error
	UpdateUserAction(userID, actionID string, status int) error

	SelectReflectionByPrimaryKey(userID, actionID string) (*domain.Reflection, error)
	InsertReflection(userID, actionID, keep, problem, try string, evaluateValue int) error
	GetDoneActionCountByUserID(userID string) (float32, error)
	GetDoneActionCountSum() (int, int, error)

	GetRecommendationByActionID(actionID string) (*domain.Recommendation, error)
	GetRecommendationsByActionID(actionIDs []interface{}) ([]domain.Recommendation, error)
	GetRecommendationsMapByActionID(actionIDs []interface{}) (map[string]domain.Recommendation, error)

	UpdateUserActionByUserActionID(tx *sql.Tx, userActionID string, status int) error
	SelectUserActionsOfAllUsers(status int) ([]*domain.UserAction, error)

	SendEmailToOperation(from, to string, userActionName, userName, do, reflection, nextAction string) error
}
