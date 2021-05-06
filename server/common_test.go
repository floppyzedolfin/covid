package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareDates(t *testing.T) {
	tt := map[string]struct {
		lhs      string
		rhs      string
		expected bool
	}{
		"nominal": {
			lhs:      "2020-05-08",
			rhs:      "2020-05-12",
			expected: true,
		},
		"lanimon": {
			lhs:      "2020-05-12",
			rhs:      "2020-05-08",
			expected: false,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			res := dateBefore(tc.lhs, tc.rhs)
			assert.Equal(t, tc.expected, res)
		})
	}
}
