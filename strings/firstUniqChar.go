package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	counts := make(map[rune]int)
	for _, ch := range s {
		counts[ch]++
	}
	for i, ch := range s {
		if counts[ch] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	s := "aabb"
	fmt.Println("Index of first non-repeating character: ", firstUniqChar(s))
}
