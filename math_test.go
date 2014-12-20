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

func BenchmarkMathAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Abs(neg)
	}
}

func BenchmarkMathCos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Cos(mathv1)
	}
}

func BenchmarkMathSin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Sin(mathv1)
	}
}

func BenchmarkMathPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Pow(mathv1, mathv2)
	}
}

func BenchmarkMathLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Log(mathv1)
	}
}

func BenchmarkMathExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Exp(mathv1)
	}
}

func BenchmarkMathSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Sqrt(mathv1)
	}
}

func BenchmarkMathMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Max(mathv1, mathv2)
	}
}

func BenchmarkMathMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = math.Min(mathv1, mathv2)
	}
}

func BenchmarkMathRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		outi = rand.Int()
	}
}
