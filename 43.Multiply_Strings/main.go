package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	stack := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			n1, _ := strconv.Atoi(string(num1[i]))
			n2, _ := strconv.Atoi(string(num2[j]))
			prevRem := stack[i+j+1]
			m := n1*n2 + prevRem
			digit := m % 10
			carry := math.Floor(float64(m) / 10)
			stack[i+j+1] = digit
			stack[i+j] += int(carry)
		}
	}

	var result strings.Builder
	result.Grow(len(stack))
	var k int
	for i := 0; i < len(stack); i++ {
		k = i
		if stack[i] != 0 {
			break
		}
	}
	stack = stack[k:]

	for i := 0; i < len(stack); i++ {
		result.WriteString(strconv.Itoa(stack[i]))
	}
	return result.String()
}

func main() {
	fmt.Println(multiply("9133", "0"))
}
