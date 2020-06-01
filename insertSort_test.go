package goAlgorithm

import (
	"fmt"
	"testing"
)

type NumsTable struct {
	nums []int
}

var nums0 = []int{1, 2, 3, 3, 2, 1}
var nums1 = []int{1, 3, 5, 4, 6, 5}
var nums2 = []int{1, 1, 2, 98, 6, 5}

var table = []NumsTable{
	{nums0},
	{nums1},
	{nums2},
}

func TestInsertSort2(t *testing.T) {
	res := InsertSort([]int{1, 3, 4, 2, 1})
	fmt.Println(res)
}

//func TestInsertSort(t *testing.T) {
//	resNums0 := InsertSort(table[0].nums)
//	resNums1 := InsertSort(table[1].nums)
//	resNums2 := InsertSort(table[2].nums)
//	corNums0 := []int{1, 1, 2, 2, 3, 3}
//	corNums1 := []int{1, 3, 4, 5, 5, 6}
//	corNums2 := []int{1, 1, 2, 5, 6, 98}
//	if reflect.DeepEqual(resNums0, corNums0) || reflect.DeepEqual(resNums1, corNums1) || reflect.DeepEqual(resNums2, corNums2) {
//		t.Errorf("计算出来的结果为： #{resNums0}, 正确的结果应该为： #{corNums0")
//		t.Errorf("计算出来的结果为： #{resNums1}, 正确的结果应该为： #{corNums1")
//		t.Errorf("计算出来的结果为： #{resNums2}, 正确的结果应该为： #{corNums2")
//	}
//}