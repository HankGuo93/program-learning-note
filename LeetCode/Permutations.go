package main

import "fmt"

var Visited []bool
var StepsRecord = New()
var Nums []int
var Results [][]int

//description : https://leetcode.com/problems/permutations/submissions/
//time complexity : O(nlogn)
//space complexity : O(n)
func permute(nums []int) [][]int {
	Results = [][]int{}
	Visited = make([]bool, len(nums))
	Nums = nums
	FindCombinations()
	return Results
}

func FindCombinations() {
	if CheckAllValueIsTure() {
		sliceRecordSliced := StepsRecord.ToSlice()
		sliceRecordSlicedToInt := ConvertInterfacesToInts(sliceRecordSliced)
		Results = append(Results, sliceRecordSlicedToInt)
	}
	for i, value := range Visited {
		if value == false {
			Record(i)
			FindCombinations()
			UnRecord(i)
		}
	}
}

func Record(i int) {
	Visited[i] = true
	StepsRecord.Push(Nums[i])
}

func UnRecord(i int) {
	StepsRecord.Pop()
	Visited[i] = false
}

func CheckAllValueIsTure() bool {
	for i := 0; i < len(Visited); i++ {
		if Visited[i] == false {
			return false
		}
	}
	return true
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
