package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/osamah22/open-mart/internal/models"
	"github.com/osamah22/open-mart/internal/validator"
)

var (
	ErrNoCategories      = fmt.Errorf("%w: no categories exists", NotFound)
	ErrDuplicateCategory = fmt.Errorf("%w: Category name already exists", NameTaken)
)

type CategoryService interface {
	ListCategories(ctx context.Context) (*[]models.Category, error)
	CreateCategory(ctx context.Context, req models.CreateCategoryParams) (*models.Category, error)
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
		return nil, err
	}

	return &categories, nil
}

func (s *categoryService) CreateCategory(ctx context.Context, req models.CreateCategoryParams) (*models.Category, error) {
	req.Name = strings.TrimSpace(req.Name)
	err := validator.ValidateCreateCategory(&req)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, ErrDuplicateCategory
		}
		return nil, err
	}

	category, err := s.queries.CreateCategory(ctx, req)
	if err != nil {
		return nil, err
	}
	return &category, err
}

func (s *categoryService) UpdateCategory(ctx context.Context, req models.UpdateCategoryParams) (*models.Category, error) {
	req.Name = strings.TrimSpace(req.Name)
	err := validator.ValidateUpdateCategory(&req)
	if err != nil {
		return nil, err
	}

	category, err := s.queries.UpdateCategory(ctx, req)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, ErrDuplicateCategory
		}
		return nil, err
	}
	return &category, err
}
