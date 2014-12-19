package gocostmodel

import "testing"

var (
	u1 uint
	u2 uint = 123
	u3 uint = 109239
)

func BenchmarkBitAnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u3 = u1 & u2
	}
}

func BenchmarkBitOr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u3 = u1 | u2
	}
}

func BenchmarkBitXor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u3 = u1 ^ u2
	}
}

func BenchmarkBitAndNot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u3 = u1 &^ u2
	}
}

func BenchmarkBitShiftLeft(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u3 = u1 << 2
	}
}

func BenchmarkBitShiftRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u3 = u1 >> 2
	}
}
