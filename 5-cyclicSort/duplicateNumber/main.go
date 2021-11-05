package main

import "fmt"

func main() {
	fmt.Println("Expected: 4")
	fmt.Println("Actual:", findNumber([]int{1, 4, 4, 3, 2}))
	fmt.Println("Expected: 3")
	fmt.Println("Actual:", findNumber([]int{2, 1, 3, 3, 5, 4}))
	fmt.Println("Expected: 4")
	fmt.Println("Actual:", findNumber([]int{2, 4, 1, 4, 4}))
}

//Question: find the duplicate number
//Time: O(N)
//Space: O(1)
func findNumber(nums []int) int {
	for i := range nums {
		for nums[i] != i+1 {
			if nums[i] == nums[nums[i]-1] {
				return nums[i]
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	return nums[len(nums)-1]
}
