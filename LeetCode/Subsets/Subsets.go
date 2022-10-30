package main

import "fmt"

var NumsLen int
var StepsRecord = New()

//description : https://leetcode.com/problems/subsets/submissions/
//time complexity : O(2n)
//space complexity : O(n)
func subsets(nums []int) [][]int {
	NumsLen = len(nums)
	results := [][]int{}
	FindSubSet(0, &results, nums)
	return results
}

func FindSubSet(layer int, results *[][]int, nums []int) {
	if NumsLen == layer {
		sliceRecordSliced := StepsRecord.ToSlice()
		sliceRecordSlicedToInt := ConvertInterfacesToInts(sliceRecordSliced)
		*results = append(*results, sliceRecordSlicedToInt)
		return
	}
	FindSubSet(layer+1, results, nums)
	StepsRecord.Push(nums[layer])
	FindSubSet(layer+1, results, nums)
	StepsRecord.Pop()
}

func ConvertInterfacesToInts(interfaces []interface{}) []int {
	ints := make([]int, len(interfaces))
	for i, value := range interfaces {
		switch typeValue := value.(type) {
		case int:
			ints[i] = typeValue
			break
		default:
			fmt.Println("Not an int: ", value)
		}
	}
	return ints
}

type Stack struct {
	list []interface{}
}

func New() *Stack {
	s := new(Stack)
	s.list = make([]interface{}, 0, 8)
	return s
}

func (s *Stack) IsEmpty() bool {
	return len(s.list) == 0
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	tmp := s.list[len(s.list)-1]

	s.list = s.list[:len(s.list)-1]
	return tmp
}

func (s *Stack) Push(element interface{}) {
	s.list = append(s.list, element)
}

func (s *Stack) ToSlice() []interface{} {
	return s.list
}
