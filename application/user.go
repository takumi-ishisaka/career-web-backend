package application

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/domain/repository"
	"github.com/nokin-all-of-career/career-web-backend/infra"
)

// UserApplication : user application
type UserApplication interface {
	UserSignup(c echo.Context) error
	EmailValidation(c echo.Context) error
	UserSignin(c echo.Context) error
	MentorSignin(c echo.Context) error
	UserGet(c echo.Context) error
	UserUpdate(c echo.Context) error
	// ContactSend(c echo.Context) error
}

type userApplication struct {
	userRepository       repository.UserRepository
	actionRepository     repository.ActionRepository
	userActionRepository repository.UserActionRepository
}

// NewUserApplication : create application about User
func NewUserApplication(ur repository.UserRepository, ar repository.ActionRepository, uar repository.UserActionRepository) UserApplication {
	return &userApplication{
		userRepository:       ur,
		actionRepository:     ar,
		userActionRepository: uar,
	}
}

// UserSignup : insert to user table.
func (ua userApplication) UserSignup(c echo.Context) error {
	userRequest := new(UserRequest)
	if err := c.Bind(userRequest); err != nil {
		return errors.Wrap(err, "userApplication.UserSignup()")
	}

	err := domain.InsertValidation(userRequest.Email, userRequest.Password)
	if err != nil {
		return errors.Wrap(err, "userApplicaton.UserSignup()")
	}

	userID, err := domain.GenerateUserID()
	if err != nil {
		return errors.Wrap(err, "userApplication.UserSignup()")
	}

	hashPassword, err := domain.GenerateHashPassword(userRequest.Password)
	if err != nil {
		return errors.Wrap(err, "userApplication.UserSignup()")
	}

	//RedisにkeyとしてUserIDを入れて、valueにハッシュ化したpassowordを入れる。
	user, err := ua.userRepository.SelectByEmail(userRequest.Email)
	if err != nil {
		return errors.Wrap(err, "userApplication.UserSignup()")
	}
	if user != nil {
		err = errors.New("the email is exit")
		return errors.Wrap(err, "userApplication.UserSignup()")
	}

	//TODO:error処理の追加
	infra.Conn.Do("APPEND", userID, userRequest.Email+":"+hashPassword)
	infra.Conn.Do("EXPIRE", userID, time.Minute*60)

	//create validation url
	validationURL := "http://localhost:8081/user/validation?id=" + userID
	fmt.Println(validationURL)
	//sendEmail
	err = ua.userRepository.SendValidationEmail(infra.OperationEmail, userRequest.Email, validationURL)
	if err != nil {
		return errors.Wrap(err, "userApplication.UserSignup()")
	}
	return nil
}

func (ua userApplication) EmailValidation(c echo.Context) error {

	emailValidationGetRequest := new(EmailValidationGetRequest)
	if err := c.Bind(emailValidationGetRequest); err != nil {
		return errors.Wrap(err, "userApplication.EmailValidation()")
	}

	exists, err := redis.Bool(infra.Conn.Do("EXISTS", emailValidationGetRequest.UserID))
	if err != nil {
		return errors.Wrap(err, "userApplication.EmailValidation()")
	}

	if exists != true {
		err = errors.New("the URI is not valid because the ID is not correct")
		return errors.Wrap(err, "userApplication.EmailValidation()")
	}
	value, err := redis.String(infra.Conn.Do("GET", emailValidationGetRequest.UserID))
	if err != nil {
		return errors.Wrap(err, "userApplication.EmailValidation()")
	}
	data := strings.Split(value, ":")
	//data[0]=email,data[1]=password
	//DBにUser情報を追加
	err = ua.userRepository.Insert(emailValidationGetRequest.UserID, data[0], data[1])
	if err != nil {
		return errors.Wrap(err, "userApplication.EmailValidation()")
	}
	//Redisから該当UserIDを削除
	infra.Conn.Do("DELETE", emailValidationGetRequest.UserID)
	// insert user_action into user_action table
	actionIDs, err := ua.actionRepository.SelectByActiontype(domain.TUTORIAL)
	for _, actionID := range actionIDs {
		userActionID, err := domain.GenerateUserActionID()
		if err != nil {
			return errors.Wrap(err, "userActionApplication.EmailValidation()")
		}
		err = ua.userActionRepository.InsertUserAction(userActionID, emailValidationGetRequest.UserID, actionID, domain.STATETODO)
		if err != nil {
			return errors.Wrap(err, "userActionApplication.EmailValidation()")
		}
	}
	return nil
}

// UserSignin : sign in by user email and pass
func (ua userApplication) UserSignin(c echo.Context) error {
	userRequest := new(UserRequest)
	if err := c.Bind(userRequest); err != nil {
		return errors.Wrap(err, "userApplication.UserSignin()")
	}

	user, err := ua.userRepository.SelectByEmail(userRequest.Email)
	if err != nil {
		return errors.Wrap(err, "userApplication.UserSignin()")
	}

	err = domain.CheckPassword(user.Password, userRequest.Password)
	if err != nil {
		return errors.Wrap(err, "userApplication.UserSignin()")
	}

	err = domain.CheckStatus(user.Status)
	if err != nil {
		return errors.Wrap(err, "userApplication.UserSignin()")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	// claims["iss"] = "https://career-272208.appspot.com/" // Issuer
	claims["admin"] = true                                // Authority
	claims["sub"] = user.UserID                           // Unique value
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // effective date

	// generate encoded token and send it as response.
	t, err := token.SignedString([]byte(domain.JWTSignature))
	if err != nil {
		return errors.Wrap(err, "userApplication.UserSignin()")
	}

	// Cookie create
	// cookie := &http.Cookie{
	// 	Name:     "token",
	// 	Value:    t,
	// 	Expires:  time.Now().Add(time.Hour * 24),
	// 	HttpOnly: true,
	// 	Secure:   true,
	// 	Domain: "all-of-career-fhjfljalxa-uc.a.run.app",
	// }
	// c.SetCookie(cookie)

	log.SetOutput(os.Stdout)
	log.Println(user.UserID)
	return c.JSON(http.StatusOK, SigninResponse{Token: t, Status: user.Status})
}

func (ua userApplication) MentorSignin(c echo.Context) error {
	userRequest := new(UserRequest)
	if err := c.Bind(userRequest); err != nil {
		return errors.Wrap(err, "userApplication.UserSignin()")
	}

	user, err := ua.userRepository.SelectByEmail(userRequest.Email)
	if err != nil {
		return errors.Wrap(err, "userApplication.UserSignin()")
	}

	err = domain.CheckPassword(user.Password, userRequest.Password)
	if err != nil {
		return errors.Wrap(err, "userApplication.UserSignin()")
	}

	err = domain.CheckMentorStatus(user.Status)
	if err != nil {
		return errors.Wrap(err, "userApplication.UserSignin()")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	// claims["iss"] = "https://career-272208.appspot.com/" // Issuer
	claims["admin"] = true                                // Authority
	claims["sub"] = user.UserID                           // unique value
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // effective date

	// generate encoded token and send it as response.
	t, err := token.SignedString([]byte(domain.JWTSignature))
	if err != nil {
		return errors.Wrap(err, "userApplication.UserSignin()")
	}

	return c.JSON(http.StatusOK, SigninResponse{Token: t, Status: user.Status})
}

// UserGet : getting user info.
func (ua userApplication) UserGet(c echo.Context) error {
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "profileApplication.ProfileGet()")
	}

	// get user information by userID
	user, err := ua.userRepository.SelectByPrimaryKey(userID)
	if err != nil {
		return errors.Wrap(err, "userApplication.UserGet()")
	}

	return c.JSON(http.StatusOK, user)
}

// UserUpdate : update user info
func (ua userApplication) UserUpdate(c echo.Context) error {
	userUpdateRequest := new(UserUpdateRequest)
	if err := c.Bind(userUpdateRequest); err != nil {
		return errors.Wrap(err, "userApplication.UserUpdate()")
	}

	err := domain.UpdateValidation(userUpdateRequest.Email)
	if err != nil {
		return errors.Wrap(err, "userApplicaton.UserUpdate()")
	}

	// get userID from JWT
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "userApplication.UserUpdate()")
	}

	err = ua.userRepository.Update(userID, userUpdateRequest.Email)
	if err != nil {
		return errors.Wrap(err, "userApplication.UserUpdate()")
	}

	return nil

}

// // ContactSend : method that send question by email
// func (ua userApplication) ContactSend(c echo.Context) error {
// 	// get userID from JWT
// 	var userID string
// 	if token := c.Get("user"); token != nil {
// 		user := token.(*jwt.Token)
// 		claims := user.Claims.(jwt.MapClaims)
// 		userID = claims["sub"].(string)
// 	} else {
// 		err := errors.New("not *jwt.Token")
// 		return errors.Wrap(err, "userApplication.ContactSend()")
// 	}

// 	contactSendRequest := new(ContactSendRequest)
// 	if err := c.Bind(contactSendRequest); err != nil {
// 		return errors.Wrap(err, "userApplication.ContactSend()")
// 	}

// 	err := domain.ContactSendValidation(contactSendRequest.Title, contactSendRequest.Content)
// 	if err != nil {
// 		return errors.Wrap(err, "userAppication.ContatSend()")
// 	}

// 	user, err := ua.userRepository.SelectByPrimaryKey(userID)
// 	if err != nil {
// 		return errors.Wrap(err, "userApplication.ContatSend()")
// 	}

// 	err = ua.userRepository.ContactSend(user.Email, contactSendRequest.Title, contactSendRequest.Content)
// 	if err != nil {
// 		return errors.Wrap(err, "userApplication.ContatSend()")
// 	}

// 	return nil
// }

// UserRequest : user request
type UserRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

// UserUpdateRequest : user update request
type UserUpdateRequest struct {
	Email string `form:"email"`
}

// ContactSendRequest : contact request
type ContactSendRequest struct {
	Title   string `form:"title"`
	Content string `form:"content"`
}

// SigninResponse : signin response
type SigninResponse struct {
	Token  string `json:"token"`
	Status int    `json:"status"`
}

type EmailValidationGetRequest struct {
	UserID string `query:"id"`
}
