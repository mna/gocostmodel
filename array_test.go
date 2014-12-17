package gocostmodel

import "testing"

var (
	v1  = [...]int{1, 2, 3, 4, 5}
	v2  = [...]int{1, 2, 3, 4, 5}
	v3  = [...]int{1, 2, 3, 4, 5}
	sl1 = []int{5, 6, 7}
)

func BenchmarkArrayIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1[i1] = i1
	}
}

func BenchmarkArrayIndex2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1[v2[i1]] = i1
	}
}

func BenchmarkArrayIndex3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1[v2[v3[i1]]] = i1
	}
}

func BenchmarkArraySlicing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sl1 = v1[:]
	}
}
