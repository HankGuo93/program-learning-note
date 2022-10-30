package main

//description : https://leetcode.com/problems/merge-two-sorted-lists/
//time complexity : O(n)
//space complexity : O(n)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var RootNode *ListNode
	var CurrentNode *ListNode
	for list1.Next != nil || list2.Next != nil {
		if list1.Val < list2.Val {
			if RootNode == nil {
				RootNode = list1
			} else {
				CurrentNode.Next = list1
			}
			CurrentNode = list1
			list1 = list1.Next
		} else {
			if RootNode == nil {
				RootNode = list2
			} else {
				CurrentNode.Next = list2
			}
			CurrentNode = list2
			list2 = list2.Next
		}
		if list1 == nil {
			if RootNode == nil {
				RootNode = list2
			} else {
				CurrentNode.Next = list2
			}
			CurrentNode = list2
			list2 = list2.Next
		}
		if list2 == nil {
			if RootNode == nil {
				RootNode = list2
			} else {
				CurrentNode.Next = list2
			}
			CurrentNode = list2
			list2 = list2.Next
		}
	}
	return RootNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
