package main

import "fmt"

type node struct {
	val  int
	next *node
}

func main() {
	head := &node{val: 1}
	head.next = &node{val: 2}
	head.next.next = &node{val: 3}
	head.next.next.next = &node{val: 4}
	head.next.next.next.next = &node{val: 5}
	fmt.Println("Expected: 3")
	fmt.Println("Actual:", findMiddle(head))
	head.next.next.next.next.next = &node{val: 6}
	fmt.Println("Expected: 4")
	fmt.Println("Actual:", findMiddle(head))
	head.next.next.next.next.next.next = &node{val: 7}
	fmt.Println("Expected: 4")
	fmt.Println("Actual:", findMiddle(head))
}

//Question: middle of linkedlist
//Time: O(N)
//Space: O(1)
func findMiddle(head *node) int {
	fast := head
	for true {
		if fast == nil || fast.next == nil {
			return head.val
		}
		fast = fast.next.next
		head = head.next
	}
	return head.val
}
