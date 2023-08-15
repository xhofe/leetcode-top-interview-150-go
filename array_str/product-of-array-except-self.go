package main

func productExceptSelf(nums []int) []int {
	leftProduct := make([]int, len(nums))
	rightProduct := make([]int, len(nums))
	leftProduct[0] = nums[0]
	rightProduct[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i]
		rightProduct[len(nums)-1-i] = rightProduct[len(nums)-i] * nums[len(nums)-1-i]
	}
	ans := make([]int, len(nums))
	ans[0], ans[len(nums)-1] = rightProduct[1], leftProduct[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		ans[i] = leftProduct[i-1] * rightProduct[i+1]
	}
	return ans
}
