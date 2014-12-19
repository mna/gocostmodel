package gocostmodel

import (
	"bytes"
	"testing"
)

var (
	bt1   = []byte("This is probably a Medium-Size source string")
	bt2   = []byte("z")
	bt3   = []byte("is")
	btix  int
	btout []byte
)

func BenchmarkByteIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		btix = bytes.Index(bt1, bt2)
	}
}

func BenchmarkByteReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		btout = bytes.Replace(bt1, bt3, bt2, -1)
	}
}

func BenchmarkByteToUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		btout = bytes.ToUpper(bt1)
	}
}

func BenchmarkByteToLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		btout = bytes.ToLower(bt1)
	}
}
