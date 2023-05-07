package application

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/infra"
	"github.com/nokin-all-of-career/career-web-backend/infra/persistence"
	"github.com/stretchr/testify/assert"
)

// テスト用に入っている値
// domain.Profile{"auth", "あいう", "えお大学", "かき専攻", "くけこになる", "よろしくお願いします。", "/path/", "週休二日制", 1}
// domain.Profile{"auth2", "さしす", "せそ大学", "たち専攻", "つてとになる", "よろしくお願いします。", "/path2/", "若いうちから挑戦できる", 2}
// domain.Profile{"auth3","name","shizuokaUniversity","computerscience",2021,"enginner","platform","mynameisTanaka","Imagepath",1,50.0}
func profileInitializaion() (*echo.Echo, ProfileApplication) {
	p := persistence.NewProfilePersistence(infra.DB)
	pua := persistence.NewUserActionPersistence(infra.DB)
	a := NewProfileApplication(p, pua)
	e := echo.New()
	return e, a
}

func TestProfileGet(t *testing.T) {
	// Setup
	// return *echo.Echo and ProfileHandler
	e, a := profileInitializaion()
	var profileGetTC = []ProfileGetTC{
		{"auth", domain.Profile{"user_id_001", "name", "shizuokaUniversity", "computerscience", 2021, "enginner", "platform", "mynameisTanaka", "Imagepath", 1, 50.0}, 200},
		{"auth", domain.Profile{}, 500},
	}

	for _, v := range profileGetTC {
		// c is echo.Context
		// rec is *http.ResponseRecoder
		c, rec := CreateProfileGetRequest(e, v.JWT)

		if assert.NoError(t, a.ProfileGet(c)) {
			// normal system testcase
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, v.Profile, rec.Body.String())
		} else {
			// abnormal system testcase
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, v.Profile, rec.Body.String())
		}
	}
}

// method to create request to test /profile/get API
func CreateProfileGetRequest(e *echo.Echo, JWT string) (echo.Context, *httptest.ResponseRecorder) {
	f := make(url.Values) // create map[string][]string
	req := httptest.NewRequest(http.MethodPost, "/profile/get", strings.NewReader(f.Encode()))

	// common processing
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cookie := new(http.Cookie)
	cookie.Value = "token=" + JWT
	c.SetCookie(cookie)
	return c, rec
}

type ProfileGetTC struct {
	JWT        string
	Profile    domain.Profile
	StatusCode int
}

func TestProfileUpsert(t *testing.T) {
	// Setup
	// return *echo.Echo and ActionHandler
	e, a := profileInitializaion()
	var profileUpsertTC = []ProfileUpsertTC{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIzYjMwNjAyMC00Y2QzLTQyYzUtYmY2ZS0xMDFlNWFhMzdjNzYiLCJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTE2MjM5MDIyfQ.5eGsvQYcN-aeABDpi2ECq7JaxzrHNaTEA5PfkvDkbmQ", "あいう", "えお大学", "かき専攻", 2021, "くけこになる", "いろんな技術を身に着けたい", "/path/", "週休二日制", 4, 200},
		// {"auth3", "なにぬ", "ねの大学", "はひ専攻", 1999, "ふへほになる", "よろしくお願いします。", "/path3/", "週休二日制", 1, 200},
		// {"auth4", "", "", "", 0, "", "", "", "", 1, 500},
	}

	for _, v := range profileUpsertTC {
		// c is echo.Context
		// rec is *http.ResponseRecoder
		c, rec := CreateProfileUpsertRequest(e, v.JWT, v.Name, v.University, v.Major, v.Dream, v.Sentence, v.ImagePath, v.CoreValue, v.GraduationYear, v.DeviationValue)

		if assert.NoError(t, a.ProfileUpsert(c)) {
			// normal system testcase
			assert.Equal(t, http.StatusOK, rec.Code)
		} else {
			// abnormal system testcase
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	}
}

// method to create request to test /profile/upsert API
func CreateProfileUpsertRequest(e *echo.Echo, JWT string, name, university, major, dream, sentence, imagePath, coreValue string, GraduationYear int, deviationValue float64) (echo.Context, *httptest.ResponseRecorder) {
	f := make(url.Values) // create map[string][]string
	f.Set("name", name)
	f.Set("university", university)
	f.Set("major", major)
	f.Set("graduation_year", string(GraduationYear))
	f.Set("dream", dream)
	f.Set("sentence", sentence)
	f.Set("imagePath", imagePath)
	f.Set("coreValue", coreValue)
	f.Set("deviationValue", strconv.FormatFloat(deviationValue, 'f', 4, 64))
	req := httptest.NewRequest(http.MethodPost, "/profile/upsert", strings.NewReader(f.Encode()))

	// request date type is form value（application/json is not.）
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	// common processing
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cookie := new(http.Cookie)
	cookie.Value = "token=" + JWT
	c.SetCookie(cookie)
	return c, rec
}

type ProfileUpsertTC struct {
	JWT            string
	Name           string
	University     string
	Major          string
	GraduationYear int
	Dream          string
	Sentence       string
	ImagePath      string
	CoreValue      string
	DeviationValue float64
	StatusCode     int
}
