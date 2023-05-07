package application

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/infra"
	"github.com/nokin-all-of-career/career-web-backend/infra/persistence"
)

// テスト用に入っている値
// domain.UserAction{"auth", "action_id", "", "1", "2020/2/16"}
// domain.UserAction{"auth", "action_id2", "", "1", "2020/2/15"}
// domain.UserAction{"auth", "action_id3", "reflection_id3", "2", "2020/2/14"}
// domain.UserAction{"auth", "action_id4", "reflection_id4", "2", "2020/2/13"}

func userActionInitializaion() (*echo.Echo, UserActionApplication) {
	uap := persistence.NewUserActionPersistence(infra.DB)
	urp := persistence.NewRecommendationPersistence(infra.DB)
	a := NewUserActionApplication(uap, urp)
	e := echo.New()
	return e, a
}

func TestUserTodoActionGet(t *testing.T) {
	// Setup
	// return *echo.Echo and ProfileApplicatin
	e, a := userActionInitializaion()
	var userActions []domain.UserAction
	data, _ := time.Parse("2006-01-02", "2020-02-16")
	data2, _ := time.Parse("2006-01-02", "2020-02-15")
	userActions = append(userActions,
		domain.UserAction{"user_id_001", "action_id", 1, data},
		domain.UserAction{"user_id_002", "action_id2", 1, data2},
	)
	var userActionGetTC = []UserActionGetTC{
		{"auth", userActions, 200},
		{"auth", userActions, 500},
	}

	for _, v := range userActionGetTC {
		// c is echo.Context
		// rec is *http.ResponseRecoder
		c, rec := CreateUserTodoActionGetRequest(e, v.JWT)

		if assert.NoError(t, a.SelectTodoByUserID(c)) {
			// normal system testcase
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, v.UserActions, rec.Body.String())
		} else {
			// abnormal sytem testcase
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, v.UserActions, rec.Body.String())
		}
	}
}

// method to create request to test /userAction/todo API
func CreateUserTodoActionGetRequest(e *echo.Echo, JWT string) (echo.Context, *httptest.ResponseRecorder) {
	f := make(url.Values) // create map[string][]string
	req := httptest.NewRequest(http.MethodPost, "/userAction/todo", strings.NewReader(f.Encode()))

	// common processing
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cookie := new(http.Cookie)
	cookie.Value = "token=" + JWT
	c.SetCookie(cookie)
	return c, rec
}

func TestUserDoneActionGet(t *testing.T) {
	// Setup
	// return *echo.Echo and ProfileApplication
	e, a := userActionInitializaion()
	var userActions []domain.UserAction
	data, _ := time.Parse("2006-01-02", "2020-02-14")
	data2, _ := time.Parse("2006-01-02", "2020-02-13")
	userActions = append(userActions,
		domain.UserAction{"user_id_003", "action_id3", 2, data},
		domain.UserAction{"user_id_004", "action_id4", 2, data2},
	)
	var userActionGetTC = []UserActionGetTC{
		{"auth", userActions, 200},
		{"auth", userActions, 500},
	}

	for _, v := range userActionGetTC {
		// c is echo.Context
		// rec is *http.ResponseRecoder
		c, rec := CreateUserDoneActionGetRequest(e, v.JWT)

		if assert.NoError(t, a.SelectDoneByUserID(c)) {
			// normal system testcase

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, v.UserActions, rec.Body.String())
		} else {
			// abnormal system testcase
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, v.UserActions, rec.Body.String())
		}
	}
}

// method to create request to test /userAction/done API
func CreateUserDoneActionGetRequest(e *echo.Echo, JWT string) (echo.Context, *httptest.ResponseRecorder) {
	f := make(url.Values) // create map[string][]string
	req := httptest.NewRequest(http.MethodPost, "/userAction/done", strings.NewReader(f.Encode()))

	// common processing
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cookie := new(http.Cookie)
	cookie.Value = "token=" + JWT
	c.SetCookie(cookie)
	return c, rec
}

type UserActionGetTC struct {
	JWT         string
	UserActions []domain.UserAction
	StatusCode  int
}

func TestUserActionStatusChange(t *testing.T) {
	// Setup
	// return *echo.Echo and ActionApplication
	e, a := userActionInitializaion()
	var userActionStatusChangeTC = []UserActionStatusChangeTC{
		{"auth", "action_id", 2, 200},
		{"auth", "action_id2", 0, 200},
		{"auth", "", 4, 500},
	}

	for _, v := range userActionStatusChangeTC {
		// c is echo.Context
		// rec is *http.ResponseRecoder
		c, rec := CreateUserActionStatusChangeRequest(e, v.JWT, v.ActionID, v.Status)

		if assert.NoError(t, a.ChangeActionStatus(c)) {
			// normal system testcase
			assert.Equal(t, http.StatusOK, rec.Code)
		} else {
			// abnormal system testcase
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	}
}

// method to create request to test /userAction/change API
func CreateUserActionStatusChangeRequest(e *echo.Echo, JWT string, actionID string, status int) (echo.Context, *httptest.ResponseRecorder) {
	f := make(url.Values) // create map[string][]strin
	f.Set("status", string(status))
	req := httptest.NewRequest(http.MethodPost, "/userAction/change?action_id="+actionID, strings.NewReader(f.Encode()))

	// request date typw is form value（application/json is not）
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	// common processing
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cookie := new(http.Cookie)
	cookie.Value = "token=" + JWT
	c.SetCookie(cookie)
	return c, rec
}

type UserActionStatusChangeTC struct {
	JWT        string
	ActionID   string
	Status     int
	StatusCode int
}

type UserActionStatusFinishTC struct {
	JWT        string
	ActionID   string
	StatusCode int
}

//TODO:
func TestHandleUserActionSelectReflectionByPrimaryKey(t *testing.T) {
	e, a := userActionInitializaion()
	var userActionReflectionTC = []UserActionReflectionTC{
		{"auth", "action_id", 200},
	}
	for _, v := range userActionReflectionTC {
		c, rec := CreateSelectReflectionByPrimaryKeyRequest(e, v.JWT, v.ActionID)

		if assert.NoError(t, a.SelectReflectionByPrimaryKey(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		} else {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	}
}

//TODO:
func CreateSelectReflectionByPrimaryKeyRequest(e *echo.Echo, JWT string, actionID string) (echo.Context, *httptest.ResponseRecorder) {
	f := make(url.Values) // create map[string][]string
	f.Set("action_id", actionID)
	req := httptest.NewRequest(http.MethodGet, "/user_action/reflection/get", strings.NewReader(f.Encode()))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cookie := new(http.Cookie)
	cookie.Value = "token=" + JWT
	c.SetCookie(cookie)

	return c, rec
}

type UserActionReflectionTC struct {
	JWT        string
	ActionID   string
	StatusCode int
}

//TODO:
func TestHandleUserActionInsertReflection(t *testing.T) {
	e, a := userActionInitializaion()
	var userActionInsertReflectionTC = []UserActionInsertReflectionTC{
		{"auth", "action_id", "keep", "problem", "try", 4, 200},
	}
	for _, v := range userActionInsertReflectionTC {
		c, rec := CreateInsertReflectionRequest(e, v.JWT, v.ActionID, v.Keep, v.Problem, v.Try, v.EvaluateValue)

		if assert.NoError(t, a.SelectReflectionByPrimaryKey(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		} else {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	}
}

//TODO:
func CreateInsertReflectionRequest(e *echo.Echo, JWT string, actionID string, keep string, problem string, try string, evaluatevalue int) (echo.Context, *httptest.ResponseRecorder) {
	f := make(url.Values) // create map[string][]string
	req := httptest.NewRequest(http.MethodGet, "/user_action/reflection/"+actionID, strings.NewReader(f.Encode()))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cookie := new(http.Cookie)
	cookie.Value = "token=" + JWT
	c.SetCookie(cookie)

	return c, rec
}

type UserActionInsertReflectionTC struct {
	JWT           string
	ActionID      string
	Keep          string
	Problem       string
	Try           string
	EvaluateValue int
	StatusCode    int
}
