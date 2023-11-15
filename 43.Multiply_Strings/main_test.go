package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiply(t *testing.T) {
	type testData struct {
		name string
		in1  string
		in2  string
		want string
	}

	tests := []testData{
		{name: "test-1", in1: "2", in2: "6", want: "12"},
		{name: "test-zero", in1: "0", in2: "1", want: "0"},
		{name: "test-123", in1: "123", in2: "456", want: "56088"},
		{name: "test-999", in1: "999", in2: "999", want: "998001"},
		{name: "test-9133", in1: "9133", in2: "0", want: "0"},
		{name: "test-big", in1: "123456789", in2: "987654321", want: "121932631112635269"},
	}
	t.Parallel()
	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			res := multiply(testCase.in1, testCase.in2)
			require.Equal(t, testCase.want, res)
		})
	}
}
