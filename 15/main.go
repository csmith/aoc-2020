package main

import (
	"github.com/csmith/aoc-2020/common"
)

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

func runWithArray(starting []int) (part1 int, part2 int) {
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

func main() {
	p1, p2 := runWithArray(common.ReadCsvAsInts("15/input.txt"))
	println(p1)
	println(p2)
}
