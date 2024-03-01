package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums := []int{1, 2, 3, 3, 3, 4, 4, 4, 5, 6}
	k := removeDuplicates(nums)
	fmt.Println("Numbers: ", nums)
	fmt.Println("Uniques: ", k)

}
