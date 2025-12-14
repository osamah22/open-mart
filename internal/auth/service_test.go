package auth

import (
	"testing"

	"github.com/osamah22/open-mart/internal/models"
	"github.com/stretchr/testify/require"
)

func TestGetOrCreateUser_CreatesUser(t *testing.T) {
	pool, cleanup := models.SetupTestDB(t)
	defer cleanup()
	q := models.New(pool)
	svc := NewService(q)

	user, err := svc.GetOrCreateUser(
		t.Context(),
		"google-123",
		"test@example.com",
		"https://avatar.url",
	)

	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "google-123", user.GoogleID)
	require.NotEmpty(t, user.Username)
}

func TestGetOrCreateUser_GetUserByGoogleId(t *testing.T) {
	pool, cleanup := models.SetupTestDB(t)
	defer cleanup()
	q := models.New(pool)
	svc := NewService(q)

	user := createUserHelper(t, svc)
	tests := []struct {
		name     string
		googleId string
		wantErr  error
	}{
		{
			"Valid id",
			user.ID.String(),
			nil,
		},

		{
			"Not existing id",
			user.ID.String(),
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.GetByGoogleID(t.Context(), tt.googleId)
			require.Equal(t, tt.wantErr, err)
		})
	}
}

// func TestGetByUserId_Get
func createUserHelper(t *testing.T, svc AuthService) *models.User {
	user, err := svc.GetOrCreateUser(
		t.Context(),
		"google-123",
		"test@example.com",
		"https://avatar.url",
	)

	require.NoError(t, err)
	require.NotNil(t, user)
	require.NotEmpty(t, user)
	return user
}
