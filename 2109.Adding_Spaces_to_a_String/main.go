package main

import (
	"fmt"
	"strings"
)

func addSpaces(s string, spaces []int) string {
	var result strings.Builder
	result.Grow(len(s) + len(spaces))
	j, l, r := 0, 0, 0
	for j < len(spaces) {
		r = spaces[j]
		if l == r {
			result.WriteString(" ")
		} else {
			// result += string(s[l:r])
			// result += " "
			result.WriteString(fmt.Sprintf("%s ", s[l:r]))
		}
		l = r
		j++
	}
	result.WriteString(s[r:])
	return result.String()
}

func main() {
	fmt.Println(addSpaces("lol", []int{1, 2})) // l o l
}
