package application

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/infra"
	"github.com/nokin-all-of-career/career-web-backend/infra/persistence"
)

//テスト用に入っている値
// 	domain.User{"自動生成されたID", "aaa@aaa.com", "aaaaaaaaのハッシュ化"}
// 	domain.User{"自動生成されたID", "bbb@bbb.com", "bbbbbbbbのハッシュ化"}
// 	domain.User{"自動生成されたID", "ccc@ccc.com", "ccccccccのハッシュ化"}

func userInitializaion() (*echo.Echo, UserApplication) {
	p := persistence.NewUserPersistence(infra.DB)
	a := NewUserApplication(p)
	e := echo.New()
	return e, a
}

func TestUserSignup(t *testing.T) {
	// return *echo.Echo and UserHandler
	e, a := userInitializaion()
	var signupTC = []SignupTC{
		{"xxx@xxx.com", "xxxxxxxx", 200},
		{"yyy@yyy.com", "yyyyyyyyyyyyy", 200},
		{"yyy@yyy", "yyyyyyy", 500},
		//{"zzzzzz.com", "zzzzzzzz", 500},
	}

	for _, v := range signupTC {
		// c is echo.Context
		// rec is *http.ResponseRecoder
		c, rec := CreateUserSignupRequest(e, v.Email, v.Password)
		err := a.UserSignup(c)

		if assert.NoError(t, err) {
			// normal system testcase
			assert.Equal(t, http.StatusOK, rec.Code)
		} else {
			// abnormal system testcase
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	}
}

func TestUserSignin(t *testing.T) {
	// return *echo.Echo and UserHandler
	e, a := userInitializaion()
	var signinTC = []SigninTC{
		{"aaa@aaa.com", "aaaaaaaa", 200},
		{"xxx@xxx.com", "xxxxxxxx", 500},
	}

	for _, v := range signinTC {
		// c is echo.Context
		// rec is *http.ResponseRecoder
		c, rec := CreateUserSigninRequest(e, v.Email, v.Password)

		if assert.NoError(t, a.UserSignin(c)) {
			// normal system testcase
			assert.Equal(t, http.StatusOK, rec.Code)
		} else {
			// abnormal system testcase
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	}
}

func TestUserGet(t *testing.T) {
	// return *echo.Echo and UserHandler
	e, a := userInitializaion()
	var userGetTC = []UserGetTC{
		{"auth", domain.User{"a", "aaa@aaa.com", "aaaaaaaa", 1}, 200},
		// {"auth", domain.User{}, 500},
	}

	for _, v := range userGetTC {
		// c is echo.Context
		// rec is *http.ResponseRecoder
		c, rec := CreateUserGetRequest(e, v.JWT)

		if assert.NoError(t, a.UserGet(c)) {
			// normal system testcase
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, v.User, rec.Body.String())
		} else {
			// abnormal system testcase
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, v.User, rec.Body.String())
		}
	}
}

// method to create request to test /user/signup API
func CreateUserSignupRequest(e *echo.Echo, email, password string) (echo.Context, *httptest.ResponseRecorder) {
	f := make(url.Values) // create map[string][]string for form request
	f.Set("email", email)
	f.Set("password", password)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8081/user/signup", bytes.NewBufferString((f.Encode())))

	// request date type is form value（application/json is not）
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	// common processing
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

// method to create request to test /user/signin API
func CreateUserSigninRequest(e *echo.Echo, email, password string) (echo.Context, *httptest.ResponseRecorder) {
	f := make(url.Values) // create map[string][]string
	f.Set("email", email)
	f.Set("password", password)
	req := httptest.NewRequest(http.MethodPost, "/user/signin", bytes.NewBufferString((f.Encode())))

	// request date type is form value（application/json is not）
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	// common processing
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

// method to create request to test /user/get API
func CreateUserGetRequest(e *echo.Echo, JWT string) (echo.Context, *httptest.ResponseRecorder) {
	//f := make(url.Values) // create map[string][]string
	req := httptest.NewRequest(http.MethodPost, "/user/signup", nil)

	// common processing
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cookie := new(http.Cookie)
	cookie.Value = "token=" + JWT
	c.SetCookie(cookie)
	return c, rec
}

type SignupTC struct {
	Email      string
	Password   string
	StatusCode int
}

type SigninTC struct {
	Email      string
	Password   string
	StatusCode int
}

type UserGetTC struct {
	JWT        string
	User       domain.User
	StatusCode int
}
