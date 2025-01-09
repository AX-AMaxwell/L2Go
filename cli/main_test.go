package main

import "testing"

var (
	v1 = "0b1001010"
	v2 = "0x1001010"
)

func BenchmarkM1(b *testing.B) {
	for range b.N {
		isHex(v2)
	}
}

func BenchmarkM2(b *testing.B) {
	for range b.N {
		isBin(v1)
	}
}
