package main

//Question: pair with target sum
//Time: O(N)
//Space: O(1)
func search(arr []int, targetSum int) []int {
	p1, p2 := 0, len(arr)-1
	for p1 < p2 {
		sum := arr[p1] + arr[p2]
		if sum < targetSum {
			p1++
		} else if sum > targetSum {
			p2--
		} else {
			return []int{p1, p2}
		}
	}
	return nil
}
