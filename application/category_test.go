package application

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/infra"
	"github.com/nokin-all-of-career/career-web-backend/infra/persistence"
	"github.com/stretchr/testify/assert"
)

//テスト用に入っている値
// 	domain.Category{"category_001", "自己分析"}
// 	domain.Category{"category_002", "企業分析"}

func categoryInitializaion() (*echo.Echo, CategoryApplication) {
	p := persistence.NewCategoryPersistence(infra.DB)
	a := NewCategoryApplication(p)
	e := echo.New()
	return e, a
}

func TestCategoryList(t *testing.T) {
	// Setup
	// return *echo.Echo and CategoryApplication
	e, a := categoryInitializaion()
	var categories []domain.Category
	categories = append(categories,
		domain.Category{"category_001", "自己分析"},
		domain.Category{"category_002", "企業分析"},
	)
	var categoryListTC = []CategoryListTC{
		{"auth", categories, 200},
		{"auth", nil, 500},
	}

	for _, v := range categoryListTC {
		// c is echo.Context
		// rec is *http.ResponseRecoder
		c, rec := CreateCategoryListRequest(e, v.JWT)

		if assert.NoError(t, a.CategoryList(c)) {
			// normal system testcase
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, v.Categories, rec.Body.String())
		} else {
			// abnormal system testcase
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, v.Categories, rec.Body.String())
		}
	}
}

// method to create request to test /category/list API
func CreateCategoryListRequest(e *echo.Echo, JWT string) (echo.Context, *httptest.ResponseRecorder) {
	f := make(url.Values) // create map[string][]string
	req := httptest.NewRequest(http.MethodPost, "/category/list", strings.NewReader(f.Encode()))

	// common processing
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cookie := new(http.Cookie)
	cookie.Value = "token=" + JWT
	c.SetCookie(cookie)
	return c, rec
}

type CategoryListTC struct {
	JWT        string
	Categories []domain.Category
	StatusCode int
}
