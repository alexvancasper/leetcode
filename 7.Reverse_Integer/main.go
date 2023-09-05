package main

import (
	"fmt"
)

func reverse(x int) int {
	rightLimit := 2147483648
	leftLimit := -2147483648
	var result, remain int
	if leftLimit-x == 0 || rightLimit+x == 0 {
		return 0
	}
	for x != 0 {
		remain = x % 10
		result = result * 10
		result = result + remain
		x = x / 10
	}
	if result > rightLimit || result < leftLimit {
		return 0
	}
	return result
}

func main() {
	// a := 2147483648
	a := -2147483648
	res := reverse(a)
	fmt.Printf("%d\n", res)

}
