package gocostmodel

import (
	"reflect"
	"testing"
)

type reflectStr struct {
	a1 int
}

var (
	reflectStra    = reflectStr{}
	reflectStrp    = &reflectStr{}
	reflectStrVal  = reflect.ValueOf(reflectStr{1})
	reflectStrPVal = reflect.ValueOf(&reflectStr{1})
)

func BenchmarkReflFromVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflectStra = reflectStrVal.Interface().(reflectStr)
	}
}

func BenchmarkReflToVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflectStrVal = reflect.ValueOf(reflectStra)
	}
}

func BenchmarkReflPtrFromVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflectStrp = reflectStrPVal.Interface().(*reflectStr)
	}
}

func BenchmarkReflPtrToVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflectStrPVal = reflect.ValueOf(reflectStrp)
	}
}
