package repository

import "github.com/nokin-all-of-career/career-web-backend/domain"

// CategoryRepository : category repository
type CategoryRepository interface {
	SelectCategories() (*[]domain.Category, error)
	SelectByPrimaryKey(categoryID string) (*domain.Category, error)
}
