package main

import "fmt"

func main(){
	fmt.Println("Expected: 3")
	fmt.Println("Actual:", findLength("aabccbb"))
	fmt.Println("Expected: 2")
	fmt.Println("Actual:", findLength("abbbb"))
	fmt.Println("Expected: 3")
	fmt.Println("Actual:", findLength("abccde"))
}

//Question: longest substring with distinct characters
//Time: O(N)
//Space: O(N)
func findLength(str string) int {
	m := make(map[int]int)
	var(
		start int
		max int
	)

	for i,v := range str{
		_, found := m[int(v)]
		if found {
			m[int(v)]++
		} else{
			m[int(v)] = 1
		}

		for m[int(v)] > 1{
			m[int(str[start])]--
			start++
		}

		max = maxInt(max, i-start+1)
	}
	return max
}

func maxInt(x int, y int) int {
	if x > y {
		return x
	}
	return y
}