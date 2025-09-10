package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected int64
	}{
		{"Empty file", "./testdata/yarn.lock", 0},
		{"Non-empty file", "./testdata", 38},
		{"Empty folder", "./testdata/lib", 0},
		{"Non-empty folder", "./testdata/doc", 13663},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetSize(tc.path)
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}
