package main

//description : https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
//time complexity : O(n)
//space complexity : O(n)
func removeDuplicates(nums []int) int {
	result := []int{}
	exist := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		value := (nums)[i]
		_, ok := exist[value]
		if ok {
			continue
		}
		exist[value] = value
		result = append(result, value)
	}
	for i := 0; i < len(result); i++ {
		nums[i] = result[i]
	}
	return len(result)
}
