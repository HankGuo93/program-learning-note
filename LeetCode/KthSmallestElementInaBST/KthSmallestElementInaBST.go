package main

import "sort"

var nums []int

//description : https://leetcode.com/problems/kth-smallest-element-in-a-bst/
//time complexity : O(nlogn)
//space complexity : O(n)
func kthSmallest(root *TreeNode, k int) int {
	nums = []int{}
	findAllElement(root)
	sort.Ints(nums)
	return nums[k-1]
}

func findAllElement(node *TreeNode) {
	if node == nil {
		return
	}
	nums = append(nums, node.Val)
	findAllElement(node.Left)
	findAllElement(node.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
