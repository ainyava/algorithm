package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftLen := maxDepth(root.Left)
	rightLen := maxDepth(root.Right)
	if leftLen > rightLen {
		return 1 + leftLen
	} else {
		return 1 + rightLen
	}
}

func main() {
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println("Length of tree:", maxDepth(&root))
}
