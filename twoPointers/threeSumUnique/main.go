package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Expected: [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]")
	fmt.Println("Actual:", searchTriplets([]int{-3, 0, 1, 2, -1, 1, -2}))
	fmt.Println("Expected: [[-5, 2, 3], [-2, -1, 3]]")
	fmt.Println("Actual:", searchTriplets([]int{-5, 2, -1, -2, 3}))
}

//Question: remove deplicates
//Time: O(N^2)
//Space: O(N)
func searchTriplets(arr []int) [][]int {
	answer := make([][]int, 0)
	sort.Ints(arr)

	for i := 0; i < len(arr)-2; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		p1, p2, sum := i+1, len(arr)-1, -arr[i]
		for p1 < p2 {
			if arr[p1]+arr[p2] < sum {
				p1++
			} else if arr[p1]+arr[p2] > sum {
				p2--
			} else {
				if p1-1 == i || arr[p1-1] != arr[p1] {
					answer = append(answer, []int{arr[i], arr[p1], arr[p2]})
				}
				p1++
			}
		}
	}
	return answer
}
