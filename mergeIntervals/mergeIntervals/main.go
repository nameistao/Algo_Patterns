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
	intervalOne, intervalTwo, intervalThree := interval{1, 4}, interval{2, 5}, interval{7, 9}
	intervals := []interval{intervalOne, intervalTwo, intervalThree}
	fmt.Println("Expected: [1,5] [7,9]")
	fmt.Println("Actual:", merge(intervals))
	intervalOne, intervalTwo, intervalThree = interval{6, 7}, interval{2, 4}, interval{5, 9}
	intervals = []interval{intervalOne, intervalTwo, intervalThree}
	fmt.Println("Expected: [2,4] [5,9]")
	fmt.Println("Actual:", merge(intervals))
	intervalOne, intervalTwo, intervalThree = interval{1, 4}, interval{2, 6}, interval{3, 5}
	intervals = []interval{intervalOne, intervalTwo, intervalThree}
	fmt.Println("Expected: [1,6]")
	fmt.Println("Actual:", merge(intervals))
}

//Question: merge intervals
//Time: O(NlogN)
//Space: O(N)
func merge(intervals []interval) []interval {
	mergedIntervals := make([]interval, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		eval := intervals[i]
		if cur.end >= eval.start {
			cur.end = maxInt(cur.end, eval.end)
		} else {
			mergedIntervals = append(mergedIntervals, cur)
			cur = eval
		}
	}
	mergedIntervals = append(mergedIntervals, cur)
	return mergedIntervals
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
