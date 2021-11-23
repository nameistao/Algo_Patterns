package main

//Question: longest substring with same letters after replacement
//Time: O(N)
//Space: O(1)
func findLength(k int, str string) int {
	var (
		start    int
		maxCount int
		maxLen   int
		maxChar  int
	)

	m := make(map[int]int)

	for i, v := range str {
		_, found := m[int(v)]
		if found {
			m[int(v)]++
		} else {
			m[int(v)] = 1
		}
		if m[int(v)] >= maxCount {
			maxCount = m[int(v)]
			maxChar = int(v)
		}

		for i-start+1-maxCount > k {
			m[int(str[start])]--
			if int(str[start]) == maxChar {
				maxCount--
			}
			start++
		}

		maxLen = maxInt(maxLen, i-start+1)
	}
	return maxLen
}

func maxInt(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
