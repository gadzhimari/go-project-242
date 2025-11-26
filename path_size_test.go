package code

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	testCases := []struct {
		name      string
		path      string
		recursive bool
		all       bool
		expected  int64
	}{
		{"Empty file", "./testdata/yarn.lock", false, false, 0},
		{"Non-empty file", "./testdata", false, false, 38},
		{"Empty folder", "./testdata/lib", false, false, 0},
		{"Non-empty folder", "./testdata/doc", false, false, 13663},
		{"Folder without hidden files", "./testdata/examples", false, false, 3345},
		{"Folder with hidden files", "./testdata/examples", false, true, 9493},
		{"Recursive without hidden files", "./testdata", true, false, 2511882},
		{"Recursive with hidden files", "./testdata", true, true, 2518391},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetSize(tc.path, tc.recursive, tc.all)
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}

func TestFormatSize(t *testing.T) {
	testCases := []struct {
		name     string
		size     int64
		human    bool
		expected string
	}{
		{"Non-human: zero", 0, false, "0B"},
		{"Non-human: large", 999999, false, "999999B"},
		{"Human: 1024B = 1KB", 1024, true, "1.0KB"},
		{"Human: 1536B = 1.5KB", 1536, true, "1.5KB"},
		{"Human: 2048B = 2KB", 2048, true, "2.0KB"},
		{"Human: 1025B ~ 1.0KB", 1025, true, "1.0KB"},
		{"Human: 1.5MB", 3 * 1024 * 512, true, "1.5MB"},
		{"Human: 1.7GB", int64(math.Round(1.7 * 1024 * 1024 * 1024)), true, "1.7GB"},
		{"Human: 3.2TB", int64(math.Round(3.2 * 1024 * 1024 * 1024 * 1024)), true, "3.2TB"},
		{"Human: 1PB exact", 1024 * 1024 * 1024 * 1024 * 1024, true, "1.0PB"},
		{"Human: 2.1EB", int64(math.Round(2.1 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024)), true, "2.1EB"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := formatSize(tc.size, tc.human)
			require.Equal(t, tc.expected, got)
		})
	}
}
