package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/osamah22/open-mart/internal/models"
)

func TestCreateCategory(t *testing.T) {
	testDB := models.NewTestingDatabase(t)
	queries := models.New(testDB)
	s := NewCategoryService(queries)
	testTables := []struct {
		TestName string
		Name     string
		ParentID uuid.NullUUID
	}{
		{
			TestName: "Valid",
			Name:     "category",
		},
		{
			TestName: "duplicate",
			Name:     "duplicate",
		},
	}

	for _, test := range testTables {
		t.Run(test.TestName, func(t *testing.T) {
			params := models.CreateCategoryParams{
				Name:     test.TestName,
				ParentID: test.ParentID,
			}
			s.CreateCategory(t.Context(), params)
		})
	}
}
