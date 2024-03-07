package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var (
		prev    *ListNode
		current = head
		next    *ListNode
	)
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	// 1 2 3 4
	node4 := ListNode{Val: 4}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}

	node := reverseList(&node1)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
