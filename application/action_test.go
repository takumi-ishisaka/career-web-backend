package application

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/infra"
	"github.com/nokin-all-of-career/career-web-backend/infra/persistence"
)

// テスト用に入っている値
// 	domain.Action{"action_00001", "自己分析", "自己分析をしてみよう", "友達とやるのがいいかも", "category_001"}
// 	domain.Action{"action_00002", "企業分析", "企業分析をしてみよう", "ネットを使おう", "category_002"}

func actionInitializaion() (*echo.Echo, ActionApplication) {
	pa := persistence.NewActionPersistence(infra.DB)
	pc := persistence.NewCategoryPersistence(infra.DB)
	pr := persistence.NewRecommendationPersistence(infra.DB)
	a := NewActionApplication(pa, pr, pc)
	e := echo.New()
	return e, a
}

func TestActionShow(t *testing.T) {
	// Setup
	// return *echo.Echo and ActionApplication
	e, a := actionInitializaion()
	var actionShowTC = []ActionShowTC{
		{"auth", "action_0001", domain.Action{"action_0001", "自己分析", "自己分析をしてみよう", "友達とやるのがいいかも", "category_001"}, 200},
		{"auth", "action_0000", domain.Action{}, 500},
	}

	for _, v := range actionShowTC {
		// c is echo.Context
		// rec is *http.ResponseRecoder
		c, rec := CreateActionShowRequest(e, v.JWT, v.ActionID)

		if assert.NoError(t, a.ActionsShow(c)) {
			// normal system testcase
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, v.Action, rec.Body.String())
		} else {
			// abnormal system testcase
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, v.Action, rec.Body.String())
		}
	}
}

func TestActionCategory(t *testing.T) {
	// Setup
	// return is *echo.Echo and ActionHandler
	e, a := actionInitializaion()
	var actions []domain.Action
	actions = append(actions,
		domain.Action{"action_00002", "企業分析", "企業分析をしてみよう", "ネットを使おう", "category_002"},
		domain.Action{"action_00002", "企業分析", "企業分析をしてみよう", "ネットを使おう", "category_002"},
	)
	var actionCategoryTC = []ActionCategoryTC{
		{"auth", "category_001", actions, 200},
		{"auth", "category_000", actions, 500},
	}

	for _, v := range actionCategoryTC {
		// c is echo.Context
		// rec is *http.ResponseRecoder
		c, rec := CreateActionCategoryRequest(e, v.JWT, v.CategoryID)

		if assert.NoError(t, a.ActionDetailShow(c)) {
			// normal system testcase
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, v.Actions, rec.Body.String())
		} else {
			// abnormal system testcase
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, v.Actions, rec.Body.String())
		}
	}
}

// method to create request to test /action/show API
func CreateActionShowRequest(e *echo.Echo, JWT, actionID string) (echo.Context, *httptest.ResponseRecorder) {
	f := make(url.Values) // create map[string][]string
	req := httptest.NewRequest(http.MethodPost, "/action/show?action_id="+actionID, strings.NewReader(f.Encode()))

	// common processing
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cookie := new(http.Cookie)
	cookie.Value = "token=" + JWT
	c.SetCookie(cookie)
	return c, rec
}

// mthod to create request to test /actions API
func CreateActionCategoryRequest(e *echo.Echo, JWT, categoryID string) (echo.Context, *httptest.ResponseRecorder) {
	f := make(url.Values) // create map[string][]string
	req := httptest.NewRequest(http.MethodPost, "/actions?category_id="+categoryID, strings.NewReader(f.Encode()))

	// common processing
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cookie := new(http.Cookie)
	cookie.Value = "token=" + JWT
	c.SetCookie(cookie)
	return c, rec
}

type ActionShowTC struct {
	JWT        string
	ActionID   string
	Action     domain.Action
	StatusCode int
}

type ActionCategoryTC struct {
	JWT        string
	CategoryID string
	Actions    []domain.Action
	StatusCode int
}
