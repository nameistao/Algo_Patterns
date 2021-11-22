package main

import "math"

//Question: smallest subarray with a given sum
//Time: O(N)
//Space: O(1)
func findMinSubArray(s int, arr []int) int {
	sum, start, min := 0, 0, math.MaxInt32

	for i, v := range arr {
		sum += v

		for sum >= s {
			min = minInt(min, i-start+1)
			sum -= arr[start]
			start++
		}
	}

	if min == math.MaxInt32 {
		return 0
	}

	return min
}

func minInt(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
