package main

import "fmt"

type interval struct {
	start int
	end   int
}

func main() {
	inputOneOne, inputOneTwo, inputOneThree := interval{1, 3}, interval{5, 6}, interval{7, 9}
	inputTwoOne, intervalTwoTwo := interval{2, 3}, interval{5, 7}
	inputOne := []interval{inputOneOne, inputOneTwo, inputOneThree}
	inputTwo := []interval{inputTwoOne, intervalTwoTwo}
	fmt.Println("Expected: [2,3] [5,6] [7,7]")
	fmt.Println("Actual:", merge(inputOne, inputTwo))
	inputOneOne, inputOneTwo, inputOneThree = interval{1, 3}, interval{5, 7}, interval{9, 12}
	inputTwoOne = interval{5, 10}
	inputOne = []interval{inputOneOne, inputOneTwo, inputOneThree}
	inputTwo = []interval{inputTwoOne}
	fmt.Println("Expected: [5,7] [9,10]")
	fmt.Println("Actual:", merge(inputOne, inputTwo))
}

//Question: intervals intersection
//Time: O(N+M)
//Space: O(N+M)
func merge(x, y []interval) []interval {
	intersections := make([]interval, 0)
	p1, p2 := 0, 0

	for p1 < len(x) && p2 < len(y) {
		intersect, interval := intersection(x[p1], y[p2])
		if intersect {
			intersections = append(intersections, interval)
		}
		if x[p1].end < y[p2].end {
			p1++
		} else {
			p2++
		}
	}

	return intersections
}

func intersection(x, y interval) (bool, interval) {
	if minInt(x.end, y.end)-maxInt(x.start, y.start) >= 0 {
		return true, interval{maxInt(x.start, y.start), minInt(x.end, y.end)}
	}
	return false, interval{-1, -1}
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
