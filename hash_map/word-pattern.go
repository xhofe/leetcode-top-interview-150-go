package main

import "strings"

func wordPattern(pattern string, s string) bool {
	segments := strings.Split(s, " ")
	if len(segments) != len(pattern) {
		return false
	}
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		b := pattern[i]
		seg := segments[i]
		if v, ok := m1[b]; ok && v != seg {
			return false
		}
		if v, ok := m2[seg]; ok && v != b {
			return false
		}
		m1[b] = seg
		m2[seg] = b
	}
	return true
}
