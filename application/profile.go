package application

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/domain/repository"
)

// ProfileApplication : profile application
type ProfileApplication interface {
	ProfileGet(c echo.Context) error
	ProfileUpsert(c echo.Context) error
	ProfilesSearch(c echo.Context) error
}

type profileApplication struct {
	profileRepository    repository.ProfileRepository
	userActionRepository repository.UserActionRepository
}

// NewProfileApplication : create Application about profile
func NewProfileApplication(pr repository.ProfileRepository, uar repository.UserActionRepository) ProfileApplication {
	return &profileApplication{
		profileRepository:    pr,
		userActionRepository: uar,
	}
}

func (pa profileApplication) ProfileGet(c echo.Context) error {
	// get userID from JWT
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "profileApplication.ProfileGet()")
	}

	// calculate done action count by user_id
	doneCount, err := pa.userActionRepository.GetDoneActionCountByUserID(userID)
	if err != nil {
		return errors.Wrap(err, "profileApplication.ProfileGet()")
	}

	var deviationValue float32
	if doneCount != 0 {
		// calculate avarage of done action count per user_id
		doneCountSum, users, err := pa.userActionRepository.GetDoneActionCountSum()
		doneCountAverage := float32(doneCountSum / users)
		if err != nil {
			return errors.Wrap(err, "profileApplication.ProfileGet()")
		}
		deviationValue = domain.DeviationValueCalculation(doneCount, doneCountAverage)
	}

	err = pa.profileRepository.InsertDeviationValue(deviationValue, userID)
	if err != nil {
		return errors.Wrap(err, "profileApplication.ProfileGet()")
	}

	profile, err := pa.profileRepository.SelectByPrimaryKey(userID)
	if err != nil {
		return errors.Wrap(err, "profileApplication.ProfileGet()")
	}

	profile.DeviationValue = float32(int(profile.DeviationValue*10.0) / 10.0)

	return c.JSON(http.StatusOK, profile)
}

func (pa profileApplication) ProfileUpsert(c echo.Context) error {
	// get userID from JWT
	var userID string
	if token := c.Get("user"); token != nil {
		user := token.(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID = claims["sub"].(string)
	} else {
		err := errors.New("not *jwt.Token")
		return errors.Wrap(err, "profileApplication.ProfileGet()")
	}

	profileUpsertRequest := new(ProfileUpsertRequest)
	if err := c.Bind(profileUpsertRequest); err != nil {
		return errors.Wrap(err, "profileApplication.ProfileUpsert()")
	}

	err := domain.UpsertValidation(profileUpsertRequest.Name, profileUpsertRequest.University, profileUpsertRequest.Major, profileUpsertRequest.AspiringOccupation, profileUpsertRequest.AspiringField, profileUpsertRequest.Sentence, profileUpsertRequest.ImagePath, profileUpsertRequest.GraduationYear)
	if err != nil {
		return errors.Wrap(err, "profileApplication.ProfileUpsert()")
	}

	// calculate done action count by user_id
	doneCount, err := pa.userActionRepository.GetDoneActionCountByUserID(userID)
	if err != nil {
		return errors.Wrap(err, "profileApplication.ProfileUpsert()")
	}

	var deviationValue float32
	if doneCount != 0 {
		// calculate avarage of done action count per user_id
		doneCountSum, users, err := pa.userActionRepository.GetDoneActionCountSum()
		doneCountAverage := float32(doneCountSum / users)
		if err != nil {
			return errors.Wrap(err, "profileApplication.ProfileUpsert()")
		}
		deviationValue = domain.DeviationValueCalculation(doneCount, doneCountAverage)
	}

	err = pa.profileRepository.Upsert(userID, profileUpsertRequest.Name, profileUpsertRequest.University, profileUpsertRequest.Major, profileUpsertRequest.AspiringOccupation, profileUpsertRequest.AspiringField, profileUpsertRequest.Sentence, profileUpsertRequest.ImagePath, profileUpsertRequest.GraduationYear, profileUpsertRequest.JobHuntingStatus, deviationValue)
	if err != nil {
		return errors.Wrap(err, "profileApplication.ProfileUpsert()")
	}

	err = pa.profileRepository.InsertStorage(userID, profileUpsertRequest.Image)
	if err != nil {
		return errors.Wrap(err, "profileApplication.ProfileUpsert()")
	}

	return nil
}

func (pa profileApplication) ProfilesSearch(c echo.Context) error {
	// retriev profile data for searching user from http request.
	profilesSearchRequest := new(ProfilesSearchRequest)
	if err := c.Bind(profilesSearchRequest); err != nil {
		return errors.Wrap(err, "profileApplication.ProfilesSearch()")
	}

	// make slice which is used argument of SelectByProfileInfo()
	profileInfo := make(map[string]interface{})
	if profilesSearchRequest.Name != "" {
		profileInfo["name"] = profilesSearchRequest.Name
	}
	if profilesSearchRequest.University != "" {
		profileInfo["university"] = profilesSearchRequest.University
	}
	if profilesSearchRequest.GraduationYear != 0 {
		profileInfo["graduation_year"] = profilesSearchRequest.GraduationYear
	}
	if profilesSearchRequest.AspiringOccupation != "" {
		profileInfo["aspiring_occupation"] = profilesSearchRequest.AspiringOccupation
	}
	if profilesSearchRequest.AspiringField != "" {
		profileInfo["aspiring_field"] = profilesSearchRequest.AspiringField
	}

	// call function which select user who fill search condition in profile repository.
	profiles, err := pa.profileRepository.SelectByProfileInfo(profileInfo)
	if err != nil {
		return errors.Wrap(err, "profileApplication.ProfilesSearch()")
	}
	// set retrieved data in http response.
	c.JSON(http.StatusOK, profiles)
	return nil
}

// ProfileUpsertRequest : ProfileUpsert() request struct
type ProfileUpsertRequest struct {
	Name               string `form:"name"`
	University         string `form:"university"`
	Major              string `form:"major"`
	GraduationYear     int    `form:"graduation_year"`
	AspiringOccupation string `form:"aspiring_occupation"`
	AspiringField      string `form:"aspiring_field"`
	Sentence           string `form:"sentence"`
	ImagePath          string `form:"image_path"`
	Image              string `form:"image"` // not test
	JobHuntingStatus   int    `form:"job_hunting_status"`
}

// ProfilesSearchRequest : ProfilesSearch() request struct
type ProfilesSearchRequest struct {
	Name               string `form:"name"`
	University         string `form:"university"`
	GraduationYear     int    `form:"graduation_year"`
	AspiringOccupation string `form:"aspiring_occupation"`
	AspiringField      string `form:"aspiring_field"`
}
