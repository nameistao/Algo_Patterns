package main

import (
	"fmt"
)

func main(){
	fmt.Println("Expected: 9")
	fmt.Println("Actual:", findMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2 }));
	fmt.Println("Expected: 7")
	fmt.Println("Actual:", findMaxSumSubArray(2, []int{2, 3, 4, 1, 5 }))
}

//Question: maximum subarray with a given sum
func findMaxSumSubArray(k int, arr []int) int {
	var (
		sum int
		max int
	)
	
	for i,v := range arr {
		sum += v

		if(i+1-k > 0){
			sum -= arr[i-k]	
		}

		max = maxInt(max, sum)
	}
	return max
}

func maxInt(x int, y int) int {
	if x > y{
		return x
	} 
	return y
}