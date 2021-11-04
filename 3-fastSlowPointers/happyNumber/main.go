package main

import "fmt"

func main() {
	fmt.Println("Expected: true")
	fmt.Println("Actual:", find(23))
	fmt.Println("Expected: false")
	fmt.Println("Actual:", find(12))
}

//Question: happy number
//Time: O(log(N))
//Space: O(1)
func find(num int) bool {
	fast, slow := num, num
	for fast != 1 {
		fast = next(next(fast))
		slow = next(slow)
		if fast == slow && fast != 1 {
			return false
		}
	}
	return true
}

func next(num int) (sum int) {
	for num > 0 {
		sum += (num % 10) * (num % 10)
		num /= 10
	}
	return
}
