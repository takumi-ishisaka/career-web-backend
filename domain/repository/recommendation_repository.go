package repository

import "github.com/nokin-all-of-career/career-web-backend/domain"

//Recommendation is the interface to recommendation struct meat
type RecommendationRepository interface {
	UpsertRecommendation(actionID string, userRecomendationValue int) error
	SelectByPrimaryKey(actionID string) (*domain.Recommendation, error)
	SelectByPrimaryKeys(actionIDs []interface{}) ([]domain.Recommendation, error)
}
