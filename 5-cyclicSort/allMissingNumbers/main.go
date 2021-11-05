package main

import (
	"fmt"
)

func main() {
	fmt.Println("Expected: [4 6 7]")
	fmt.Println("Actual:", findNumbers([]int{2, 3, 1, 8, 2, 3, 5, 1}))
	fmt.Println("Expected: [3]")
	fmt.Println("Actual:", findNumbers([]int{2, 4, 1, 2}))
	fmt.Println("Expected: [4]")
	fmt.Println("Actual:", findNumbers([]int{2, 3, 2, 1}))
}

//Question: find all missing numbers
//Time: O(N)
//Space: O(N)
func findNumbers(nums []int) []int {
	for i := range nums {
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	answer := make([]int, 0)
	for i := range nums {
		if nums[i] != i+1 {
			answer = append(answer, i+1)
		}
	}
	return answer
}
