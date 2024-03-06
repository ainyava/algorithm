package main

import (
	"fmt"
)

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func main() {
	s := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	fmt.Printf("%s ", s)
	reverseString(s)
	fmt.Printf("reversed: %s", s)
}
