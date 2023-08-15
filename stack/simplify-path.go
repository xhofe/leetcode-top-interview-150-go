package main

import "strings"

func simplifyPath(path string) string {
	segments := strings.Split(path, "/")
	ans := []string{}
	for _, seg := range segments {
		switch seg {
		case "", ".":
			continue
		case "..":
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
		default:
			ans = append(ans, seg)
		}
	}
	return "/" + strings.Join(ans, "/")
}
