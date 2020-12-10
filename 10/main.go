package main

import (
	"github.com/csmith/aoc-2020/common"
	"sort"
)

func combinations(from int, chunk []int) int {
	if len(chunk) == 0 {
		return 1
	}

	res := 0
	for i := range chunk {
		if chunk[i] - from <= 3 {
			res += combinations(chunk[i], chunk[i+1:])
		}
	}
	return res
}

func main() {
	lines := common.ReadFileAsInts("10/input.txt")
	sort.Ints(lines)

	var chunks [][]int
	lastChunk := 0

	var (
		last   = 0
		ones   = 0
		threes = 1 // The built-in adapter at the end is always a jump of three
	)
	for i := range lines {
		if lines[i]-last == 1 {
			ones++
		} else if lines[i]-last == 3 {
			threes++

			chunks = append(chunks, lines[lastChunk:i])
			lastChunk = i
		}
		last = lines[i]
	}

	println(ones * threes)

	// Any jump of 3 must be in our chain, as there's no other way to bridge the gap. The total number of combinations
	// is therefore the product of the combinations within each 3-separated chunk, which turns out to be a manageable
	// calculation.
	last = 0
	total := 1
	for i := range chunks {
		total *= combinations(last, chunks[i])
		last = chunks[i][len(chunks[i])-1]
	}
	println(total)
}
