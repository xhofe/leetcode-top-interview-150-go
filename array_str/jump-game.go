package main

func canJump(nums []int) bool {
	index := len(nums) - 1
	for j := len(nums) - 2; j >= 0; j-- {
		if j+nums[j] >= index {
			index = j
		}
	}
	return index == 0
}
