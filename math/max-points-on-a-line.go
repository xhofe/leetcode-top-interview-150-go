package main

func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}
	ans := 2
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			cnt := 2
			for k := j + 1; k < len(points); k++ {
				if isLine(points[i], points[j], points[k]) {
					cnt++
				}
			}
			ans = max(ans, cnt)
		}
	}
	return ans
}

func isLine(p1, p2, p3 []int) bool {
	return (p1[0]-p2[0])*(p1[1]-p3[1]) == (p1[0]-p3[0])*(p1[1]-p2[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
