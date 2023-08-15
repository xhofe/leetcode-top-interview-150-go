package main

func maxProfit(prices []int) int {
	mi := prices[0]
	ans := 0
	for _, v := range prices {
		ans = max(ans, v-mi)
		mi = min(mi, v)
	}
	return ans
}
