package main

func TwoSum(nums []int, target int) []int {
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
