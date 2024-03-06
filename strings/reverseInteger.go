package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	coefficient := 1
	if x < 0 {
		x = -x
		coefficient = -1
	}
	n := 0
	for x > 0 {
		n *= 10
		n += x % 10
		x /= 10
	}
	if n > math.MaxInt32 {
		return 0
	}
	return coefficient * n
}

func main() {
	x := -123
	fmt.Println("Reversed: ", reverse(x))
}
