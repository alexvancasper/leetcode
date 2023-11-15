package main

import (
	"fmt"
	"math"
)

func maximumSumScore(nums []int) int64 {
	pSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		pSum[i+1] = pSum[i] + nums[i]
	}
	var maxVal int64 = math.MinInt64
	for i := 0; i < len(nums); i++ {
		num1 := pSum[i+1]
		num2 := pSum[len(nums)] - pSum[i]
		if num1 >= num2 {
			if int64(num1) > maxVal {
				maxVal = int64(num1)
			}
		} else {
			if int64(num2) > maxVal {
				maxVal = int64(num2)
			}
		}
	}
	return maxVal
}

func max(num1, num2 int) int64 {
	if num1 > num2 {
		return int64(num1)
	}
	return int64(num2)
}

func main() {
	fmt.Println(maximumSumScore([]int{-100000, -100000, -100000, -100000}))
}
