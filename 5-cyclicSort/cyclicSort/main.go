package main

import "fmt"

func main() {
	fmt.Println("Expected: [1 2 3 4 5]")
	fmt.Println("Actual:", sort([]int{1, 2, 3, 4, 5}))
	fmt.Println("Expected: [1 2 3 4 5 6]")
	fmt.Println("Actual:", sort([]int{2, 6, 4, 3, 1, 5}))
	fmt.Println("Expected: [1 2 3 4 5 6]")
	fmt.Println("Actual:", sort([]int{1, 5, 6, 4, 3, 2}))
}

//Question: cyclic sort
//Time: O(N)
//Space: O(1)
func sort(nums []int) []int {
	for i := range nums {
		for nums[i] != i+1 {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	return nums
}
