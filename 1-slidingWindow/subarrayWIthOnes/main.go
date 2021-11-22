package main

//Question: longest subarray with ones after replacement
//Time: O(N)
//Space: O(1)
func findLength(k int, arr []int) int {
	var (
		ones  int
		max   int
		start int
	)

	for i, v := range arr {
		if v == 1 {
			ones++
		}

		for i-start+1-ones > k {
			if arr[start] == 1 {
				ones--
			}
			start++
		}

		if i-start+1 > max {
			max = i - start + 1
		}
	}
	return max
}
