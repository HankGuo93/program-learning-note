package main

//description : https://leetcode.com/problems/kth-smallest-element-in-a-bst/
//time complexity : O(n)
//space complexity : O(1)
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i+1 {
            continue
		}
        cur := nums[i]
        for cur <= len(nums) && cur > 0 && nums[cur-1] != cur {
            tmp := nums[cur-1]
            nums[cur-1] = cur
            cur = tmp
        }
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}