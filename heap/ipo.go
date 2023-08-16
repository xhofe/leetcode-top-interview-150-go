package main

import (
	"container/heap"
	"sort"
)

type hp2 struct {
	sort.IntSlice
}

func (h *hp2) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

// Pop implements heap.Interface.
func (h *hp2) Pop() any {
	x := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return x
}

// Push implements heap.Interface.
func (h *hp2) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

var _ heap.Interface = (*hp2)(nil)

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	type pair struct{ profit, need int }
	var pairs []pair
	for i, p := range profits {
		pairs = append(pairs, pair{p, capital[i]})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].need < pairs[j].need })
	index := 0
	h := &hp2{}
	for k > 0 {
		for index < len(pairs) && pairs[index].need <= w {
			heap.Push(h, pairs[index].profit)
			index++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int)
		k--
	}
	return w
}
