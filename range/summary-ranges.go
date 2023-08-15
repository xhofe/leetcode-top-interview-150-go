package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	ans := make([]string, 0)
	start := 0
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i] != nums[i-1]+1 {
			if start == i-1 {
				ans = append(ans, strconv.Itoa(nums[start]))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
			}
			start = i
		}
	}
	return ans
}
