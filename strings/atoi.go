package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	n := 0
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	coef := 1
	if i < len(s) && s[i] == '-' {
		coef = -1
		i++
	} else if i < len(s) && s[i] == '+' {
		i++
	}
	for i < len(s) && (s[i] >= 48 && s[i] <= 57) {
		n *= 10
		n += int(s[i] - 48)
		i++
		if n > math.MaxInt32 {
			n = math.MaxInt32 + 1
			if coef == 1 {
				n -= 1
			}
			break
		}
	}

	return coef * n
}

func main() {
	s := "   -18446744073709551617 with words"
	fmt.Println("Integer: ", myAtoi(s))
}
