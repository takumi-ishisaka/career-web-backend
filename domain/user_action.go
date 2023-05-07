package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

const (
	// NOTSELECT : notselect string
	NOTSELECT = "notselect"
	// STATENOTSELECT : notselect number
	STATENOTSELECT = 0
	// TODO : todo string
	TODO = "todo"
	// STATETODO : todo number
	STATETODO = 1
	// DONE : done string
	DONE = "done"
	// STATEDONE : done number
	STATEDONE = 2
	// APPROVAL : approval string
	APPROVAL = "approval"
	// STATEAPPROVAL : approval number
	STATEAPPROVAL = 3
	// AGAIN : again string
	AGAIN = "again"
	// STATEAGAIN : again number
	STATEAGAIN = 4
)

// UserAction : userAction domain
type UserAction struct {
	UserActionID  string    `json:"user_action_id"`
	UserID        string    `json:"user_id"`
	ActionID      string    `json:"action_id"`
	Status        int       `json:"status"`
	Do            string    `json:"do" form:"do"`
	Reflection    string    `json:"reflection" form:"reflection"`
	NextAction    string    `json:"next_action" form:"next_action"`
	EvaluateValue int       `json:"evaluate_value" form:"evaluate_value"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GenerateUserActionID : generate user action ID
func GenerateUserActionID() (string, error) {
	// get userID by UUID
	userActionID, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "domain.useraction.GenerateUserActionID")
	}

	return userActionID.String(), nil
}

// Recommendation : recommendation domain
type Recommendation struct {
	ActionID            string
	RecommendationValue float32
	DoneUserNum         int
}

const (
	// FIRSTDONEUSERNUM : first done user num
	FIRSTDONEUSERNUM = 1
)

// Reflection : reflection domain
type Reflection struct {
	Do            string `json:"do" form:"do"`
	Reflection    string `json:"reflection" form:"reflection"`
	NextAction    string `json:"next_action" form:"next_action"`
	EvaluateValue int    `json:"evaluate_value" form:"evaluate_value"`
}

// ReflectionValidation : reflection request validation
func ReflectionValidation(do, reflection, nextAction string) error {
	insertReflection := &ReflectionValidate{Do: do, Reflection: reflection, NextAction: nextAction}
	validate := validator.New()
	err := validate.Struct(insertReflection)
	if err != nil {
		return errors.Wrap(err, "domain.user_action.ReflectionValidation()")
	}
	return nil
}

// ReflectionValidate : reflection validation struct
type ReflectionValidate struct {
	Do         string `validate:"required"`
	Reflection string `validate:"required"`
	NextAction string `validate:"required"`
}
