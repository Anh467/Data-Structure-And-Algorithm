package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	temp := head
	for {
		temp = temp.Next
		pointer := head
		for pointer.Next != temp.Next {
			if pointer.Val > temp.Val {
				pointer.Val = pointer.Val ^ temp.Val
				temp.Val = temp.Val ^ pointer.Val
				pointer.Val = pointer.Val ^ temp.Val
			}
			pointer = pointer.Next
		}
		if temp.Next == nil {
			break
		}
	}
	return head
}
func printPointer(pointer *ListNode) {
	for {
		fmt.Print(pointer.Val)
		if pointer.Next == nil {
			break
		}
		pointer = pointer.Next
	}
	fmt.Println()
}
func main() {
	/*
		//case 2
		node4 := ListNode{Val: 0, Next: nil}
		node3 := ListNode{Val: 4, Next: &node4}
		node2 := ListNode{Val: 3, Next: &node3}
		node1 := ListNode{Val: 5, Next: &node2}
		node0 := ListNode{Val: -1, Next: &node1}
	*/

	// case 1
	node3 := ListNode{Val: 3, Next: nil}
	node2 := ListNode{Val: 1, Next: &node3}
	node1 := ListNode{Val: 2, Next: &node2}
	node0 := ListNode{Val: 4, Next: &node1}

	pointer := node0
	printPointer(&pointer)
	node := insertionSortList(&node0)
	pointers := node
	printPointer(pointers)

}
