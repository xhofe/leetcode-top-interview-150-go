package main

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans := 0
	preNum := m[s[0]]
	for i := 1; i < len(s); i++ {
		num := m[s[i]]
		if preNum < num {
			ans -= preNum
		} else {
			ans += preNum
		}
		preNum = num
	}
	ans += preNum
	return ans
}
