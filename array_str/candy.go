package main

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		if i == 0 || ratings[i] <= ratings[i-1] {
			candies[i] = 1
		} else {
			candies[i] = candies[i-1] + 1
		}
	}
	ans := 0
	for i := len(ratings) - 1; i >= 0; i-- {
		if i != len(ratings)-1 && ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i],candies[i+1] + 1)
		}
		ans += candies[i]
	}
	return ans
}
