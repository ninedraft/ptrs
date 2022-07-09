package ptrs_test

import (
	"ptrs"
	"testing"
)

const N = 1024

type stub [32]byte

func allocNaive[E any](n int) []*E {
	var ptrs = make([]*E, n)
	for i := range ptrs {
		ptrs[i] = new(E)
	}
	return ptrs
}

var naiveValues []*stub

func BenchmarkNaive(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		naiveValues = allocNaive[stub](N)
	}
}

var makeValues []*stub

func BenchmarkMake(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		makeValues = ptrs.Make[stub](N, nil)
	}
}

var makeModifyValues []*stub

func BenchmarkModifyMake(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		makeModifyValues = ptrs.Make(N, func(i int, ptr *stub) {
			ptr[0] = byte(i)
		})
	}
}
