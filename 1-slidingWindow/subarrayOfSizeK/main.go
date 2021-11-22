package main

//Question: maximum subarray with a given sum
//Time: O(N)
//Space: O(1)
func findMaxSumSubArray(k int, arr []int) int {
	var (
		sum int
		max int
	)

	for i, v := range arr {
		sum += v

		if i+1-k > 0 {
			sum -= arr[i-k]
		}

		max = maxInt(max, sum)
	}
	return max
}

func maxInt(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
