package main

import (
	"strings"
)

//description : https://leetcode.com/problems/make-the-string-great/description/
//time complexity : O(n)
//space complexity : O(1)
func makeGood(s string) string {
	tmpString := s
	for i := 0; i < len(s); i++ {
		char := s[i]
		comparingBadString1 := getBadStringSet1(string(char))
		tmpString = strings.ReplaceAll(tmpString, comparingBadString1, "")
		comparingBadString2 := getBadStringSet2(string(char))
		tmpString = strings.ReplaceAll(tmpString, comparingBadString2, "")
	}
	return tmpString
}

func getBadStringSet1(s string) string {
	lowerString := strings.ToLower(s)
	upperString := strings.ToUpper(s)
	return lowerString + upperString
}

func getBadStringSet2(s string) string {
	lowerString := strings.ToLower(s)
	upperString := strings.ToUpper(s)
	return upperString + lowerString
}
