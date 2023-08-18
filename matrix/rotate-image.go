package main

func rotate(matrix [][]int) {
	n := len(matrix)
	pad := 0
	for n-pad*2 > 1 {
		for j := pad; j < n-pad-1; j++ {
			matrix[pad][j], matrix[n-1-j][pad], matrix[n-1-pad][n-1-j], matrix[j][n-1-pad] = matrix[n-1-j][pad], matrix[n-1-pad][n-1-j], matrix[j][n-1-pad], matrix[pad][j]
		}
		pad++
	}
}
