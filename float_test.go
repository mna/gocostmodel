package gocostmodel

import (
	"math"
	"testing"
)

var (
	f1 float32 = 1.23
	f2 float32 = 3.45
	f3 float32 = 6.78

	d1 float64 = 1.23
	d2 float64 = 3.45
	d3 float64 = 6.78
)

func BenchmarkFloat32Assign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1 = f2
	}
}

func BenchmarkFloat32Inc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1++
	}
}

func BenchmarkFloat32Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1 = f2 + f3
	}
}

func BenchmarkFloat32Sub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1 = f2 - f3
	}
}

func BenchmarkFloat32Mul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1 = f2 * f3
	}
}

func BenchmarkFloat32Mod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1 = float32(math.Mod(float64(f2), float64(f3)))
	}
}

func BenchmarkFloat64Assign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1 = d2
	}
}

func BenchmarkFloat64Inc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1++
	}
}

func BenchmarkFloat64Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1 = d2 + d3
	}
}

func BenchmarkFloat64Sub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1 = d2 - d3
	}
}

func BenchmarkFloat64Mul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1 = d2 * d3
	}
}

func BenchmarkFloat64Mod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1 = math.Mod(d2, d3)
	}
}
