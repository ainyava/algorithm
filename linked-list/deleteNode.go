package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	*node = *(node.Next)
}

func main() {
	// 4 5 1 9
	node9 := ListNode{Val: 9}
	node1 := ListNode{Val: 1, Next: &node9}
	node5 := ListNode{Val: 5, Next: &node1}
	node := ListNode{Val: 4, Next: &node5}

	deleteNode(&node1)
	for node.Next != nil {
		fmt.Println(node.Val)
		node = *node.Next
	}
	fmt.Println(node.Val)
}
