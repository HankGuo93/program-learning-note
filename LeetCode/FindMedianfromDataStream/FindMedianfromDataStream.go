package main

import "container/heap"

//description : https://leetcode.com/problems/find-median-from-data-stream/description/
//time complexity : O(log n)
//space complexity : O(n)
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	min := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return min
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	min := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return min
}

type MedianFinder struct {
	minHeap MinHeap
	maxHeap MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: MinHeap{},
		maxHeap: MaxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 || num <= this.maxHeap[0] {
		heap.Push(&this.maxHeap, num)
	} else {
		heap.Push(&this.minHeap, num)
	}
	if this.maxHeap.Len() > this.minHeap.Len()+1 {
		maxLeft := heap.Pop(&this.maxHeap)
		heap.Push(&this.minHeap, maxLeft)
	}
	if this.minHeap.Len() > this.maxHeap.Len() {
		minRight := heap.Pop(&this.minHeap)
		heap.Push(&this.maxHeap, minRight)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() == this.maxHeap.Len() {
		return float64(this.minHeap[0]+this.maxHeap[0]) / 2
	}
	return float64(this.maxHeap[0])
}
