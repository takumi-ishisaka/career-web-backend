package persistence

import (
	"bytes"
	"database/sql"
	"strings"

	"github.com/pkg/errors"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/domain/repository"
	"github.com/nokin-all-of-career/career-web-backend/infra"
)

type profilePersistence struct {
	DB *sql.DB
}

// NewProfilePersistence : create profile persistence
func NewProfilePersistence(DB *sql.DB) repository.ProfileRepository {
	return &profilePersistence{
		DB: DB,
	}
}

// InsertDeviationValue: insert deviationValue in profile table
func (pp profilePersistence) InsertDeviationValue(deviationValue float32, userID string) error {
	stmt, err := pp.DB.Prepare("UPDATE profile SET deviation_value=? WHERE user_id=?")
	if err != nil {
		return errors.Wrap(err, "profilePersistence.InsertDeviationValue()")
	}
	_, err = stmt.Exec(deviationValue, userID)
	if err != nil {
		return errors.Wrap(err, "profilePersistence.InsertDeviationValue()")
	}

	return nil
}

// SelectByPrimaryKey : get user Profile by userID
func (pp profilePersistence) SelectByPrimaryKey(userID string) (*domain.Profile, error) {
	row := pp.DB.QueryRow("SELECT * FROM profile WHERE user_id=?", userID)
	profile, err := convertToProfile(row)
	if err != nil {
		return nil, errors.Wrap(err, "profilePersistence.SelectByPrimaryKey()")
	}

	return profile, nil
}

func convertToProfile(row *sql.Row) (*domain.Profile, error) {
	profile := domain.Profile{}
	err := row.Scan(&profile.UserID, &profile.Name, &profile.University, &profile.Major, &profile.GraduationYear, &profile.AspiringOccupation, &profile.AspiringField, &profile.Sentence, &profile.ImagePath, &profile.JobHuntingStatus, &profile.DeviationValue)
	if err != nil {
		if err == sql.ErrNoRows {
			return &profile, nil
		}
		return nil, errors.Wrap(err, "profilePersistence.convertToProfile()")
	}

	return &profile, nil
}

// Upsert: create or update Profile using information enterted by user
func (pp profilePersistence) Upsert(userID, name, university, major, aspiringOccupation, aspiringField, sentence, imagePath string, graduationYear, jobHuntingStatus int, deviationValue float32) error {
	stmt, err := pp.DB.Prepare("INSERT INTO profile(user_id, name, university, major, graduation_year, aspiring_occupation, aspiring_field, sentence, image_path, job_hunting_status, deviation_value) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE key UPDATE name=?, graduation_year=?, aspiring_occupation=?, aspiring_field=?, sentence=?, image_path=?, job_hunting_status=?, deviation_value=?")
	if err != nil {
		return errors.Wrap(err, "profilePersistence.Upsert()")
	}
	_, err = stmt.Exec(userID, name, university, major, graduationYear, aspiringOccupation, aspiringField, sentence, imagePath, jobHuntingStatus, deviationValue, name, graduationYear, aspiringOccupation, aspiringField, sentence, imagePath, jobHuntingStatus, deviationValue)
	if err != nil {
		return errors.Wrap(err, "profilePersistence.Upsert()")
	}

	return nil
}

func (pp profilePersistence) SelectByProfileInfo(profileInfo map[string]interface{}) (*[]domain.Profile, error) {
	query := "SELECT * FROM profile WHERE "
	var info []interface{}
	for k, v := range profileInfo {
		query += k + "=? AND "
		info = append(info, v)
	}
	query = strings.TrimRight(query, "AND ")

	rows, err := pp.DB.Query(query, info...)
	if err != nil {
		return nil, errors.Wrap(err, "profilePersistence.SelectByProfileInfo()")
	}

	var profiles []domain.Profile
	for rows.Next() {
		var profile domain.Profile
		err := rows.Scan(&profile.UserID, &profile.Name, &profile.University, &profile.Major, &profile.GraduationYear, &profile.AspiringOccupation, &profile.AspiringField, &profile.Sentence, &profile.ImagePath, &profile.JobHuntingStatus, &profile.DeviationValue)
		if err != nil {
			return &profiles, err
		}
		profiles = append(profiles, profile)
	}

	return &profiles, nil
}

// InsertStorage : save profile image to gcp storage
func (pp profilePersistence) InsertStorage(userID, image string) error {
	// upload profile image
	bucketPath := userID + "/profile.txt"
	uploadObject := infra.Bucket.Object(bucketPath)

	writer := uploadObject.NewWriter(infra.CtxStorage)
	defer writer.Close()

	buff := new(bytes.Buffer)
	_, err := buff.WriteTo(writer) // exec upload
	if err != nil {
		errors.Wrap(err, "profilePersistence.InsertStorage()")
	}

	return err
}
