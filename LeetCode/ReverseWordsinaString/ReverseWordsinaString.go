package main

import "strings"

func reverseWords(s string) string {
	splitStrings := strings.Split(s, " ")
	result := ""
	for i := len(splitStrings) - 1; i >= 0; i-- {
		if splitStrings[i] == "" {
			continue
		}
		if i == 0 {
			result += splitStrings[i]
			break
		}
		result += splitStrings[i] + " "
	}
	return strings.Trim(result, " ")
}
