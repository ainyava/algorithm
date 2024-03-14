package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	low := 0
	high := len(nums)

	if low >= high {
		return nil
	}

	mid := (low + high) / 2

	root := TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return &root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Println("Tree of sorted array:", sortedArrayToBST(nums))
}
