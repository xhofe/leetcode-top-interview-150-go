package main

func trap(height []int) int {
	stack := []int{-1}
	ans := 0
	for i, v := range height {
		for len(stack) > 1 && v >= height[stack[len(stack)-1]] {
			if len(stack) > 2 {
				ans += (i - stack[len(stack)-2] - 1) * (min(v, height[stack[len(stack)-2]]) - height[stack[len(stack)-1]])
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

