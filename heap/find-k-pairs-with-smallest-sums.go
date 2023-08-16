package main

import "container/heap"

type pair struct{ i, j, val int }
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := &hp{}
	for i := 0; i < k && i < m; i++ {
		*h = append((*h), pair{i, 0, nums1[i] + nums2[0]})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(h, pair{i, j + 1, nums1[i] + nums2[j+1]})
		}
	}
	return
}
