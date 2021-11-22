package main

import (
	"testing"
)

type FindLengthTest struct {
	arg      []int
	expected int
}

var FindLengthTests = []FindLengthTest{
	{[]int{1, 2, 3, 1, 3}, 3},
	{[]int{1, 2, 3, 2, 2, 3}, 5},
	{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
}

func TestFindLength(t *testing.T) {
	for _, v := range FindLengthTests {
		actual := findLength(v.arg)
		if actual != v.expected {
			t.Errorf("expected %v, actual %v", v.expected, actual)
		}
	}
}
