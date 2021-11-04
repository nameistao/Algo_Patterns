package main

import "fmt"

func main() {
	fmt.Println("Expected: [0, 1, 4, 4, 9]")
	fmt.Println("Actual:", makeSquare([]int{-2, -1, 0, 2, 3}))
	fmt.Println("Expected: [0, 1, 1, 4, 9]")
	fmt.Println("Actual:", makeSquare([]int{-3, -1, 0, 1, 2}))
}

//Question: squaring a sorted array
//Time: O(N)
//Space: O(N)
func makeSquare(arr []int) []int {
	p1, p2, cur, ans := 0, len(arr)-1, len(arr)-1, make([]int, len(arr))

	for p1 < p2 {
		if arr[p1]*arr[p1] > arr[p2]*arr[p2] {
			ans[cur] = arr[p1] * arr[p1]
			p1++
			cur--
		} else {
			ans[cur] = arr[p2] * arr[p2]
			p2--
			cur--
		}
	}
	return ans
}
