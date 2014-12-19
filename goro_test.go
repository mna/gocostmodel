package gocostmodel

import (
	"math"
	"sync"
	"testing"
)

// this one doesn't work as intended, it presumably overwhelms the scheduler
// with tons of goros to spawn, so it ends up looking much more costly than
// it actually is.
func xBenchmarkGoroFireForget(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {
			// avoid inlining
			math.Pow(1, 1)
		}()
	}
}

func BenchmarkGoroWait(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			wg.Done()
		}()
		wg.Wait()
	}
}

func BenchmarkGoroSend(b *testing.B) {
	benchmarkGoroSend(b, 0)
}

func BenchmarkGoroSendBuf1(b *testing.B) {
	benchmarkGoroSend(b, 1)
}

func BenchmarkGoroSendBuf100(b *testing.B) {
	benchmarkGoroSend(b, 100)
}

func benchmarkGoroSend(b *testing.B, bufSize int) {
	ch := make(chan int, bufSize)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _ = range ch {
		}
		wg.Done()
	}()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ch <- i
	}

	b.StopTimer()
	close(ch)
	wg.Wait()
}
