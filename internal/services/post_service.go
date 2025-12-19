package services

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/osamah22/open-mart/internal/models"
)

var (
	ErrInvalidTitle   = errors.New("Invalid title")
	ErrInvalidContent = errors.New("Invalid title")
)

type PostService interface {
	CreatePost(ctx context.Context, req models.CreatePostParams) (*models.Post, error)
	UpdatePost(ctx context.Context, req models.UpdatePostParams) (*models.Post, error)
	DeletePost(ctx context.Context, id uuid.UUID) error
}
type postService struct {
	q *models.Queries
}

func NewPostService(queries *models.Queries) PostService {
	return &postService{q: queries}
}

func (s *postService) CreatePost(ctx context.Context, req models.CreatePostParams) (*models.Post, error) {
	if strings.TrimSpace(req.Title) == "" {
		return nil, ErrInvalidTitle
	}
	post, err := s.q.CreatePost(ctx, req)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *postService) UpdatePost(ctx context.Context, req models.UpdatePostParams) (*models.Post, error) {
	if strings.TrimSpace(req.Title) == "" {
		return nil, ErrInvalidTitle
	}
	return nil, nil
}

func (s *postService) DeletePost(ctx context.Context, id uuid.UUID) error
