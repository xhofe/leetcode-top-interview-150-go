package main

import "math/bits"

func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}
