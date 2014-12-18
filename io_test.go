package gocostmodel

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type stringer int

func (s stringer) String() string {
	return "a"
}

var (
	st  stringer = 1
	bl           = true
	ptr          = &st
)

// fmt.Fprintx(ioutil.Discard, ...) benchmarks the fmt mechanism, but not
// the actual I/O, since Discard doesn't write anywhere.

func BenchmarkPrintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintln(ioutil.Discard, "test")
	}
}

func BenchmarkEmptyPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprint(ioutil.Discard, "")
	}
}

func BenchmarkPrintPctd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%d", i)
	}
}

func BenchmarkPrintPctb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%b", i)
	}
}

func BenchmarkPrintPctc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%c", i1)
	}
}

func BenchmarkPrintPcto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%o", i)
	}
}

func BenchmarkPrintPctxInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%x", i)
	}
}

func BenchmarkPrintPctU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%U", i1)
	}
}

func BenchmarkPrintPcts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%s", "a")
	}
}

func BenchmarkPrintPctq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%q", "a")
	}
}

func BenchmarkPrintPctxString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%x", "a")
	}
}

func BenchmarkPrintPctxByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%x", 'a')
	}
}

func BenchmarkPrintPctt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%t", bl)
	}
}

func BenchmarkPrintPctT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%T", "a")
	}
}

func BenchmarkPrintPctf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%f", d1)
	}
}

func BenchmarkPrintPcte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%e", d1)
	}
}

func BenchmarkPrintPctp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%p", ptr)
	}
}

func BenchmarkPrintStringer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%s", st)
	}
}
