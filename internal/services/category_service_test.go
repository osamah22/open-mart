package services

import (
	"testing"

	"github.com/osamah22/open-mart/internal/models"
	"github.com/stretchr/testify/require"
)

func TestGetCategory(t *testing.T) {
	db, cleanup := models.SetupTestDB(t)
	defer cleanup()

	queries := models.New(db)
	s := NewCategoryService(queries)

	categories, err := s.ListCategories(t.Context())
	require.NoError(t, err)
	require.NotEmpty(t, categories)
}
