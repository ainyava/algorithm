package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	// define pointers to bigger and smaller arrays
	var (
		big   *[]int
		small *[]int
	)
	if len(nums1) > len(nums2) {
		big = &nums1
		small = &nums2
	} else {
		big = &nums2
		small = &nums1
	}
	var (
		bc [1001]int
		sc [1001]int
	)
	var i int
	for i = 0; i < len(*small); i++ {
		bc[(*big)[i]]++
		sc[(*small)[i]]++
	}
	for ; i < len(*big); i++ {
		bc[(*big)[i]]++
	}
	var ans []int
	for i = 0; i < len(bc); i++ {
		var m int
		if sc[i] > bc[i] {
			m = bc[i]
		} else {
			m = sc[i]
		}
		if m > 0 {
			for j := 0; j < m; j++ {
				ans = append(ans, i)
			}
		}
	}
	return ans
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println("Intersection: ", intersect(nums1, nums2))
}
