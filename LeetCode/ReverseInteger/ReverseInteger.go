package main

//description : https://leetcode.com/problems/reverse-integer/description/
//time complexity : O(n)
//space complexity : O(1)
func reverse(x int) int {
	times := 10
	result := 0
	for x != 0 {
		result += x % times * times
		result *= 10
		times *= 10
		x /= 10
	}
	return result
}
