package main

//description : https://leetcode.com/problems/jump-game/
//time complexity : O(n)
//space complexity : O(1)
func canJump(nums []int) bool {
	max := 0
	target := len(nums) - 1
	for i := 0; i < len(nums) && i <= max; i++ {
		if max < i+nums[i] {
			max = i + nums[i]
		}
		if max >= target {
			return true
		}
	}
	return false
}
