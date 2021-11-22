package main

import (
	"testing"
)

type FindStringAnagramsTest struct {
	arg1, arg2 string
	expected []int
}

var FindStringAnagramsTests = []FindStringAnagramsTest{
	{"ppqp", "pq", []int{1,2}},
	{"abbcabc", "abc", []int{2,3,4}},
}

func TestFindStringAnagrams(t *testing.T) {
	for _,v := range FindStringAnagramsTests{
		actual := findStringAnagrams(v.arg1, v.arg2)
		if len(actual) != len(v.expected) {
			t.Errorf("expected %v, actual %v", v.expected, actual)
		}
		for i := range actual {
			if actual[i] != v.expected[i] {
				t.Errorf("expected %v, actual %v", v.expected, actual)
			}
		}
	}
}