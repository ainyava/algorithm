package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	exists := make(map[int]int)
	for i, n := range nums {
		j := exists[target-n]
		if j > 0 {
			return []int{i, j - 1}
		}
		exists[n] = (i + 1)
	}
	return []int{}
}

func main() {
	nums, target := []int{3, 2, 4}, 6
	fmt.Println("Numerse: ", nums)
	fmt.Println("TwoSums of: ", twoSum(nums, target), "equal to", target)
}
