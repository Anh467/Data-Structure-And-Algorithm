package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	lTemp := l
	var prev *ListNode
	var numTemp int = 0
	var totalTemp int = 0
	var n1, n2 int
	for {
		if l1 == nil && l2 == nil {
			if numTemp != 0 {
				lTemp.Val = numTemp
				lTemp.Next = nil
			} else {
				prev.Next = nil
			}
			break
		}
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		totalTemp = n1 + n2 + numTemp
		if totalTemp > 9 {
			numTemp = 1
			totalTemp = totalTemp % 10
		} else {
			numTemp = 0
		}
		lTemp.Val = totalTemp
		lTemp.Next = &ListNode{
			Next: nil,
		}
		prev = lTemp
		lTemp = lTemp.Next
	}
	lTemp = nil
	return l
}

func initList(nums []int) *ListNode {
	l := &ListNode{}
	lTemp := l
	for index, n := range nums {
		lTemp.Val = n
		if index == len(nums)-1 {
			lTemp.Next = nil
			break
		}
		lTemp.Next = &ListNode{}
		lTemp = lTemp.Next
	}
	lTemp = nil
	return l
}
func printList(l *ListNode) {
	println()
	print("List: ")
	var lTemp *ListNode = l
	for lTemp != nil {
		print(lTemp.Val)
		lTemp = lTemp.Next
	}
	println()
}
func main() {
	/*
		Example 1:
		Input: l1 = [2,4,3], l2 = [5,6,4]
		Output: [7,0,8]
		Explanation: 342 + 465 = 807.
	*/
	/*
		Example 2:
		Input: l1 = [0], l2 = [0]
		Output: [0]
	*/
	/*
		Example 3:
		Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
		Output: [8,9,9,9,0,0,0,1]
	*/
	var num1 = []int{9, 9, 9, 9, 9, 9, 9}
	var num2 = []int{9, 9, 9, 9}
	l1 := initList(num1)
	l2 := initList(num2)
	l := addTwoNumbers(l1, l2)
	printList(l)
}
