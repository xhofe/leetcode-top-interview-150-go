package main

func jump(nums []int) int {
	start, end := 0, 1
	ans := 0
	for end < len(nums) {
		_end := end
		for i := start; i < _end; i++ {
			end = max(end, i+nums[i]+1)
		}
		ans++
		start = _end
	}
	return ans
}
