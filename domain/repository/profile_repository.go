package repository

import (
	"github.com/nokin-all-of-career/career-web-backend/domain"
)

// ProfileRepository : profile repository
type ProfileRepository interface {
	InsertDeviationValue(deviationValue float32, userID string) error
	SelectByPrimaryKey(userID string) (*domain.Profile, error)
	Upsert(userID, name, university, major, aspiringOccupation, aspiringField, sentence, imagePath string, graduationYear, jobHuntingStatus int, deviationvalue float32) error
	SelectByProfileInfo(profileInfo map[string]interface{}) (*[]domain.Profile, error)
	InsertStorage(userID, image string) error
}
