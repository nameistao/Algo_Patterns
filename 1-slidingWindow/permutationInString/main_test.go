package main

import (
	"testing"
)

type FindPermutationTest struct {
	arg1, arg2 string
	expected   bool
}

var FindPermutationTests = []FindPermutationTest{
	{"oidbcaf", "abc", true},
	{"odicf", "dc", false},
	{"bcdxabcdy", "bcdyabcdx", true},
	{"aaacb", "abc", true},
}

func TestFindPermutation(t *testing.T) {
	for _, v := range FindPermutationTests {
		actual := findPermutation(v.arg1, v.arg2)
		if actual != v.expected {
			t.Errorf("expected %v, actual %v", v.expected, actual)
		}
	}
}
