package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkTime(t *testing.T) {
	testCases := []struct {
		name     string
		values   string
		expected string
	}{
		{
			name:     "first",
			values:   "12:01:00AM",
			expected: "00:01:00",
		},
		{
			name:     "second",
			values:   "12:01:00PM",
			expected: "12:01:00",
		},
		{
			name:     "third",
			values:   "08:21:15PM",
			expected: "20:21:15",
		},
		{
			name:     "fourth",
			values:   "08:21:15AM",
			expected: "08:21:15",
		},
		{
			name:     "fifth",
			values:   "00:00:00AM",
			expected: "00:00:00",
		},
		{
			name:     "sixth",
			values:   "11:11:11PM",
			expected: "23:11:11",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			gotResp, gotErr := WorkTime(tc.values)
			assert.Equal(t, tc.expected, gotResp, gotErr)
		})
	}
}
