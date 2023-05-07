package repository

import (
	"github.com/nokin-all-of-career/career-web-backend/domain"
)

// ActionRepository : action repository
type ActionRepository interface {
	SelectByActiontype(actionType int) ([]string, error)
	SelectByCategoryID(categaryID string) ([]domain.Action, error)
	SelectByPrimaryKey(actionID string) (domain.Action, error)
	SelectByPrimaryKeys(actionIDs []interface{}) (map[string]domain.Action, error)
}
