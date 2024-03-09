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

func isSymmetric(root *TreeNode) bool {
	vlrQueue := []*TreeNode{root}
	vlrVals := []int{}
	vrlQueue := []*TreeNode{root}
	vrlVals := []int{}
	i := 0
	treeLen := 1
	for ; i < treeLen; i, treeLen = i+1, len(vlrQueue) {
		if vlrQueue[i] != nil {
			vlrVals = append(vlrVals, vlrQueue[i].Val)
			vlrQueue = append(vlrQueue, vlrQueue[i].Left)
			vlrQueue = append(vlrQueue, vlrQueue[i].Right)
		} else {
			vlrVals = append(vlrVals, math.MinInt32)
		}
		if vrlQueue[i] != nil {
			vrlVals = append(vrlVals, vrlQueue[i].Val)
			vrlQueue = append(vrlQueue, vrlQueue[i].Right)
			vrlQueue = append(vrlQueue, vrlQueue[i].Left)
		} else {
			vrlVals = append(vrlVals, math.MinInt32)
		}
	}
	// checking symetric?
	for i := 0; i < len(vlrVals); i++ {
		if vlrVals[i] != vrlVals[i] {
			return false
		}
	}
	return true
}

func main() {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println("Is tree symmetric?", isSymmetric(&root))
}
