package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Expected: 2")
	fmt.Println("Actual:", findMinSubArray(7, []int{2, 1, 5, 2, 3, 2}))
	fmt.Println("Expected: 1")
	fmt.Println("Actual:", findMinSubArray(7, []int{2, 1, 5, 2, 8}))
	fmt.Println("Expected: 3")
	fmt.Println("Actual:", findMinSubArray(8, []int{3, 4, 1, 1, 6}))
}

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
