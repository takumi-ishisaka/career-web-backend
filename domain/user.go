package domain

import (
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/nokin-all-of-career/career-web-backend/configs"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

// User : user domain
type User struct {
	UserID        string    `json:"user_id"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Status        int       `json:"status"`
	LastLoginTime time.Time `json:"last_login_time"`
}

// this variable is status value
const (
	MENTOR = 75392 // Server-accepted mentors status
	USER   = 1     // Server-accepted users status
	FREEZE = -1    // Server-unaccepted users status
)

// JWTSignature : signature for jwt
var JWTSignature string

// InitSignature : set signature foe jwt
func InitSignature(runOption int) {
	if runOption == 2 {
		JWTSignature = os.Getenv("SIGNATURE")
	} else if runOption == 1 {
		JWTSignature = configs.Config.Signature
	}
}

// InsertValidation : validate Insert() elements
func InsertValidation(email, password string) error {
	insert := &InsertValidate{Email: email, Password: password}
	validate := validator.New()
	err := validate.Struct(insert)
	if err != nil {
		return errors.Wrap(err, "domain.user.InsertValidation()")
	}

	return nil
}

// UserIDs : for mentor user
var UserIDs []string = []string{"709e07db-5837-451e-a2d2-265bdc8f8a73"}

// InsertValidate : Insert() validation struct
type InsertValidate struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=19"`
}

// GenerateUserID : generate userID
func GenerateUserID() (string, error) {
	// get userID by UUID
	userID, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "domain.user.GenerateUserID()")
	}

	return userID.String(), nil
}

// GenerateHashPassword : generate hash pass
func GenerateHashPassword(password string) (string, error) {
	passwordDigest, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "domain.user.GenerateHashPassword()")
	}

	return string(passwordDigest), nil
}

// CheckPassword : check password
func CheckPassword(hashPassword, password string) error {
	// check password
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return errors.Wrap(err, "domain.user.CheckPassword()")
	}
	return nil
}

// CheckStatus : middleware which check user's status
func CheckStatus(status int) error {
	if status == FREEZE {
		err := errors.New("This account is Freeze account")
		return errors.Wrap(err, "domain.user.CheckStatus()")
	}
	return nil
}

// CheckMentorStatus : middleware which check mentor's status
func CheckMentorStatus(status int) error {
	if status == FREEZE {
		err := errors.New("This account is Freeze account")
		return errors.Wrap(err, "domain.user.CheckStatus()")
	} else if status == USER {
		err := errors.New("This account is not Mentor account")
		return errors.Wrap(err, "domain.user.CheckStatus()")
	}
	return nil
}

// UpdateValidation : validate Upsert() element
func UpdateValidation(email string) error {
	update := &UpdateValidate{Email: email}
	validate := validator.New()
	err := validate.Struct(update)
	if err != nil {
		return errors.Wrap(err, "domain.user.UpdateValidation()")
	}

	return nil
}

// UpdateValidate : Upsert() validation struct
type UpdateValidate struct {
	Email string `validate:"required,email"`
}

// ContactSendValidation : validate ContactSend() elements
func ContactSendValidation(title, content string) error {
	contactSend := &ContactSendValidate{Title: title, Content: content}
	validate := validator.New()
	err := validate.Struct(contactSend)
	if err != nil {
		return errors.Wrap(err, "domain.user.ContactSendValidation()")
	}

	return nil
}

// ContactSendValidate : ContactSend() validation struct
type ContactSendValidate struct {
	Title   string `validate:"required"`
	Content string `validate:"required"`
}
