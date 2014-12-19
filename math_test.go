package gocostmodel

import (
	"math"
	"math/rand"
	"testing"
)

var (
	neg    float64 = -3.456
	mathv1 float64 = 6.4
	mathv2 float64 = 1.2
	out    float64
	outi   int
)

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Abs(neg)
	}
}

func BenchmarkCos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Cos(mathv1)
	}
}

func BenchmarkSin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Sin(mathv1)
	}
}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Pow(mathv1, mathv2)
	}
}

func BenchmarkLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Log(mathv1)
	}
}

func BenchmarkExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Exp(mathv1)
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Sqrt(mathv1)
	}
}

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Max(mathv1, mathv2)
	}
}

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Min(mathv1, mathv2)
	}
}

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outi = rand.Int()
	}
}
