package main

var results [][]int

//description : https://leetcode.com/problems/combination-sum/
//time complexity : O(nlogn)
//space complexity : O(1)
func combinationSum(candidates []int, target int) [][]int {
	results = [][]int{}
	fundAllSetThatFitTarget(target, candidates, []int{}, 0)
	return results
}

func fundAllSetThatFitTarget(target int, candidates []int, set []int, sumValue int) {
	if sumValue > target {
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
		var tmpSetLen = len(tmpSet)
		if tmpSetLen >= 2 && tmpSet[tmpSetLen-2] > tmpSet[tmpSetLen-1] {
			continue
		}
		fundAllSetThatFitTarget(target, candidates, tmpSet, tmp)
	}
}
