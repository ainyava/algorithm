package main

import (
	"fmt"
	"slices"
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
	decimals := make(map[int]int, 4)
	for i := len(s) - 1; i >= 0; i-- {
		if i > 0 {
			if slices.Contains([]byte{'V', 'X'}, s[i]) && s[i-1] == 'I' {
				decimals[0] += mappings[s[i]] - 1
				i--
				continue
			} else if slices.Contains([]byte{'L', 'C'}, s[i]) && s[i-1] == 'X' {
				decimals[1] += mappings[s[i]] - 10
				i--
				continue
			} else if slices.Contains([]byte{'D', 'M'}, s[i]) && s[i-1] == 'C' {
				decimals[1] += mappings[s[i]] - 100
				i--
				continue
			}
		}
		n := mappings[s[i]]
		if n < 10 {
			decimals[0] += n
		} else if n < 100 {
			decimals[1] += n
		} else if n < 1000 {
			decimals[2] += n
		} else if n < 10000 {
			decimals[3] += n
		}
	}
	return decimals[3] + decimals[2] + decimals[1] + decimals[0]
}

func main() {
	n := "MCMXCIV"
	fmt.Println("Roman to integer:", romanToInt(n))
}
