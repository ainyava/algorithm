package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	maxSum := math.MinInt32
	curSum := 0
	for _, num := range nums {
		curSum = max(num, num+curSum)
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

func main() {
	arr := []int{2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Max sub array sum:", maxSubArray(arr))
}
