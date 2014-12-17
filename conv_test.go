package gocostmodel

import "testing"

func BenchmarkConvIntToFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1 = float64(i1)
	}
}

func BenchmarkConvFloat64ToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i1 = int(d1)
	}
}

func BenchmarkConvIntToFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1 = float32(i1)
	}
}

func BenchmarkConvFloat32ToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i1 = int(f1)
	}
}
