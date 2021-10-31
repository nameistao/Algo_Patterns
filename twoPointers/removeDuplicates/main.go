package main

import "fmt"

func main() {
	fmt.Println("Expected: 4")
	fmt.Println("Actual:", remove([]int{2, 3, 3, 3, 6, 9, 9}))
	fmt.Println("Expected: 2")
	fmt.Println("Actual:", remove([]int{2, 2, 2, 11}))
}

//Question: remove deplicates
//Time: O(N)
//Space: O(1)
func remove(arr []int) int {
	p1 := 0
	for _, v := range arr {
		if v != arr[p1] {
			p1++
			arr[p1] = v
		}
	}
	return p1 + 1
}
