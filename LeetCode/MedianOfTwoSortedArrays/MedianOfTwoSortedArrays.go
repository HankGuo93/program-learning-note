package main

import "sort"

//description : https://leetcode.com/problems/median-of-two-sorted-arrays/
//time complexity : O(nlogn)
//space complexity : O(n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var mergeSet = append(nums1, nums2...)
	sort.Ints(mergeSet)
	var medianIndex = len(mergeSet) / 2
	var remainder = len(mergeSet) % 2
	if remainder == 0 {
		return (float64(mergeSet[medianIndex]) + float64(mergeSet[medianIndex+1])) / 2
	}
	return float64(mergeSet[medianIndex+1])
}
