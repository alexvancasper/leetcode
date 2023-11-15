package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSpacest(t *testing.T) {
	type testData struct {
		name string
		in1  string
		in2  []int
		want string
	}

	tests := []testData{
		{name: "test-1", in1: "LeetcodeHelpsMeLearn", in2: []int{8, 13, 15}, want: "Leetcode Helps Me Learn"},
		{name: "test-2", in1: "icodeinpython", in2: []int{1, 5, 7, 9}, want: "i code in py thon"},
		{name: "test-spacing", in1: "spacing", in2: []int{0, 1, 2, 3, 4, 5, 6}, want: " s p a c i n g"},
		{name: "test-lol", in1: "lol", in2: []int{1, 2}, want: "l o l"},
		{name: "test-lol-zero", in1: "lol", in2: []int{0}, want: " lol"},
		{name: "test-lol-last", in1: "lol", in2: []int{3}, want: "lol "},
		{name: "test-lol-levery", in1: "lol", in2: []int{0, 1, 2, 3}, want: " l o l "},
	}
	t.Parallel()
	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			res := addSpaces(testCase.in1, testCase.in2)
			require.Equal(t, testCase.want, res)
		})
	}
}
