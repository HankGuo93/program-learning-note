package main

var maxValue *int

//description : https://leetcode.com/problems/binary-tree-maximum-path-sum/description/
//time complexity : O(n)
//space complexity : O(1)
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxValue := root.Val

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		sum := root.Val + left + right
		if sum > maxValue {
			maxValue = sum
		}

		return max(left, right) + root.Val
	}

	dfs(root)
	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
