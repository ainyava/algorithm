package main

import (
	"fmt"
)

func rob(nums []int) int {
	prev := 0
	max := 0
	for _, num := range nums {
		temp := max
		if prev+num > max {
			max = prev + num
		}
		prev = temp
	}
	return max
}

func main() {
	arr := []int{2, 7, 9, 3, 1}
	fmt.Println("Max money you can rob:", rob(arr))
}
