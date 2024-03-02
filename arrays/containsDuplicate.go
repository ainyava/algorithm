package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	contain := make(map[int]bool)
	for _, num := range nums {
		if contain[num] {
			return true
		}
		contain[num] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("Contains duplicate?", containsDuplicate(nums))
}
