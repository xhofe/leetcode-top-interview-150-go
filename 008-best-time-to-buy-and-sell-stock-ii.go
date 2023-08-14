package main

func maxProfitII(prices []int) int {
	ans := 0
	for i, v := range prices {
		if i > 0 && v > prices[i-1] {
			ans += v - prices[i-1]
		}
	}
	return ans
}
