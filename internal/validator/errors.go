package validator

import "fmt"

var (
	ErrEmptyString = fmt.Errorf("empty string")
	ErrNotInRange  = fmt.Errorf("not in range")
)
