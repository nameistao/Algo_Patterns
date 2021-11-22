package main

import (
	"testing"
)

type FindLengthTest struct {
	arg1 int
	arg2 []int
	expected int
}

var FindLengthTests = []FindLengthTest{
	{2, []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, 6},
	{3, []int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}, 9},
}

func TestFindLength(t *testing.T) {
	for _,v := range FindLengthTests {
		actual := findLength(v.arg1, v.arg2)
		if actual != v.expected {
			t.Errorf("expected %v, actual %v", v.expected, actual)
		}
	}
}