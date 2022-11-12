package main

import (
	"strings"
)

//description : https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
//time complexity : O(n^2)
//space complexity : O(1)
func removeDuplicates(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			s = strings.Replace(s, string(s[i])+string(s[i]), "", -1)
			i = -1
		}
	}
	return s
}
