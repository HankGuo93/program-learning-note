package main

import "math/rand"

type RandomizedSet struct {
	mySet  map[int]int
	myInts []int
}

//description : https://leetcode.com/problems/insert-delete-getrandom-o1/description/
//time complexity : O(1)
//space complexity : O(n)
func Constructor() RandomizedSet {
	return RandomizedSet{
		mySet:  make(map[int]int),
		myInts: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exist := this.mySet[val]
	if exist {
		return false
	}
	this.myInts = append(this.myInts, val)
	this.mySet[val] = len(this.myInts) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	_, exist := this.mySet[val]
	if exist {
		this.myInts[this.mySet[val]] = this.myInts[len(this.myInts)-1]
		this.mySet[this.myInts[len(this.myInts)-1]] = this.mySet[val]
		this.myInts = this.myInts[:len(this.myInts)-1]
		delete(this.mySet, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.myInts[rand.Intn(len(this.myInts))]
}
