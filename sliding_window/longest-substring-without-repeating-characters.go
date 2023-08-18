package main

func lengthOfLongestSubstring(s string) int {
	ans := 0
	cnt := make(map[byte]int)
	l := -1
	for i, b := range []byte(s) {
		if j, ok := cnt[b]; ok {
			l = max(l, j)
		}
		cnt[b] = i
		ans = max(ans, i-l)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
