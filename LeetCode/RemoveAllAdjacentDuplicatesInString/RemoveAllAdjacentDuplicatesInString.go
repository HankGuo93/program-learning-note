package main

import "strings"

func removeDuplicates(s string) string {
	sSplit := strings.Split(s, "")
	if len(sSplit) == 1 {
		return s
	}
	result := ""
	for true {
		if !innerRemoveDuplicates(s, sSplit, &result) {
			break
		}
	}
	return result
}

func innerRemoveDuplicates(s string, sSplit []string, result *string) bool {
	centralIndex := -1
	for i := 0; i < len(sSplit)-1; i++ {
		if sSplit[i] == sSplit[i+1] {
			centralIndex = i
			break
		}
	}
	if centralIndex == -1 {
		return true
	}
	duplicates := ""
	leftIndex := centralIndex
	rightIndex := centralIndex + 1
	for true {
		leftString := sSplit[leftIndex]
		rightString := sSplit[rightIndex]
		duplicates = leftString + duplicates + rightString
		leftIndex -= 1
		rightIndex += 1
		if leftIndex < 0 || rightIndex > len(sSplit)-1 {
			break
		}
		if sSplit[leftIndex] != sSplit[rightIndex] {
			break
		}
	}
	replaceString := strings.Replace(s, duplicates, "", -1)
	result = &replaceString
	return false
}
