package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	i := 0
	for i = range nums {
		j := i
		for ; j < len(nums) && nums[j] == nums[i]; j++ {
		}
		for k := 0; k+j < len(nums); k++ {
			nums[i+1+k] = nums[j+k]
		}
		if i > 0 && nums[i] <= nums[i-1] {
			i--
			break
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 2, 3, 3, 3, 4, 4, 4, 5, 6}
	k := removeDuplicates(nums)
	fmt.Println("Numbers: ", nums)
	fmt.Println("Uniques: ", k)

}
