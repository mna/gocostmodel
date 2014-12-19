package gocostmodel

import (
	"strconv"
	"testing"
)

var (
	int5    []int
	int100  []int
	str5    []string
	str100  []string
	mint5   = make(map[int]string)
	mint100 = make(map[int]string)
	mstr5   = make(map[string]string)
	mstr100 = make(map[string]string)
	mapOut  string
)

func init() {
	for i := 0; i < 100; i++ {
		key := "key " + strconv.Itoa(i)
		val := "value for " + strconv.Itoa(i)
		if i < 5 {
			int5 = append(int5, i)
			str5 = append(str5, key)
			mint5[i] = key
			mstr5[key] = val
		}
		int100 = append(int100, i)
		str100 = append(str100, key)
		mint100[i] = key
		mstr100[key] = val
	}
}

func BenchmarkFindSliceInt5(b *testing.B) {
	key := 4
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(int5); j++ {
			if int5[j] == key {
				break
			}
		}
	}
}

func BenchmarkFindSliceStr5(b *testing.B) {
	key := "key 4"
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(str5); j++ {
			if str5[j] == key {
				break
			}
		}
	}
}

func BenchmarkFindSliceInt100(b *testing.B) {
	key := 99
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(int100); j++ {
			if int100[j] == key {
				break
			}
		}
	}
}

func BenchmarkFindSliceStr100(b *testing.B) {
	key := "key 99"
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(str100); j++ {
			if str100[j] == key {
				break
			}
		}
	}
}

func BenchmarkFindMapInt5(b *testing.B) {
	key := 4
	for i := 0; i < b.N; i++ {
		mapOut = mint5[key]
	}
}

func BenchmarkFindMapStr5(b *testing.B) {
	key := "key 4"
	for i := 0; i < b.N; i++ {
		mapOut = mstr5[key]
	}
}

func BenchmarkFindMapInt100(b *testing.B) {
	key := 99
	for i := 0; i < b.N; i++ {
		mapOut = mint100[key]
	}
}

func BenchmarkFindMapStr100(b *testing.B) {
	key := "key 99"
	for i := 0; i < b.N; i++ {
		mapOut = mstr100[key]
	}
}
