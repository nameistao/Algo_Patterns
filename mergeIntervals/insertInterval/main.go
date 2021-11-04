package main

import "fmt"

type interval struct {
	start int
	end   int
}

func main() {
	intervalOne, intervalTwo, intervalThree := interval{1, 3}, interval{5, 7}, interval{8, 12}
	intervals := []interval{intervalOne, intervalTwo, intervalThree}
	newInterval := interval{4, 6}
	fmt.Println("Expected: [1,3] [4,7] [8,12]")
	fmt.Println("Actual:", insert(intervals, newInterval))
	newInterval = interval{4, 10}
	fmt.Println("Expected: [1,3] [4,12]")
	fmt.Println("Actual:", insert(intervals, newInterval))
	intervalOne, intervalTwo = interval{2, 3}, interval{5, 7}
	intervals = []interval{intervalOne, intervalTwo}
	newInterval = interval{1, 4}
	fmt.Println("Expected: [1,4] [5,7]")
	fmt.Println("Actual:", insert(intervals, newInterval))
}

//Question: insert interval
//Time: O(N)
//Space: O(N)
func insert(intervals []interval, newInterval interval) []interval {
	for i := 0; i < len(intervals); i++ {
		if overlap(intervals[i], newInterval) {
			intervals[i].start = minInt(intervals[i].start, newInterval.start)
			intervals[i].end = maxInt(intervals[i].end, newInterval.end)
			break
		}
	}
	return merge(intervals)
}

func merge(intervals []interval) []interval {
	mergedIntervals := make([]interval, 0)
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		eval := intervals[i]
		if overlap(cur, eval) {
			cur.start = minInt(intervals[i].start, cur.start)
			cur.end = maxInt(intervals[i].end, cur.end)
		} else {
			mergedIntervals = append(mergedIntervals, cur)
			cur = eval
		}
	}
	mergedIntervals = append(mergedIntervals, cur)
	return mergedIntervals
}

func overlap(x, y interval) bool {
	if minInt(x.end, y.end)-maxInt(x.start, y.start) > 0 {
		return true
	}
	return false
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
