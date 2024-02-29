package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	last := len(digits) - 1
	digits[last]++
	for i := last; i >= 0; i-- {
		if digits[i] != 10 {
			break
		}

		digits[i] = 0
		if i != 0 {
			digits[i-1] += 1
		} else {
			return append([]int{1}, digits...)
		}
	}
	return digits
}

func main() {
	nums := []int{4, 3, 2, 1}
	fmt.Println("PlusOne: ", plusOne(nums))
}
