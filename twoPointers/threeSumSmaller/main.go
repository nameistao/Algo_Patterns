package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Expected: 2")
	fmt.Println("Actual:", searchTriplets([]int{-1, 0, 2, 3}, 3))
	fmt.Println("Expected: 4")
	fmt.Println("Actual:", searchTriplets([]int{-1, 4, 2, 1, 3}, 5))
}

//Question: squaring a sorted array
//Time: O(N^2)
//Space: O(1)
func searchTriplets(arr []int, target int) int {
	count := 0
	sort.Ints(arr)

	for i := 0; i < len(arr)-2; i++ {
		newTarget := target - arr[i]
		p1, p2 := i+1, len(arr)-1

		for p1 < p2 {
			if arr[p1]+arr[p2] < newTarget {
				count += p2 - p1
				p1++
			} else {
				p2--
			}
		}
	}

	return count
}
