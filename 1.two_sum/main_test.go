package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	req := require.New(t)
	tests := map[string]struct {
		arr    []int
		target int
		want   []int
	}{
		// "pre-test-1": {arr: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		// "pre-test-2": {arr: []int{11, 15, 7, 1, 2}, target: 9, want: []int{2, 4}},
		// "pre-test-3": {arr: []int{-1, -3, -5, 7, 1}, target: 4, want: []int{1, 3}},
		// "pre-test-4": {arr: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		// "pre-test-5": {arr: []int{3, 3}, target: 6, want: []int{0, 1}},
		// "pre-test-6": {arr: []int{-5, 10, 0, 4, 4, 4, -4}, target: -9, want: []int{0, 6}},
		"pre-test-1": {arr: []int{2, 7, 11, 15}, target: 9, want: []int{2, 7}},
		"pre-test-2": {arr: []int{11, 15, 7, 1, 2}, target: 9, want: []int{7, 2}},
		"pre-test-3": {arr: []int{-1, -3, -5, 7, 1}, target: 4, want: []int{-3, 7}},
		"pre-test-4": {arr: []int{3, 2, 4}, target: 6, want: []int{2, 4}},
		"pre-test-5": {arr: []int{3, 3}, target: 6, want: []int{3, 3}},
		"pre-test-6": {arr: []int{-5, 10, 0, 4, 4, 4, -4}, target: -9, want: []int{-5, -4}},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := TwoSum(testCase.arr, testCase.target)
			req.Equal(testCase.want, res)
		})
	}
}

func TestTwoSumIdx(t *testing.T) {
	req := require.New(t)
	tests := map[string]struct {
		arr    []int
		target int
		want   []int
	}{
		"pre-test-1": {arr: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		"pre-test-2": {arr: []int{11, 15, 7, 1, 2}, target: 9, want: []int{2, 4}},
		"pre-test-3": {arr: []int{-1, -3, -5, 7, 1}, target: 4, want: []int{1, 3}},
		"pre-test-4": {arr: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		"pre-test-5": {arr: []int{3, 3}, target: 6, want: []int{0, 1}},
		"pre-test-6": {arr: []int{-5, 10, 0, 4, 4, 4, -4}, target: -9, want: []int{0, 6}},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := TwoSumIdx(testCase.arr, testCase.target)
			req.Equal(testCase.want, res)
		})
	}
}

func TestTwoSumIdx167(t *testing.T) {
	req := require.New(t)
	tests := map[string]struct {
		arr    []int
		target int
		want   []int
	}{
		"pre-test-1": {arr: []int{2, 7, 11, 15}, target: 9, want: []int{1, 2}},
		"pre-test-2": {arr: []int{11, 15, 7, 1, 2}, target: 9, want: []int{3, 5}},
		"pre-test-3": {arr: []int{-1, -3, -5, 7, 1}, target: 4, want: []int{2, 4}},
		"pre-test-4": {arr: []int{3, 2, 4}, target: 6, want: []int{2, 3}},
		"pre-test-5": {arr: []int{3, 3}, target: 6, want: []int{1, 2}},
		"pre-test-6": {arr: []int{-5, 10, 0, 4, 4, 4, -4}, target: -9, want: []int{1, 7}},
		"pre-test-7": {arr: []int{2, 3, 4}, target: 6, want: []int{1, 3}},
		"pre-test-8": {arr: []int{-1, 0}, target: -1, want: []int{1, 2}},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := TwoSumIdx167(testCase.arr, testCase.target)
			req.Equal(testCase.want, res)
		})
	}
}
