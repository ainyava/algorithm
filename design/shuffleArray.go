package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	original []int
	nums     []int
}

func Constructor(nums []int) Solution {
	original := make([]int, len(nums))
	copy(original, nums)
	return Solution{
		original,
		nums,
	}
}

func (sol *Solution) Reset() []int {
	sol.nums = make([]int, len(sol.original))
	copy(sol.nums, sol.original)
	return sol.nums
}

func (sol *Solution) Shuffle() []int {
	for i := range sol.nums {
		j := rand.Intn(i + 1)
		sol.nums[i], sol.nums[j] = sol.nums[j], sol.nums[i]
	}
	return sol.nums
}

func main() {
	nums := []int{1, 2, 3}
	obj := Constructor(nums)
	fmt.Println("Shuffled:", obj.Shuffle())
	fmt.Println("Resetted:", obj.Reset())
	fmt.Println("Shuffled:", obj.Shuffle())
	fmt.Println("Object:", obj)
}
