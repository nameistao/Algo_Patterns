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
	p1, p2 := 0, 1
	for p2 < len(arr) {
		if arr[p2] == 0 && arr[p1] != 0 {
			arr[p2], arr[p1] = arr[p1], arr[p2]
		}
		p2++
	}
	p2 = p1 + 1
	for p2 < len(arr) {
		if arr[p2] == 1 && arr[p1] != 1 {
			temp := arr[p2]
			arr[p2] = arr[p1]
			arr[p1] = temp
			p1++
		}
		p2++
	}
	return arr
}
