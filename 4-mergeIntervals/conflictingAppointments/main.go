package main

import (
	"fmt"
	"sort"
)

type interval struct {
	start int
	end   int
}

func main() {
	intervals := []interval{{1, 4}, {2, 5}, {7, 9}}
	fmt.Println("Expected: false")
	fmt.Println("Actual:", canAttendAllAppointments(intervals))
	intervals = []interval{{6, 7}, {2, 4}, {8, 12}}
	fmt.Println("Expected: true")
	fmt.Println("Actual:", canAttendAllAppointments(intervals))
	intervals = []interval{{4, 5}, {2, 3}, {3, 6}}
	fmt.Println("Expected: false")
	fmt.Println("Actual:", canAttendAllAppointments(intervals))
}

//Question: conflicting appointments
//Time: O(N)
//Space: O(N)
func canAttendAllAppointments(intervals []interval) bool {
	sort.Slice(intervals, func(x, y int) bool {
		return intervals[x].start < intervals[y].start
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i].start <= intervals[i-1].end {
			return false
		}
	}
	return true
}
