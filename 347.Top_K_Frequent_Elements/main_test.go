package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTopKFrequent(t *testing.T) {
	req := require.New(t)

	tests := map[string]struct {
		in1  []int
		in2  int
		want []int
	}{
		"pre-test-1": {in1: []int{1, 1, 1, 2, 2, 3}, in2: 2, want: []int{1, 2}},
		"pre-test-2": {in1: []int{2, 3, 2, 1, 1, 1}, in2: 2, want: []int{1, 2}},
		"pre-test-3": {in1: []int{1, 3, 2, 1, 2, 1}, in2: 2, want: []int{1, 2}},
		"pre-test-4": {in1: []int{1}, in2: 1, want: []int{1}},
		"pre-test-5": {in1: []int{1, 9, 3, 2}, in2: 3, want: []int{1, 2, 3}},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := topKFrequent(testCase.in1, testCase.in2)
			req.Equal(testCase.want, res)
		})
	}
}
