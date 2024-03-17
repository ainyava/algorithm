package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	n1 := nums1[:m]
	i, j, index := m-1, n-1, m+n-1

	for ; i >= 0 && j >= 0; index-- {
		if n1[i] > nums2[j] {
			nums1[index] = n1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
	}
	for i >= 0 {
		nums1[index] = n1[i]
		index--
		i--
	}
	for j >= 0 {
		nums1[index] = nums2[j]
		index--
		j--
	}
}
func main() {
	nums1, m, nums2, n := []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3
	merge(nums1, m, nums2, n)
	fmt.Println("Sorted array: ", nums1)
}
