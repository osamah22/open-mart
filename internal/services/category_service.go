package services

import (
	"context"
	"fmt"

	"github.com/osamah22/open-mart/internal/models"
)

var (
	ErrNoCategories      = fmt.Errorf("%w: no categories exists", NotFound)
	ErrDuplicateCategory = fmt.Errorf("%w: Category name already exists", NameTaken)
)

type CategoryService interface {
	ListCategories(ctx context.Context) (*[]models.Category, error)
	SlugExists(ctx context.Context, slug string) bool
}

type categoryService struct {
	queries *models.Queries
}

func NewCategoryService(queries *models.Queries) CategoryService {
	return &categoryService{
		queries: queries,
	}
}

func (s *categoryService) ListCategories(ctx context.Context) (*[]models.Category, error) {
	categories, err := s.queries.ListCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("ListCategories: %w", err)
	}

	if len(categories) == 0 {
		return nil, ErrNoCategories
	}

	return &categories, nil
}

func (s *categoryService) SlugExists(ctx context.Context, slug string) bool {
	exists, err := s.queries.CateogryExists(ctx, slug)
	if err != nil {
		return false
	}
	return exists
}
