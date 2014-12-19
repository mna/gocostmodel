package gocostmodel

import "testing"

var (
	x         int
	val       = 1
	rng       = []int{1}
	blockedCh = make(chan bool)
	closedCh  = make(chan bool)
)

func init() {
	close(closedCh)
}

func BenchmarkSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch val {
		case 0:
			x++
		case 1:
			x--
		default:
			x += 2
		}
	}
}

func BenchmarkIfEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if val == 1 {
			x++
		}
	}
}

func BenchmarkIfNotEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if val != 5 {
			x++
		}
	}
}

func BenchmarkFor1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for val < 0 {
			x++
		}
	}
}

func BenchmarkFor3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// loop only once, the goal is to measure the loop construct
		for j := 0; j < val; j++ {
			x++
		}
	}
}

func BenchmarkForRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// do not use empty range, would force go1.4
		for _ = range rng {
			x++
		}
	}
}

func BenchmarkForRangeClosedChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// do not use empty range, would force go1.4
		for _ = range closedCh {
			x++
		}
	}
}

func BenchmarkSelectBlockedDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case <-blockedCh:
			x--
		default:
			x++
		}
	}
}

func BenchmarkSelectBlockedClosed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case <-blockedCh:
			x--
		case <-closedCh:
			x++
		}
	}
}

// simple sumx funcs get inlined (go test -c -gcflags -m)

func BenchmarkFunc1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = sum1(i1)
	}
}

func BenchmarkFunc2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = sum2(i1, i2)
	}
}

func BenchmarkFunc3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = sum3(i1, i2, i3)
	}
}

func sum1(n int) int {
	x := n
	return x
}

func sum2(n1, n2 int) int {
	x := n1 + n2
	return x
}

func sum3(n1, n2, n3 int) int {
	x := n1 + n2 + n3
	return x
}
