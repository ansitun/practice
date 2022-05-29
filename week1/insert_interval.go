package main

import (
	"fmt"
)

func main() {
	//intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	intervals := make([][]int, 0)

	fmt.Printf("%v", insert(intervals, []int{4, 8}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	start := newInterval[0]
	end := newInterval[1]
	result := make([][]int, 0)
	i := 0
	tempInterval := []int{start, end}

	for ; i < len(intervals) && start > intervals[i][1]; i++ {
	}
	result = append(result, intervals[:i]...)

	for ; i < len(intervals) && end >= intervals[i][0]; i++ {
		start = min(start, intervals[i][0])
		end = max(end, intervals[i][1])
		tempInterval = []int{start, end}
	}
	result = append(result, tempInterval)

	result = append(result, intervals[i:]...)

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	} else {
		return a
	}
}

func min(a, b int) int {
	if b < a {
		return b
	} else {
		return a
	}
}
