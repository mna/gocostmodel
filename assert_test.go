package gocostmodel

import (
	"testing"
)

type assertStr struct {
	a1 int
}

var (
	assertPiface interface{} = &assertStr{}
	assertIface  interface{} = assertStr{}
	assertStra   assertStr
	assertStrp   *assertStr
)

var ok bool

func BenchmarkIfacePtrTypeAssertShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assertStrp = assertPiface.(*assertStr)
	}
}

func BenchmarkIfacePtrTypeAssertWithOk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assertStrp, ok = assertPiface.(*assertStr)
	}
}

func BenchmarkIfaceTypeAssertShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assertStra = assertIface.(assertStr)
	}
}

func BenchmarkIfaceTypeAssertWithOk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assertStra, ok = assertIface.(assertStr)
	}
}
