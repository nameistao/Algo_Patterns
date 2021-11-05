package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	head := &node{value: 2}
	head.next = &node{value: 4}
	head.next.next = &node{value: 6}
	head.next.next.next = &node{value: 8}
	head.next.next.next.next = &node{value: 10}
	fmt.Println("Expected: 10 8 6 4 2")
	head = reverse(head)
	fmt.Print("Actual: ")
	for head != nil {
		fmt.Print(head.value, " ")
		head = head.next
	}
	fmt.Println()
}

//Question: reverse a linkedlist
//Time: O(N)
//Space: O(1)
func reverse(head *node) *node {
	var prev *node
	cur := head
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}
