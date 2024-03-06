package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make(map[byte]int)
	for i := range s {
		counts[s[i]]++
		counts[t[i]]--
	}
	for ch := range counts {
		if counts[byte(ch)] != 0 {
			return false
		}
	}
	return true
}
func main() {
	s, t := "a", "ab"
	fmt.Println("Anagram?", isAnagram(s, t))
}
