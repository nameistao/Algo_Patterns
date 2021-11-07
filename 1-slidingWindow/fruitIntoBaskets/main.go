package main

import "fmt"

func main() {
	fmt.Println("Expected: 3")
	fmt.Println("Actual:", findLength([]int{1, 2, 3, 1, 3}))
	fmt.Println("Expected: 5")
	fmt.Println("Actual:", findLength([]int{1, 2, 3, 2, 2, 3}))
	fmt.Println("Expected: 5")
	fmt.Println("Actual:", findLength([]int{3,3,3,1,2,1,1,2,3,3,4}))
}

//Question: fruits into baskets
//Time: O(N)
//Space: O(1)
func findLength(fruits []int) int {
	start, cur, max := 0,0,0
	m := make(map[int]int)

	for _, v := range fruits {
		_, found := m[v]
		if found {
			m[v]++
		} else {
			m[v] = 1
		}
		cur++

		for len(m) > 2 {
			val := fruits[start]
			if m[val] > 1 {
				m[val]--
			} else {
				delete(m, val)
			}
			cur--
            start++
		}

		max = maxInt(max, cur)
	}

	return max
}

func maxInt(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
