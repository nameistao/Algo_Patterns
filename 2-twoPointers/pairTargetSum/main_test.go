package main

import (
	"testing"
)

type SearchTest struct {
	arg1     []int
	arg2     int
	expected []int
}

var SearchTests = []SearchTest{
	{[]int{1, 2, 3, 4, 6}, 6, []int{1, 3}},
	{[]int{2, 5, 9, 11}, 11, []int{0, 2}},
}

func TestSearch(t *testing.T) {
	for _, v := range SearchTests {
		actual := search(v.arg1, v.arg2)
		if len(actual) != len(v.expected) {
			t.Errorf("expected %v, actual %v", v.expected, actual)
		}

		for i := range v.expected {
			if v.expected[i] != actual[i] {
				t.Errorf("expected %v, actual %v", v.expected, actual)
			}
		}
	}
}
