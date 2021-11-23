package main

import (
	"testing"
)

type FindLengthTest struct {
	arg1     int
	arg2     string
	expected int
}

var FindLengthTests = []FindLengthTest{
	{2, "aabccbb", 5},
	{1, "abbcb", 4},
	{1, "abccde", 3},
}

func TestFindLength(t *testing.T) {
	for _, v := range FindLengthTests {
		actual := findLength(v.arg1, v.arg2)
		if actual != v.expected {
			t.Errorf("expected %v, actual %v", v.expected, actual)
		}
	}
}
