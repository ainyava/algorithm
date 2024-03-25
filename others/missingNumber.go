package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums) + 1
	sum := n * (n - 1) / 2
	for _, i := range nums {
		sum -= i
	}
	return sum
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println("Missing number:", missingNumber(nums))
}
