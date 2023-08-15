package main

import "strings"

func reverseWords(s string) string {
	segments := strings.Split(s, " ")
	ans := make([]string, 0)
	for i := len(segments) - 1; i >= 0; i-- {
		if segments[i] == "" {
			continue
		}
		ans = append(ans, segments[i])
	}
	return strings.Join(ans, " ")
}
