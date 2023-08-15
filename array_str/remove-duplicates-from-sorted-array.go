package main

func removeDuplicates(nums []int) int {
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[index-1] {
			continue
		}
		nums[index] = nums[i]
		index++
	}
	return index
}
