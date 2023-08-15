package main

import "sort"

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	last := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if cur[0] <= last[1] {
			last[1] = max(last[1], cur[1])
		} else {
			ans = append(ans, last)
			last = cur
		}
	}
	ans = append(ans, last)
	return ans
}
