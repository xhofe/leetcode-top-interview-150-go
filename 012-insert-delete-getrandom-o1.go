package main

import "math/rand"

type RandomizedSet struct {
	m map[int]int
	l []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		l: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.l = append(this.l, val)
	this.m[val] = len(this.l) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.m[val]; ok {
		delete(this.m, val)
		if i != len(this.l)-1 {
			this.l[i] = this.l[len(this.l)-1]
			this.m[this.l[len(this.l)-1]] = i
		}
		this.l = this.l[:len(this.l)-1]
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Int() % len(this.l)
	return this.l[idx]
}
