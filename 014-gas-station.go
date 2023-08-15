package main

func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		left, cnt := 0, 0
		for cnt < n {
			j := (i + cnt) % n
			left += gas[j] - cost[j]
			if left < 0 {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}
