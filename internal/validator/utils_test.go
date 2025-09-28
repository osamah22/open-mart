package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInRange(t *testing.T) {
	min, max := 2, 10
	testCases := []struct {
		name  string
		value string
		want  bool
	}{
		{
			"valid",
			"something",
			true,
		},
		{
			"small",
			"s",
			false,
		},
		{
			"big",
			"something long",
			false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := inRange(tt.value, min, max)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	testCases := []struct {
		name  string
		value string
		want  bool
	}{
		{
			"empty",
			"  ",
			true,
		},

		{
			"not empty",
			"something",
			false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := isEmpty(tt.value)
			require.Equal(t, tt.want, res)
		})
	}
}
