package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			for j, ch := range needle {
				if i+j >= len(haystack) {
					break
				} else if rune(haystack[i+j]) != ch {
					break
				} else if j == len(needle)-1 {
					return i
				}
			}
		}
	}
	return -1
}

func main() {
	haystack, needle := "mississippi", "issipi"
	fmt.Println(needle, "index in", haystack, ":", strStr(haystack, needle))
}
