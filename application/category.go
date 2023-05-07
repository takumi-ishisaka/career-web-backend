package application

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/domain/repository"
)

// CategoryApplication : category application
type CategoryApplication interface {
	CategoryList(echo.Context) error
}

type categoryApplication struct {
	categoryRepository repository.CategoryRepository
}

// NewCategoryApplication : create application about categoru
func NewCategoryApplication(cr repository.CategoryRepository) CategoryApplication {
	return &categoryApplication{
		categoryRepository: cr,
	}
}

func (ca categoryApplication) CategoryList(c echo.Context) error {
	categories, err := ca.categoryRepository.SelectCategories()
	if err != nil {
		return errors.Wrap(err, "categoryApplication.CategoryList()")
	}

	return c.JSON(http.StatusOK, CategoriesResponse{Categories: categories})
}

// CategoriesResponse : categories response
type CategoriesResponse struct {
	Categories *[]domain.Category `json:"categories"`
}
