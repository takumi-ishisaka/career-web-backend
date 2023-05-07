package server

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"

	"github.com/nokin-all-of-career/career-web-backend/application"
	"github.com/nokin-all-of-career/career-web-backend/configs"
	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/infra"
	"github.com/nokin-all-of-career/career-web-backend/infra/persistence"
)

// ValidMentorMiddleware : valid mentor account middleware
func ValidMentorMiddleware(userIDs []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user")
			user := token.(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			userID := claims["sub"].(string)
			for _, v := range userIDs {
				if v == userID {
					return next(c)
				}
			}
			err := errors.New("Valid User by MentorID")
			return errors.Wrap(err, "Middleware.ValidMentorMiddleware()")
		}
	}
}

// Run : run server
func Run(runOption int) error {

	domain.InitSignature(runOption)

	if runOption == 2 {
		err := infra.NewDBConnection()
		if err != nil {
			return errors.Wrap(err, "Run()")
		}
	} else {
		err := configs.InitConfig()
		if err != nil {
			return errors.Wrap(err, "Run()")
		}
		err = infra.NewLocalDBConnection()
		if err != nil {
			return errors.Wrap(err, "Run()")
		}
	}

	infra.NewSMTPConnection()

	if runOption == 1 {
		infra.NewRedisLocalConnection()
	} else {
		infra.NewRedisConnection()
	}

	err := infra.NewStorageConnection()
	if err != nil {
		return errors.Wrap(err, "Run()")
	}

	// dependency injection
	// persistemce depends on DB
	userPersistence := persistence.NewUserPersistence(infra.DB)
	categoryPersistence := persistence.NewCategoryPersistence(infra.DB)
	actionPersistence := persistence.NewActionPersistence(infra.DB)
	userActionPersistence := persistence.NewUserActionPersistence(infra.DB)
	profilePersistence := persistence.NewProfilePersistence(infra.DB)
	feedbackPersistence := persistence.NewFeedBackPersistence(infra.DB)

	// application depends on persistence
	userApplication := application.NewUserApplication(userPersistence, actionPersistence, userActionPersistence)
	categoryApplication := application.NewCategoryApplication(categoryPersistence)
	actionApplication := application.NewActionApplication(actionPersistence, userActionPersistence, categoryPersistence)
	userActionApplication := application.NewUserActionApplication(userActionPersistence, actionPersistence, userPersistence)
	profileApplication := application.NewProfileApplication(profilePersistence, userActionPersistence)
	feedbackApplication := application.NewFeedBackApplication(feedbackPersistence, userActionPersistence, actionPersistence, userPersistence)

	// put Action date on memory
	err = persistence.InitAction()
	if err != nil {
		return errors.Wrap(err, "Run()")
	}

	// put Category date on memory
	err = persistence.InitCategory()
	if err != nil {
		return errors.Wrap(err, "Run()")
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// cache certificates
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	// add middleware
	// e.Use(middleware.Recover())
	// TODO: change if host decided
	// TODO: if we use cookie in ajax, we should add Access-Control-Allow-Credentials header
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://career-web-frontend-wd2f3kdf2q-an.a.run.app", "http://localhost:8080/"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderXRequestedWith, echo.HeaderAccept, echo.HeaderAccessControlAllowOrigin, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	// logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time": "${time_rfc3339}", "method": "${method}", "uri": "${uri}", "status": "${status}", ` +
			`"error": "${error}"}` + "\n",
	}))

	// unauthenticated route
	e.POST("/user/signup", userApplication.UserSignup)
	e.GET("/user/validation", userApplication.EmailValidation)
	e.POST("/user/signin", userApplication.UserSignin)

	JWTConf := middleware.JWTConfig{SigningKey: []byte(domain.JWTSignature)}

	// authenticated route
	e.GET("/user/get", userApplication.UserGet, middleware.JWTWithConfig(JWTConf))       // OK
	e.PUT("/user/update", userApplication.UserUpdate, middleware.JWTWithConfig(JWTConf)) // OK
	// e.POST("/contact", userApplication.ContactSend, middleware.JWTWithConfig(JWTConf))                      // NOTOK
	e.GET("/actions", actionApplication.ActionsShow, middleware.JWTWithConfig(JWTConf))                     // OK
	e.GET("/action/show", actionApplication.ActionDetailShow, middleware.JWTWithConfig(JWTConf))            // OK
	e.GET("/category/list", categoryApplication.CategoryList, middleware.JWTWithConfig(JWTConf))            // OK
	e.GET("/profile", profileApplication.ProfileGet, middleware.JWTWithConfig(JWTConf))                     // OK
	e.PUT("/profile/upsert", profileApplication.ProfileUpsert, middleware.JWTWithConfig(JWTConf))           // OK
	e.POST("/profile/search", profileApplication.ProfilesSearch, middleware.JWTWithConfig(JWTConf))         // OK
	e.GET("/user_action/todo", userActionApplication.TodoUserActionsGet, middleware.JWTWithConfig(JWTConf)) // OK
	e.GET("/user_action/done", userActionApplication.DoneUserActionsGet, middleware.JWTWithConfig(JWTConf)) // OK
	e.GET("/user_action/approval", userActionApplication.ApprovalUserActionsGet, middleware.JWTWithConfig(JWTConf))
	e.GET("/user_action/again", userActionApplication.AgainUserActionsGet, middleware.JWTWithConfig(JWTConf))
	e.GET("/user_action/done_approval", userActionApplication.DoneApprovalUserActionsGet, middleware.JWTWithConfig(JWTConf))          // OK
	e.PUT("/user_action/change", userActionApplication.UserActionStatusChange, middleware.JWTWithConfig(JWTConf))                     // OK
	e.GET("/user_action/reflection", userActionApplication.UserActionReflectionGet, middleware.JWTWithConfig(JWTConf))                // OK
	e.POST("/user_action/reflection/:action_id", userActionApplication.UserActionReflectionInsert, middleware.JWTWithConfig(JWTConf)) // OK

	// TODO: create valid mentor on middleware
	e.POST("/mentor/signin", userApplication.MentorSignin) // OK

	e.GET("/mentor/done_user_actions", userActionApplication.DoneUserActionList, middleware.JWTWithConfig(JWTConf), ValidMentorMiddleware(domain.UserIDs)) // OK
	e.POST("/mentor/feedback", feedbackApplication.FeedBackInsert, middleware.JWTWithConfig(JWTConf), ValidMentorMiddleware(domain.UserIDs))               // OK
	e.GET("/feedback/get", feedbackApplication.FeedBacksGet, middleware.JWTWithConfig(JWTConf))                                                            // OK
	e.GET("/user_action/feedback/list", feedbackApplication.FeedBackList, middleware.JWTWithConfig(JWTConf))                                               // OK
	// designate the port number and start
	// e.Logger.Fatal(e.StartAutoTLS(":443"))
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8081"
	}
	e.Logger.Fatal(e.Start(port))
	return nil
}
