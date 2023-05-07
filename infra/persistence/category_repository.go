package persistence

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/domain/repository"
	"github.com/nokin-all-of-career/career-web-backend/infra"
)

type categoryPersistence struct {
	DB *sql.DB
}

// NewCategoryPersistence : create category persistence
func NewCategoryPersistence(DB *sql.DB) repository.CategoryRepository {
	return &categoryPersistence{
		DB: DB,
	}
}

// InitCategory : get whole category from category table and put in map
func InitCategory() error {
	rows, err := ReadAllCategories()
	if err != nil {
		return errors.Wrap(err, "categoryPersistence.InitCategory()")
	}

	return InsertCategoryMap(rows)
}

// ReadAllCategories : get all category
func ReadAllCategories() (*sql.Rows, error) {
	rows, err := infra.DB.Query("SELECT * FROM category")
	if err != nil {
		return nil, errors.Wrap(err, "categoryPersistence.ReadAllCategories()")
	}

	return rows, nil
}

// InsertCategoryMap : insert in map
func InsertCategoryMap(rows *sql.Rows) error {
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.CategoryID, &category.Name, &category.Goal)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil
			}
			return errors.Wrap(err, "categoryPersistence.InsertCategoryMap()")
		}
		domain.CategoryMap[category.CategoryID] = category
	}

	return nil
}

func (cp categoryPersistence) SelectCategories() (*[]domain.Category, error) {
	var categories []domain.Category
	for _, v := range domain.CategoryMap {
		categories = append(categories, v)
	}
	return &categories, nil
}

func (cp categoryPersistence) SelectByPrimaryKey(categoryID string) (*domain.Category, error) {
	category := domain.CategoryMap[categoryID]
	return &category, nil
}
