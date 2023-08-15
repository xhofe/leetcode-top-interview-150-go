package main

func removeDuplicatesII(nums []int) int {
	index := 1
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[index-1] && cnt > 1 {
			continue
		}
		if nums[i] != nums[index-1] {
			cnt = 1
		} else {
			cnt++
		}
		nums[index] = nums[i]
		index++
	}
	return index
}
