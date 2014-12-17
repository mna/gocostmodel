package gocostmodel

import "testing"

var (
	i1 int = 1
	i2 int = 2
	i3 int = 5
)

func BenchmarkIntAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i1 = i2
	}
}

func BenchmarkIntInc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i1++
	}
}

func BenchmarkIntAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i1 = i2 + i3
	}
}

func BenchmarkIntSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i1 = i2 - i3
	}
}

func BenchmarkIntMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i1 = i2 * i3
	}
}

func BenchmarkIntDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i1 = i2 / i3
	}
}

func BenchmarkIntMod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i1 = i2 % i3
	}
}
