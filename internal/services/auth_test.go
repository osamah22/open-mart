package services

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/osamah22/open-mart/internal/models"
	"github.com/stretchr/testify/require"
)

// -------------------- HELPERS --------------------

func createUser(t *testing.T, svc AuthService) *models.User {
	user, err := svc.GetOrCreateUser(
		t.Context(),
		"google-123",
		"test@example.com",
		"https://avatar.url",
	)
	require.NoError(t, err)
	return user
}

// -------------------- TESTS --------------------

func TestGetOrCreateUser_CreatesUser(t *testing.T) {
	t.Parallel()
	pool, cleanup := models.SetupTestDB(t)
	defer cleanup()

	svc := NewService(models.New(pool))

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

func TestGetByGoogleID(t *testing.T) {
	t.Parallel()
	pool, cleanup := models.SetupTestDB(t)
	defer cleanup()

	svc := NewService(models.New(pool))
	user := createUser(t, svc)

	got, err := svc.GetByGoogleID(t.Context(), user.GoogleID)
	require.NoError(t, err)
	require.Equal(t, user.ID, got.ID)
}

func TestUpdateUsername(t *testing.T) {
	t.Parallel()
	pool, cleanup := models.SetupTestDB(t)
	defer cleanup()

	svc := NewService(models.New(pool))
	user := createUser(t, svc)

	now := time.Now().UTC()

	updated, err := svc.UpdateUsername(
		t.Context(),
		user.ID.String(),
		"new_username",
		now,
	)

	require.NoError(t, err)
	require.Equal(t, "new_username", updated.Username)
}

func TestUpdateUsername_Cooldown(t *testing.T) {
	t.Parallel()
	pool, cleanup := models.SetupTestDB(t)
	defer cleanup()

	q := models.New(pool)
	svc := NewService(q)
	user := createUser(t, svc)

	now := time.Now().UTC()

	_, err := q.UpdateUsername(t.Context(), models.UpdateUsernameParams{
		ID:       user.ID,
		Username: "first",
	})
	require.NoError(t, err)

	_, err = svc.UpdateUsername(
		t.Context(),
		user.ID.String(),
		"second",
		now.Add(5*24*time.Hour),
	)

	require.ErrorIs(t, err, ErrUsernameChangeCooldown)
}

func TestCanChangeUsernameLogic(t *testing.T) {
	t.Parallel()
	now := time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		user    *models.User
		wantErr error
	}{
		{
			name: "never changed",
			user: &models.User{
				LastUsernameChange: pgtype.Timestamp{Valid: false},
			},
			wantErr: nil,
		},
		{
			name: "5 days ago",
			user: &models.User{
				LastUsernameChange: pgtype.Timestamp{
					Time:  now.AddDate(0, 0, -5),
					Valid: true,
				},
			},
			wantErr: ErrUsernameChangeCooldown,
		},
		{
			name: "14 days ago",
			user: &models.User{
				LastUsernameChange: pgtype.Timestamp{
					Time:  now.AddDate(0, 0, -14),
					Valid: true,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CanChangeUsername(tt.user, now)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
