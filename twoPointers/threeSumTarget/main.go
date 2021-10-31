package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Expected: 1")
	fmt.Println("Actual:", searchTriplet([]int{-2, 0, 1, 2}, 2))
	fmt.Println("Expected: 0")
	fmt.Println("Actual:", searchTriplet([]int{-3, -1, 1, 2}, 1))
	fmt.Println("Expected: 3")
	fmt.Println("Actual:", searchTriplet([]int{1, 0, 1, 1}, 100))
}

//Question: triplet sum close to target
//Time: O(N^2)
//Space: O(1)
func searchTriplet(arr []int, targetSum int) int {
	diff, output := int(^uint(0)>>1), 0
	sort.Ints(arr)

	for i := 0; i < len(arr)-2; i++ {
		p1, p2 := i+1, len(arr)-1

		for p1 < p2 {
			curDiff := abs(targetSum - (arr[i] + arr[p1] + arr[p2]))
			if curDiff < diff {
				diff = curDiff
				output = arr[i] + arr[p1] + arr[p2]
			}

			if arr[i]+arr[p1]+arr[p2] < targetSum {
				p1++
			} else {
				p2--
			}
		}
	}
	return output
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
