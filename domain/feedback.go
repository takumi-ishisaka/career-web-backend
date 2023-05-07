package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type FeedBack struct {
	FeedBackID   string    `json:"feedback_id"`
	UserActionID string    `json:"user_action_id"`
	Comment      string    `json:"comment"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func GenerateFeedBackID() (string, error) {
	feedbackID, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "domain.user.GenerateUserID()")
	}

	return feedbackID.String(), nil
}
