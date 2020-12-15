package main

import (
	"github.com/csmith/aoc-2020/common"
	"testing"
)

func Benchmark_runWithMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		starting := common.ReadCsvAsInts("input.txt")
		runWithMap(starting)
	}
}

func Benchmark_runWithArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		starting := common.ReadCsvAsInts("input.txt")
		runWithArray(starting)
	}
}
