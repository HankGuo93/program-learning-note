package main

import (
	"strconv"
	"strings"
)

//description : https://leetcode.com/problems/maximum-69-number/description/
//time complexity : O(n)
//space complexity : O(1)
func maximum69Number(num int) int {
	numToString := strconv.Itoa(num)
	numToStrings := strings.Split(numToString, "")
	sixIndex := -1
	for i := 0; i < len(numToStrings); i++ {
		if numToStrings[i] == "6" {
			sixIndex = i
			break
		}
	}
	if sixIndex == -1 {
		return num
	}
	times := 1
	for i := 0; i < len(numToString)-sixIndex-1; i++ {
		times *= 10
	}
	num -= 6 * times
	num += 9 * times
	return num
}
