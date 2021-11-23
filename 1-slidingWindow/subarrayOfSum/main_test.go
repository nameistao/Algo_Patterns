package main

import (
	"testing"
)

type FindMinSubArrayTest struct {
	arg1     int
	arg2     []int
	expected int
}

var FindMinSubArrayTests = []FindMinSubArrayTest{
	{7, []int{2, 1, 5, 2, 3, 2}, 2},
	{7, []int{2, 1, 5, 2, 8}, 1},
	{8, []int{3, 4, 1, 1, 6}, 3},
}

func TestFindSubArray(t *testing.T) {
	for _, v := range FindMinSubArrayTests {
		actual := findMinSubArray(v.arg1, v.arg2)
		if actual != v.expected {
			t.Errorf("expected %v, actual %v", v.expected, actual)
		}
	}
}
