package main

import "fmt"

func main() {
	fmt.Println("Expected: [2], [5], [2, 5], [3], [5, 3], [10]")
	fmt.Println("Actual:", findSubarrays([]int{2, 5, 3, 10}, 30))
	fmt.Println("Expected: [8], [2], [8, 2], [6], [2, 6], [5], [6, 5] ")
	fmt.Println("Actual:", findSubarrays([]int{8, 2, 6, 5}, 50))
}

//Question: subarrays with product less than a target
//Time: O(N^2)
//Space: O(1)
func findSubarrays(arr []int, target int) [][]int {
	mult, start := 1, 0
	answer := make([][]int,0)

	for i,v := range arr{
		mult *= v

		for start < i && mult >= target {
			mult /= arr[start]
			start++
		}

		temp := make([]int, 0)
		for j:=i;j>=start;j--{
			temp = append([]int{arr[j]}, temp...)
			answer = append(answer, temp)
		}
	}

	return answer
}