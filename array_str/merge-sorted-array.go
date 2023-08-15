package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1
	i, j := m-1, n-1
	for index >= 0 {
		if j < 0 || (i >= 0 && nums1[i] > nums2[j]) {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
		index--
	}
}
