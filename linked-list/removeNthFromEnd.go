package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	start := &ListNode{Next: head}
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	length -= n
	current := start
	for length > 0 {
		length--
		current = current.Next
	}
	current.Next = current.Next.Next
	return start.Next
}

func main() {
	// 1 2 3 4
	node4 := ListNode{Val: 4}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}

	node := removeNthFromEnd(&node1, 1)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
