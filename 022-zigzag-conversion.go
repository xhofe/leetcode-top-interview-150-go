package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([][]byte, numRows)
	down := true
	i := 0
	for j := range s {
		rows[i] = append(rows[i], s[j])
		if down {
			i++
			if i == numRows-1 {
				down = false
			}
		} else {
			i--
			if i == 0 {
				down = true
			}
		}
	}
	ans := make([]byte, 0, len(s))
	for _, v := range rows {
		ans = append(ans, v...)
	}
	return string(ans)
}
