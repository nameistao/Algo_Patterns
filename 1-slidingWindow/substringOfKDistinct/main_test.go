package main

import (
	"testing"
)

type LongestSubstringKDistinctTest struct {
	arg1 int
	arg2 string
	expected int
}

var LongestSubstringKDistinctTests = []LongestSubstringKDistinctTest{
	{2, "araaci", 4},
	{1, "araaci", 2},
	{3, "cbbebi",5},
	{10, "cbbebi",6},
}

func TestLongestSubstringKDistinct(t *testing.T) {
	for _,v := range LongestSubstringKDistinctTests {
		actual := longestSubstringKDistinct(v.arg1, v.arg2)
		if actual != v.expected {
			t.Errorf("expected %v, actual %v", v.expected, actual)
		}
	}
}