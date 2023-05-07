package domain

import (
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

// Profile : profile domain
type Profile struct {
	UserID             string  `json:"user_id"`
	Name               string  `json:"name"`
	University         string  `json:"university"`
	Major              string  `json:"major"`
	GraduationYear     int     `json:"graduation_year"`
	AspiringOccupation string  `json:"aspiring_occupation"`
	AspiringField      string  `json:"aspiring_field"`
	Sentence           string  `json:"sentence"`
	ImagePath          string  `json:"iamge_path"`
	JobHuntingStatus   int     `json:"job_hunting_status"`
	DeviationValue     float32 `json:"deviation_value"` // Done tasks number
}

// DeviationValueCalculation : calucurate deviation value
func DeviationValueCalculation(doneNum, doneAverage float32) float32 {
	deviationValue := (doneNum-doneAverage)/2.0 + 50.0
	deviationValue = float32(int(deviationValue*10.0) / 10.0)

	return deviationValue
}

// UpsertValidation : validate ProfileUpsert request
func UpsertValidation(name, university, major, aspiringOccupation, aspiringField, sentence, imagePath string, graduationYear int) error {
	upsert := &UpsertValidate{Name: name, University: university, Major: major, AspiringOccupation: aspiringOccupation, AspiringField: aspiringField, Sentence: sentence, ImagePath: imagePath, GraduationYear: graduationYear}
	validate := validator.New()
	err := validate.Struct(upsert)
	if err != nil {
		return errors.Wrap(err, "profileApplication.UpsertValidation()")
	}

	return nil
}

// UpsertValidate : upsert validate struct
type UpsertValidate struct {
	Name               string `validate:"required"`
	University         string `validate:"required"`
	Major              string `validate:"required"`
	GraduationYear     int    `validate:"required"`
	AspiringOccupation string `validate:""`
	AspiringField      string `validate:""`
	Sentence           string `validata:""`
	ImagePath          string `validata:""`
}
