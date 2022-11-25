package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//description : https://leetcode.com/problems/palindrome-linked-list/description/
//time complexity : O(n^2)
//space complexity : O(2n)
func isPalindrome(head *ListNode) bool {
	slc := []int{}
	rev_slc := []int{}
	for head != nil {
		slc = append(slc, head.Val)
		head = head.Next
	}

	for i := len(slc) - 1; i >= 0; i-- {
		rev_slc = append(rev_slc, slc[i])
	}
	return Equal(slc, rev_slc)
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
