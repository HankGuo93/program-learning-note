package main

import (
	"sort"
)

var results [][]int

//description : https://leetcode.com/problems/combination-sum/
//time complexity : O(nlogn)
//space complexity : O(1)
func combinationSum(candidates []int, target int) [][]int {
	results = [][]int{}
	sort.Ints(candidates)
	fundAllSetThatFitTarget(target, candidates, []int{}, 0)
	return results
}

func fundAllSetThatFitTarget(target int, candidates []int, set []int, sumValue int) {
	if sumValue > target {
		return
	}
	if len(set) >= 2 && set[len(set)-2] > set[len(set)-1] {
		return
	}
	if sumValue == target {
		results = append(results, set)
		return
	}
	for i := 0; i < len(candidates); i++ {
		var tmpSet = make([]int, len(set))
		var tmp = sumValue
		copy(tmpSet[:], set)
		tmp += candidates[i]
		tmpSet = append(tmpSet, candidates[i])
		fundAllSetThatFitTarget(target, candidates, tmpSet, tmp)
	}
}
