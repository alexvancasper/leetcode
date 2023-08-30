package main

import "fmt"

func main() {
	// nums := []int{2, 7, 11, 15}
	nums := []int{-1, -3, -5, 7, 1}
	target := 4
	result := TwoSum(nums, target)
	fmt.Printf("%v", result)
}

func TwoSum(nums []int, target int) []int {
	var diff int
	v := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff = target - nums[i]
		if _, ok := v[diff]; ok {
			return []int{diff, nums[i]}
		} else {
			v[nums[i]] = diff
		}
	}
	return []int{}
}

func TwoSumIdx(nums []int, target int) []int {
	var diff int
	v := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff = target - nums[i]
		if _, ok := v[diff]; ok {
			return []int{v[diff], i}
		} else {
			v[nums[i]] = i
		}
	}
	return []int{}
}

func TwoSumIdx167(nums []int, target int) []int {
	var diff int
	v := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff = target - nums[i]
		if _, ok := v[diff]; ok {
			return []int{v[diff], i + 1}
		} else {
			v[nums[i]] = i + 1
		}
	}
	return []int{}
}
