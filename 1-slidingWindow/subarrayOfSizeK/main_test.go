package main

import (
	"testing"
)

type FindMaxSumSubarrayTest struct {
	arg1 int
	arg2 []int
	expected int
}

var FindMaxSumSubarrayTests = []FindMaxSumSubarrayTest{
	{3,[]int{2, 1, 5, 1, 3, 2}, 9},
	{2,[]int{2, 3, 4, 1, 5}, 7},
}

func TestFindMaxSumSubarray(t *testing.T) {
	
}