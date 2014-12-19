package gocostmodel

import (
	"net/http"
	"testing"
	"time"
)

var (
	m    map[string]string
	sl   []string
	ch   chan string
	ps   *string
	ptm  *time.Time
	tm   time.Time
	preq *http.Request
	req  http.Request
	fn   func()
	cln  int
)

func BenchmarkMakeMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m = make(map[string]string)
	}
}

func BenchmarkMakeMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m = make(map[string]string, 100)
	}
}

func BenchmarkMakeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sl = make([]string, 0)
	}
}

func BenchmarkMakeSliceLen100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sl = make([]string, 100)
	}
}

func BenchmarkMakeSliceCap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sl = make([]string, 0, 100)
	}
}

func BenchmarkMakeChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch = make(chan string)
	}
}

func BenchmarkMakeChan1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch = make(chan string, 1)
	}
}

func BenchmarkMakeChan100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch = make(chan string, 100)
	}
}

func BenchmarkNewString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ps = new(string)
	}
}

func BenchmarkNewSmallStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ptm = new(time.Time)
	}
}

func BenchmarkNewSmallStructLit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tm = time.Time{}
	}
}

func BenchmarkNewSmallStructLitP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ptm = &time.Time{}
	}
}

func BenchmarkNewBigStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		preq = new(http.Request)
	}
}

func BenchmarkNewBigStructLit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req = http.Request{}
	}
}

func BenchmarkNewBigStructLitP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		preq = &http.Request{}
	}
}

func BenchmarkNewClosure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn = func() {
			cln = i
		}
	}
}
