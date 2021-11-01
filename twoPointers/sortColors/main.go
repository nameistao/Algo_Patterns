package main

import "fmt"

func main() {
	fmt.Println("Expected: [0 0 1 1 2]")
	fmt.Println("Actual:", sort([]int{1, 0, 2, 1, 0}))
	fmt.Println("Expected: [0 0 1 2 2 2]")
	fmt.Println("Actual:", sort([]int{2, 2, 0, 1, 2, 0}))
}

//Question: sort colors
//Time: O(N)
//Space: O(1)
func sort(arr []int) []int {
	p1, p2 := 0, len(arr)-1

	for i := 0; i <= p2; {
		if arr[i] == 0 {
			arr[i], arr[p1] = arr[p1], arr[i]
			p1++
			i++
		} else if arr[i] == 2 {
			arr[i], arr[p2] = arr[p2], arr[i]
			p2--
		} else {
			i++
		}
	}

	return arr
}
