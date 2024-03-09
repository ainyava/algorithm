package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidSubTree(node *TreeNode, lower, higher int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= higher {
		return false
	}
	return isValidSubTree(node.Left, lower, node.Val) && isValidSubTree(node.Right, node.Val, higher)
}

func isValidBST(root *TreeNode) bool {
	return isValidSubTree(root, math.MinInt64, math.MaxInt64)
}

func main() {
	root := TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println("Is valid BST?", isValidBST(&root))
}
