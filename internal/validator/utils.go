package validator

import "strings"

func inRange(value string, min, max int) bool {
	return len(value) >= min && len(value) <= max
}

func isEmpty(value string) bool {
	return strings.TrimSpace(value) == ""
}
