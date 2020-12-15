package main

import (
	"github.com/csmith/aoc-2020/common"
)

// This file contains four different implementations of the same algorithm with different backing data structures:
// A map, a map initialised with a large capacity, a fixed-size array of ints, and a fixed-size array of int32s.
//
// Benchmark results:
//
// $ go test -bench . -benchtime 10s
// goos: linux
// goarch: amd64
// pkg: github.com/csmith/aoc-2020/15
// Benchmark_runWithMap-8           	      6	1890266689 ns/op
// Benchmark_runWithPresizedMap-8   	      6	1996602376 ns/op
// Benchmark_runWithIntArray-8      	     25	496651941 ns/op
// Benchmark_runWithInt32Array-8    	     27	407223823 ns/op

func runWithMap(starting []int) (part1 int, part2 int) {
	lastSpoken := make(map[int]int)
	last := starting[len(starting) - 1]
	for i := 0; i < len(starting) - 1; i++ {
		lastSpoken[starting[i]] = i+1
	}
	for turn := len(starting)+1; turn <= 30000000; turn++ {
		age, ok := lastSpoken[last]
		lastSpoken[last] = turn-1

		if ok {
			last = turn - age - 1
		} else {
			last = 0
		}

		if turn == 2020 {
			part1 = last
		}
	}
	part2 = last
	return
}

func runWithPresizedMap(starting []int) (part1 int, part2 int) {
	lastSpoken := make(map[int]int, 30000000)
	last := starting[len(starting) - 1]
	for i := 0; i < len(starting) - 1; i++ {
		lastSpoken[starting[i]] = i+1
	}
	for turn := len(starting)+1; turn <= 30000000; turn++ {
		age, ok := lastSpoken[last]
		lastSpoken[last] = turn-1

		if ok {
			last = turn - age - 1
		} else {
			last = 0
		}

		if turn == 2020 {
			part1 = last
		}
	}
	part2 = last
	return
}

func runWithIntArray(starting []int) (part1 int, part2 int) {
	var lastSpoken [30000000]int
	last := starting[len(starting) - 1]
	for i := 0; i < len(starting) - 1; i++ {
		lastSpoken[starting[i]] = i+1
	}
	for turn := len(starting)+1; turn <= 30000000; turn++ {
		age := lastSpoken[last]
		lastSpoken[last] = turn-1

		if age > 0 {
			last = turn - age - 1
		} else {
			last = 0
		}

		if turn == 2020 {
			part1 = last
		}
	}
	part2 = last
	return
}

func runWithInt32Array(starting []int) (part1 int, part2 int) {
	var lastSpoken [30000000]int32
	last := starting[len(starting) - 1]
	for i := 0; i < len(starting) - 1; i++ {
		lastSpoken[starting[i]] = int32(i+1)
	}
	for turn := len(starting)+1; turn <= 30000000; turn++ {
		age := lastSpoken[last]
		lastSpoken[last] = int32(turn-1)

		if age > 0 {
			last = turn - int(age) - 1
		} else {
			last = 0
		}

		if turn == 2020 {
			part1 = last
		}
	}
	part2 = last
	return
}

func main() {
	p1, p2 := runWithInt32Array(common.ReadCsvAsInts("15/input.txt"))
	println(p1)
	println(p2)
}
