package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	for ; head != nil; head = head.Next {
		if head.Val == math.MaxInt32 {
			return true
		}
		head.Val = math.MaxInt32
	}
	return false
}
func main() {
	// 1 2 3 4
	node4 := ListNode{Val: 1}
	node3 := ListNode{Val: 2, Next: &node4}
	node2 := ListNode{Val: 3, Next: &node3}
	node1 := ListNode{Val: 4, Next: &node2}
	node4.Next = &node2
	fmt.Println("HasCycle?", hasCycle(&node1))
}
