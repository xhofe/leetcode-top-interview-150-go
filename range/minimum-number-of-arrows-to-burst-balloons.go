package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	ans := 1
	pos := points[0][1]
	for _, v := range points {
		if v[0] <= pos {
			pos = min(pos, v[1])
			continue
		}
		ans++
		pos = v[1]
	}
	return ans
}
