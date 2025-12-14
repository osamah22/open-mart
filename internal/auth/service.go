package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/osamah22/open-mart/internal/models"
)

var (
	ErrUserNotFount    = errors.New("User does not exist")
	ErrInvalidUsername = errors.New("Invalid username")
	usernameRex        = regexp.MustCompile(`(?i)^[a-z0-9_]+$`)
	emailRex           = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	phoneNumberRex     = regexp.MustCompile(`^\+?[0-9]{8,15}$`)
)

type AuthService interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByGoogleID(ctx context.Context, googleID string) (*models.User, error)

	GetOrCreateUser(
		ctx context.Context,
		googleID string,
		email string,
		avatarURL string,
	) (*models.User, error)

	UpdateUsername(
		ctx context.Context,
		userID uuid.UUID,
		username string,
	) (*models.User, error)
}

type authService struct {
	q      *models.Queries
	logger *log.Logger
}

func NewService(queries *models.Queries) AuthService {
	return &authService{
		q: queries,
	}
}

func (u *authService) GetByID(ctx context.Context, id string) (*models.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user id")
	}
	user, err := u.q.GetUserByID(ctx, pgtype.UUID{Bytes: uid, Valid: true})
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, ErrUserNotFount
		}
		return nil, fmt.Errorf("invalid user id")
	}

	return &user, nil
}

func (u *authService) GetByGoogleID(ctx context.Context, googleID string) (*models.User, error) {
	return nil, nil
}

func (u *authService) GetOrCreateUser(
	ctx context.Context,
	googleID string,
	email string,
	avatarURL string,
) (*models.User, error) {
	user, err := u.q.GetUserByGoogleID(ctx, googleID)

	// ✅ user exists
	if err == nil {
		_ = u.q.UpdateUserAvatar(ctx, models.UpdateUserAvatarParams{
			ID: user.ID,
			AvatarUrl: pgtype.Text{
				String: avatarURL,
				Valid:  avatarURL != "",
			},
		})
		return &user, nil
	}

	// ✅ user does NOT exist → create
	if errors.Is(err, pgx.ErrNoRows) {
		username := generateUsername()

		user, err := u.q.CreateUser(ctx, models.CreateUserParams{
			GoogleID: googleID,
			Email:    email,
			Username: username,
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

	// ❌ real DB error
	return nil, err
}

func (u *authService) UpdateUsername(
	ctx context.Context,
	userID uuid.UUID,
	username string,
) (*models.User, error) {
	if len(username) > 3 {
		return nil, fmt.Errorf("%w: username is too short", ErrInvalidUsername)
	} else if !usernameRex.MatchString(username) {
		if username[0] == '.' || username[len(username)-1] == '.' {
			return nil, fmt.Errorf("%w: username cannot start or end with '.'", ErrInvalidUsername)
		}
		return nil, fmt.Errorf("%w: username should contains only letters, underscores or periods", ErrInvalidUsername)
	} else if _, ok := reservedUsernames[username]; ok {
		return nil, fmt.Errorf("%w: username not allowed", ErrInvalidUsername)
	}

	pgId := pgtype.UUID{Bytes: userID, Valid: true}
	user, err := u.q.UpdateUsername(ctx, models.UpdateUsernameParams{
		ID:       pgId,
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}
