package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	start := &ListNode{}
	current := start

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	return start.Next
}

func main() {
	// 1 2 4
	list1node4 := ListNode{Val: 4}
	list1node2 := ListNode{Val: 2, Next: &list1node4}
	list1node1 := ListNode{Val: 1, Next: &list1node2}
	// 1 3 4
	list2node4 := ListNode{Val: 4}
	list2node3 := ListNode{Val: 3, Next: &list2node4}
	list2node1 := ListNode{Val: 1, Next: &list2node3}

	node := mergeTwoLists(&list1node1, &list2node1)
	for node.Next != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
	fmt.Println(node.Val)
}
