package auth

import (
	"github.com/google/uuid"
)

func generateUsername() string {
	id := uuid.New()
	return "user_" + id.String()[:8]
}
