package main

import "fmt"

func main(){
	fmt.Println("Expected: 4")
	fmt.Println("Actual:", longestSubstringKDistinct(2, "araaci"))
	fmt.Println("Expected: 2")
	fmt.Println("Actual:", longestSubstringKDistinct(1, "araaci"))
	fmt.Println("Expected: 5")
	fmt.Println("Actual:", longestSubstringKDistinct(3, "cbbebi"))
	fmt.Println("Expected: 6")
	fmt.Println("Actual:", longestSubstringKDistinct(10, "cbbebi"))
}

//Question: longest substring with maximum k distinct characters
//Time: O(N)
//Space: O(M), where M is the # of distinct characters in str
func longestSubstringKDistinct(k int, str string) int {
	m := make(map[int]int)
	var start, longest int
	

	for i,v := range str{
		add(m, v)

		if len(m) > k {
			remove(m, rune(str[start]))
			start++
		}

		longest = max(longest, i-start+1)
	}
	
	return longest
}

func remove(m map[int]int, v rune){
	val := m[int(v)]
	if val == 1 {
		delete(m, int(v))
	} else {
		val--;
	}
}

func add(m map[int]int, v rune){
	val, found := m[int(v)]
	if found {
		val++
	} else {
		m[int(v)] = 1
	}
}

func max(x int, y int) int {
	if x > y{
		return x
	}
	return y
}