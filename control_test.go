package gocostmodel

import "testing"

var (
	x   int
	val = 1
	rng = []int{1}
)

func BenchmarkIfEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if val == 1 {
			x++
		}
	}
}

func BenchmarkIfNotEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i1 != 5 {
			x++
		}
	}
}

func BenchmarkFor3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < val; j++ {
			x++
		}
	}
}

func BenchmarkForRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// do not use empty range, would force go1.4
		for _ = range rng {
			x++
		}
	}
}

// simple sumx funcs probably get inlined

func BenchmarkFunc1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = sum1(i1)
	}
}

func BenchmarkFunc2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = sum2(i1, i2)
	}
}

func BenchmarkFunc3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = sum3(i1, i2, i3)
	}
}

func sum1(n int) int {
	return n
}

func sum2(n1, n2 int) int {
	return n1 + n2
}

func sum3(n1, n2, n3 int) int {
	return n1 + n2 + n3
}
