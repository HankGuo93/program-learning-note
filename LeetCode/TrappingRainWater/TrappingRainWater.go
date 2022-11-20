package main

//description : https://leetcode.com/problems/trapping-rain-water/description/
//time complexity : O(3n)
//space complexity : O(2n)
func trap(height []int) int {
	heightLen := len(height)
	left := make([]int, heightLen)
	right := make([]int, heightLen)

	for i := 1; i < heightLen; i++ {
		left[i] = Max(height[i-1], left[i-1])
	}

	for i := heightLen - 2; i >= 0; i-- {
		right[i] = Max(height[i+1], right[i+1])
	}

	answer := 0
	for i := 0; i < heightLen; i++ {
		waterlevel := Min(left[i], right[i])
		if waterlevel > height[i] {
			answer += waterlevel - height[i]
		}
	}

	return answer
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
