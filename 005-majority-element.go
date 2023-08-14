package main

func majorityElement(nums []int) int {
	ans := nums[0]
	cnt := 0
	for _, v := range nums {
		if v == ans {
			cnt++
		} else {
			cnt--
		}
		if cnt <= 0 {
			ans = v
			cnt = 1
		}
	}
	return ans
}
