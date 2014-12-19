package gocostmodel

import (
	"strings"
	"testing"
)

var (
	str1   = "This is probably a Medium-Size source string"
	str2   = "z"
	str3   = "is"
	strix  int
	strout string
)

func BenchmarkStringIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strix = strings.Index(str1, str2)
	}
}

func BenchmarkStringReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strout = strings.Replace(str1, str3, str2, -1)
	}
}

func BenchmarkStringToUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strout = strings.ToUpper(str1)
	}
}

func BenchmarkStringToLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strout = strings.ToLower(str1)
	}
}
