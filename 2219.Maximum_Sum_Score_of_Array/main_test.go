package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaximumSumScore(t *testing.T) {
	type testData struct {
		name string
		in1  []int
		want int64
	}

	tests := []testData{
		{name: "test-1", in1: []int{4, 3, -2, 5}, want: 10},
		{name: "test-2", in1: []int{-3, -5}, want: -3},
		{name: "test-100000", in1: []int{-100000, -100000, -100000, -100000}, want: -100000},
		// {name: "test-lol", in1: []int{1, 2}, want: []int{}},
		// {name: "test-lol-zero", in1: []int{0}, want: []int{}},
		// {name: "test-lol-last", in1: []int{3}, want: []int{}},
		// {name: "test-lol-levery", in1: []int{0, 1, 2, 3}, want: []int{}},
	}
	t.Parallel()
	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			res := maximumSumScore(testCase.in1)
			require.Equal(t, testCase.want, res)
		})
	}
}
