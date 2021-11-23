package main

//Question: longest substring with maximum k distinct characters
//Time: O(N)
//Space: O(M), where M is the # of distinct characters in str
func longestSubstringKDistinct(k int, str string) int {
	m := make(map[int]int)
	var start, longest int

	for i, v := range str {
		add(m, v)

		if len(m) > k {
			remove(m, rune(str[start]))
			start++
		}

		longest = max(longest, i-start+1)
	}

	return longest
}

func remove(m map[int]int, v rune) {
	val := m[int(v)]
	if val == 1 {
		delete(m, int(v))
	} else {
		val--
	}
}

func add(m map[int]int, v rune) {
	val, found := m[int(v)]
	if found {
		val++
	} else {
		m[int(v)] = 1
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
