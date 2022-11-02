package main

import (
	"container/heap"
)

//description : https://leetcode.com/problems/merge-k-sorted-lists/
//time complexity : O(2n)
//space complexity : O(n)
func mergeKLists(lists []*ListNode) *ListNode {
	h := &IntHeap{}
	root := new(ListNode)
	heap.Init(h)
	for i := 0; i < len(lists); i++ {
		currentNode := lists[i]
		for currentNode != nil {
			heap.Push(h, currentNode.Val)
			currentNode = currentNode.Next
		}
	}
	tmpNode := new(ListNode)
	root = tmpNode
	for h.Len() > 0 {
		newNode := new(ListNode)
		newNode.Val = heap.Pop(h).(int)
		tmpNode.Next = newNode
		tmpNode = tmpNode.Next
	}
	return root.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
