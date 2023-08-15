package main

import "sort"

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	for i, v := range citations {
		if v < i + 1 {
			return i
		}
	}
	return len(citations)
}
