package services

import (
	"context"
	"errors"
	"log"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/osamah22/open-mart/internal/models"
)

var (
	ErrUserNotFound           = errors.New("user does not exist")
	ErrInvalidUsername        = errors.New("invalid username")
	ErrUsernameChangeCooldown = errors.New("username can only be updated every 14 days")

	usernameRex = regexp.MustCompile(`^[a-z0-9_]+$`)
)

type AuthService interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByGoogleID(ctx context.Context, googleID string) (*models.User, error)
	GetOrCreateUser(ctx context.Context, googleID, email, avatarURL string) (*models.User, error)
	UpdateUsername(ctx context.Context, userID, username string, now time.Time) (*models.User, error)
}

type authService struct {
	q      *models.Queries
	logger *log.Logger
}

func NewService(q *models.Queries) AuthService {
	return &authService{q: q}
}

// -------------------- GET BY ID --------------------

func (s *authService) GetByID(ctx context.Context, id string) (*models.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	user, err := s.q.GetUserByID(ctx, pgtype.UUID{Bytes: uid, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// -------------------- GET BY GOOGLE ID --------------------

func (s *authService) GetByGoogleID(ctx context.Context, googleID string) (*models.User, error) {
	user, err := s.q.GetUserByGoogleID(ctx, googleID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// -------------------- GET OR CREATE --------------------

func (s *authService) GetOrCreateUser(
	ctx context.Context,
	googleID,
	email,
	avatarURL string,
) (*models.User, error) {
	user, err := s.q.GetUserByGoogleID(ctx, googleID)
	if err == nil {
		_, _ = s.q.UpdateUserAvatar(ctx, models.UpdateUserAvatarParams{
			ID: user.ID,
			AvatarUrl: pgtype.Text{
				String: avatarURL,
				Valid:  avatarURL != "",
			},
		})
		return &user, nil
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	user, err = s.q.CreateUser(ctx, models.CreateUserParams{
		GoogleID: googleID,
		Email:    email,
		Username: generateUsername(),
		AvatarUrl: pgtype.Text{
			String: avatarURL,
			Valid:  avatarURL != "",
		},
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// -------------------- UPDATE USERNAME --------------------

func (s *authService) UpdateUsername(
	ctx context.Context,
	userID,
	username string,
	now time.Time,
) (*models.User, error) {
	if len(username) < 3 {
		return nil, ErrInvalidUsername
	}
	if !usernameRex.MatchString(username) {
		return nil, ErrInvalidUsername
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	user, err := s.q.GetUserByID(ctx, pgtype.UUID{Bytes: uid, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if user.LastUsernameChange.Valid {
		cooldown := user.LastUsernameChange.Time.AddDate(0, 0, 14)
		if now.Before(cooldown) {
			return nil, ErrUsernameChangeCooldown
		}
	}

	user, err = s.q.UpdateUsername(ctx, models.UpdateUsernameParams{
		ID:       user.ID,
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}
