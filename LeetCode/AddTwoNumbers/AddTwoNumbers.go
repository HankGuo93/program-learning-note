package main

//description : https://leetcode.com/problems/add-two-numbers/
//time complexity : O(n)
//space complexity : O(3)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ten := 0
	currentNode := new(ListNode)
	RootNode := currentNode
	for l1 != nil || l2 != nil {
		currentNode.Next = new(ListNode)
		currentNode = currentNode.Next
		l1Value := 0
		l2Value := 0
		if l1 != nil {
			l1Value = l1.Val
		}
		if l2 != nil {
			l2Value = l2.Val
		}
		tmpSum := l1Value + l2Value + ten
		if tmpSum >= 10 {
			ten = 1
			tmpSum -= 10
		} else {
			ten = 0
		}

		currentNode.Val += tmpSum

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if ten == 1 {
		currentNode.Next = new(ListNode)
		currentNode.Next.Val = ten
	}
	return RootNode.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
