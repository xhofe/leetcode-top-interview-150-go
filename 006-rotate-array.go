package main

/*
nums = "----->-->"; k =3
result = "-->----->";

reverse "----->-->" we can get "<--<-----"
reverse "<--" we can get "--><-----"
reverse "<-----" we can get "-->----->"
*/
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i := 0; i < len(nums)>>1; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
