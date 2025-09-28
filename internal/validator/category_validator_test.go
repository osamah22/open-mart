package validator_test

import (
	"errors"
	"testing"

	"github.com/osamah22/open-mart/internal/models"
	"github.com/osamah22/open-mart/internal/validator"
	"github.com/stretchr/testify/require"
)

func TestValidateCreateCategory(t *testing.T) {
	tests := []struct {
		name    string
		req     *models.CreateCategoryParams
		wantErr error
	}{
		{
			name:    "valid name",
			req:     &models.CreateCategoryParams{Name: "Groceries"},
			wantErr: nil,
		},
		{
			name:    "empty name",
			req:     &models.CreateCategoryParams{Name: ""},
			wantErr: validator.ErrCategoryNameEmpty,
		},
		{
			name:    "whitespace only name",
			req:     &models.CreateCategoryParams{Name: "   "},
			wantErr: validator.ErrCategoryNameEmpty,
		},
		{
			name:    "less characters",
			req:     &models.CreateCategoryParams{Name: "e"},
			wantErr: validator.ErrNotInRange,
		},
		{
			name:    "more characters",
			req:     &models.CreateCategoryParams{Name: string(make([]byte, 260))},
			wantErr: validator.ErrNotInRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateCreateCategory(tt.req)
			require.True(t, errors.Is(err, tt.wantErr))
		})
	}
}

func TestValidateUpdateCategory(t *testing.T) {
	tests := []struct {
		name    string
		req     *models.UpdateCategoryParams
		wantErr error
	}{
		{
			name:    "valid name",
			req:     &models.UpdateCategoryParams{Name: "Electronics"},
			wantErr: nil,
		},
		{
			name:    "empty name",
			req:     &models.UpdateCategoryParams{Name: ""},
			wantErr: validator.ErrCategoryNameEmpty,
		},
		{
			name:    "whitespace only name",
			req:     &models.UpdateCategoryParams{Name: "   "},
			wantErr: validator.ErrCategoryNameEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateUpdateCategory(tt.req)
			// require.Equal(t, tt.wantErr, err)

			require.True(t, errors.Is(err, tt.wantErr))
		})
	}
}
