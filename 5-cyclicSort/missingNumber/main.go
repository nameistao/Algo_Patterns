package main

import "fmt"

func main() {
	fmt.Println("Expected: 2")
	fmt.Println("Actual:", findMissingNumber([]int{4, 0, 3, 1}))
	fmt.Println("Expected: 7")
	fmt.Println("Actual:", findMissingNumber([]int{8, 3, 5, 2, 4, 6, 0, 1}))
}

//Question: find the missing number
//Time: O(N)
//Space: O(1)
func findMissingNumber(nums []int) int {
	for i := range nums {
		for nums[i] != i {
			if nums[i] != len(nums) {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			} else {
				break
			}
		}
	}

	for i := range nums {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
