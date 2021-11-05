package main

import "fmt"

func main() {
	fmt.Println("Expected: [5 4]")
	fmt.Println("Actual:", findNumbers([]int{3, 4, 4, 5, 5}))
	fmt.Println("Expected: [3 5]")
	fmt.Println("Actual:", findNumbers([]int{5, 4, 7, 2, 3, 5, 3}))
}

//Question: find all duplicate numbers
//Time: O(N)
//Space: O(N)
func findNumbers(nums []int) []int {
	answer := make([]int, 0)
	for i := range nums {
		for nums[i] != i+1 {
			if nums[i] == nums[nums[i]-1] {
				answer = append(answer, nums[i])
				break
			}

			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	return answer
}
