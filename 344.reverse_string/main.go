package main

import "fmt"

func main() {
	s := "hello"
	bs := []byte(s)
	reverseString(bs)
}

func reverseString(s []byte) {
	// this works
	// for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
	// 	s[i], s[j] = s[j], s[i]
	// }

	// this works too
	for i := 0; i < (len(s) / 2); i++ {
		s[i], s[(len(s)-1)-i] = s[(len(s)-1)-i], s[i]
	}

	fmt.Printf("%s\n", string(s))
}
