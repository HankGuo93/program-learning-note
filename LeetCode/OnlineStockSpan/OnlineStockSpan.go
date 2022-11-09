package main

//description : https://leetcode.com/problems/online-stock-span/description/
//time complexity : O(n^2)
//space complexity : O(n)
type StockSpanner struct {
	CurrentNode *Node
}

func Constructor() StockSpanner {
	stockSpanner := new(StockSpanner)
	stockSpanner.CurrentNode = &Node{Val: -1}
	return *stockSpanner
}

func (this *StockSpanner) Next(price int) int {
	tmpNodeA := this.CurrentNode
	this.CurrentNode.Next = &Node{Val: price}
	this.CurrentNode = this.CurrentNode.Next
	this.CurrentNode.Before = tmpNodeA
	counter := 1
	tmpNodeB := this.CurrentNode
	for tmpNodeB.Before.Val != -1 {
		tmpNodeB = tmpNodeB.Before
		if tmpNodeB.Val > price {
			break
		}
		counter++
	}
	return counter
}

type Node struct {
	Val    int
	Next   *Node
	Before *Node
}
