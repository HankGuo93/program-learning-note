package main

//description : https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
//time complexity : O(n^2)
//space complexity : O(n)
func lengthOfLongestSubstring(s string) int {
	maxNumber := 0
	mapping := make(map[byte]int)
	counter := 0
	for j := 0; j < len(s); j++ {
		value, exist := mapping[s[j]]
		if exist {
			counter = 0
			mapping = make(map[byte]int)
			j = value + 1
		}
		mapping[s[j]] = j
		counter++
		if maxNumber < counter {
			maxNumber = counter
		}
	}
	return maxNumber
}
