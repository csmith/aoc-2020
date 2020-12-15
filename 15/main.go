package main

import (
	"github.com/csmith/aoc-2020/common"
)

func main() {
	starting := common.ReadCsvAsInts("15/input.txt")
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
			println(last)
		}
	}
	println(len(lastSpoken))
	println(last)
}
