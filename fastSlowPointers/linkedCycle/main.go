package main

import (
	"fmt"
)

type node struct {
	next *node
}

func main() {
	head := &node{}
	head.next = &node{}
	head.next.next = &node{}
	head.next.next.next = &node{}
	head.next.next.next.next = &node{}
	head.next.next.next.next.next = &node{}
	fmt.Println("Expected: false")
	fmt.Println("Actual:", calculateLength(head))
	head.next.next.next.next.next.next = head.next.next
	fmt.Println("Expected: true")
	fmt.Println("Actual:", calculateLength(head))
	head.next.next.next.next.next.next = head.next.next.next
	fmt.Println("Expected: true")
	fmt.Println("Actual:", calculateLength(head))
}

//Question: linkedlist cycle
//Time: O(N)
//Space: O(1)
func calculateLength(head *node) bool {
	if head == nil {
		return false
	}
	fast, slow := head, head

	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}
	}

	return false
}
