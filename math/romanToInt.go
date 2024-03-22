package main

import (
	"fmt"
)

func romanToInt(s string) int {
	mappings := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i := range s {
		if i < len(s)-1 && mappings[s[i]] < mappings[s[i+1]] {
			res -= mappings[s[i]]
		} else {
			res += mappings[s[i]]
		}
	}
	return res
}

func main() {
	n := "MCMXCIV"
	fmt.Println("Roman to integer:", romanToInt(n))
}
