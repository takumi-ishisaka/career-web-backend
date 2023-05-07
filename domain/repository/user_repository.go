package repository

import (
	"github.com/nokin-all-of-career/career-web-backend/domain"
)

// UserRepository : user repository
type UserRepository interface {
	Insert(userID, email, password string) error
	SelectByEmail(email string) (*domain.User, error)
	SelectByPrimaryKey(userID string) (*domain.User, error)
	Update(userID, email string) error
	SendValidationEmail(from, to string, url string) error
}
