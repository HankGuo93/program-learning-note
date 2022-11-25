package main

import (
	"math"
)

//description : https://leetcode.com/problems/validate-binary-search-tree/description/
//time complexity : O(n)
//space complexity : O(1)
var validTag bool

func isValidBST(root *TreeNode) bool {
	validTag = true
	findNotValidNodeExist(root, int64(math.MinInt32)-1, int64(math.MaxInt32)+1)
	return validTag
}

func findNotValidNodeExist(node *TreeNode, minLimit int64, maxLimit int64) {
	if node == nil {
		return
	}
	if int64(node.Val) <= minLimit || int64(node.Val) >= maxLimit {
		validTag = false
		return
	}
	findNotValidNodeExist(node.Left, minLimit, int64(node.Val))
	findNotValidNodeExist(node.Right, int64(node.Val), maxLimit)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
