package services

import (
	"errors"
)

var (
	NotFound  = errors.New("Not Found")
	NameTaken = errors.New("Duplicate name")
)
