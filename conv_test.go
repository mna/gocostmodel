package gocostmodel

import (
	"strconv"
	"testing"
)

var (
	s2i = "2"
	i2s string
)

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

func BenchmarkConvStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i1, _ = strconv.Atoi(s2i)
	}
}

func BenchmarkConvIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i2s = strconv.Itoa(i1)
	}
}
