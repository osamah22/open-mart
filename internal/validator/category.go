package validator

import (
	"fmt"
	"strings"

	"github.com/osamah22/open-mart/internal/models"
)

var ErrCategoryNameEmpty = fmt.Errorf("%w: category name cannot be empty", ErrEmptyString)

func ValidateCreateCategory(req *models.CreateCategoryParams) error {
	if strings.TrimSpace(req.Name) == "" {
		return ErrCategoryNameEmpty
	}
	min, max := 2, 256
	if !inRange(req.Name, min, max) {
		return fmt.Errorf("%w: category name should be in range(%d, %d)", ErrNotInRange, min, max)
	}
	return nil
}

func ValidateUpdateCategory(req *models.UpdateCategoryParams) error {
	if strings.TrimSpace(req.Name) == "" {
		return ErrCategoryNameEmpty
	}
	min, max := 2, 256
	if !inRange(req.Name, min, max) {
		return fmt.Errorf("%w: category name should be in range(%d, %d)", ErrNotInRange, min, max)
	}
	return nil
}
