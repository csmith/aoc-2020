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

func Benchmark_runWithPresizedMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		starting := common.ReadCsvAsInts("input.txt")
		runWithPresizedMap(starting)
	}
}

func Benchmark_runWithIntArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		starting := common.ReadCsvAsInts("input.txt")
		runWithIntArray(starting)
	}
}

func Benchmark_runWithInt32Array(b *testing.B) {
	for i := 0; i < b.N; i++ {
		starting := common.ReadCsvAsInts("input.txt")
		runWithInt32Array(starting)
	}
}
