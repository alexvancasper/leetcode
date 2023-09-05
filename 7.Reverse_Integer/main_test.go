package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverse(t *testing.T) {
	req := require.New(t)

	tests := map[string]struct {
		x    int
		want int
	}{
		"pre-test-1": {x: 123, want: 321},
		"pre-test-2": {x: -123, want: -321},
		"pre-test-3": {x: 120, want: 21},
		"pre-test-4": {x: 2147483648, want: 0},
		"pre-test-5": {x: -2147483648, want: 0},
		"pre-test-6": {x: 0, want: 0},
		"pre-test-7": {x: 9999999999, want: 0},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := reverse(testCase.x)
			req.Equal(testCase.want, res)
		})
	}
}
