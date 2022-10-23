package main

import "fmt"

//description : https://leetcode.com/problems/two-sum/
// Time Complexity : O(n)
// Space Complexity : O(n)
func main() {
	var nums = []int{2, 7, 11, 15}
	var tartget int = 9
	var twoSumResult = twoSum(nums, tartget)
	fmt.Printf("nums = [2,7,11,15]\n")
	fmt.Printf("expected = [0,1]\n")
	fmt.Printf("twoSumResult = [%d,%d]", twoSumResult[0], twoSumResult[1])
}

func twoSum(nums []int, target int) []int {
	var arr []int
	resultMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		value, exists := resultMap[nums[i]]
		if exists {
			arr = append(arr, value)
			arr = append(arr, i)
			return arr
		} else {
			var key int = target - nums[i]
			resultMap[key] = i
		}
	}
	return arr
}
