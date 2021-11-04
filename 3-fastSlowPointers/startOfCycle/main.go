package main

import "fmt"

type node struct {
	next *node
	val  int
}

func main() {
	head := &node{val: 1}
	head.next = &node{val: 2}
	head.next.next = &node{val: 3}
	head.next.next.next = &node{val: 4}
	head.next.next.next.next = &node{val: 5}
	head.next.next.next.next.next = &node{val: 6}
	head.next.next.next.next.next.next = head.next.next
	fmt.Println("Expected: 3")
	fmt.Println("Actual:", findCycleStart(head))
	head.next.next.next.next.next.next = head.next.next.next
	fmt.Println("Expected: 4")
	fmt.Println("Actual:", findCycleStart(head))
	head.next.next.next.next.next.next = head
	fmt.Println("Expected: 1")
	fmt.Println("Actual:", findCycleStart(head))
}

//Question: start of linkedlist cycle
//Time: O(N)
//Space: O(1)
func findCycleStart(head *node) int {
	fast, slow := head, head
	for ok := true; ok; ok = (fast != slow) {
		fast = fast.next.next
		slow = slow.next
	}

	cycleLen, cur := 0, slow
	for ok := true; ok; ok = (cur != slow) {
		cur = cur.next
		cycleLen++
	}

	fast, slow = head, head
	for i := 0; i < cycleLen; i++ {
		fast = fast.next
	}

	for fast != slow {
		fast = fast.next
		slow = slow.next
	}

	return fast.val
}
