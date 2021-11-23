package main

import (
	"testing"
)

type FindLengthTest struct {
	arg1 string
	expected int
}

var FindLengthTests = []FindLengthTest{
	{"aabccbb",3},
	{"abbbb", 2},
	{"abccde", 3},
}

func TestFindLength(t *testing.T) {
	for _,v := range FindLengthTests {
		actual := findLength(v.arg1)
		if actual != v.expected {
			t.Errorf("expected %v, actual %v", v.expected, actual)
		}
	}
}