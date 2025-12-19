package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/osamah22/open-mart/internal/models"
)

// CoolDownDuration is represented in days.
const CoolDownDuration = 14

func generateUsername() string {
	id := uuid.New()
	return "user_" + id.String()[:8]
}

func CanChangeUsername(user *models.User, now time.Time) error {
	if user.LastUsernameChange.Valid {
		cooldown := user.LastUsernameChange.Time.AddDate(0, 0, CoolDownDuration)
		if now.Before(cooldown) {
			return ErrUsernameChangeCooldown
		}
	}
	return nil
}
