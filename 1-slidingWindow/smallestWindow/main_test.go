package main

import (
	"testing"
)

type FindSubstringTest struct {
	arg1, arg2, expected string
}

var FindSubstringTests = []FindSubstringTest{
	{"aabdec", "abc", "abdec"},
	{"abdbca", "abc", "bca"},
	{"adcad", "abc", ""},
}

func TestFindSubstring(t *testing.T) {
	for _, v := range FindSubstringTests {
		actual := findSubstring(v.arg1, v.arg2)
		if actual != v.expected {
			t.Errorf("expected %v, actual %v", v.expected, actual)
		}
	}
}
