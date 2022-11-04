package main

type MinStack struct {
	nums []int
}

func Constructor() MinStack {
	return *new(MinStack)
}

func (this *MinStack) Push(val int) {
	this.nums = append(this.nums, val)
}

func (this *MinStack) Pop() {
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
	popInt := this.nums[len(this.nums)-1]
	return popInt
}

func (this *MinStack) GetMin() int {
	min := this.nums[0]
	for i := 1; i < len(this.nums); i++ {
		if min > this.nums[i] {
			min = this.nums[i]
		}
	}
	return min
}
