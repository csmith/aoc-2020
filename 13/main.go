package main

import (
	"github.com/csmith/aoc-2020/common"
	"math"
	"strings"
)

func main() {
	lines := common.ReadFileAsStrings("13/input.txt")
	arrivalTime := common.MustAtoi(lines[0])
	busIds := strings.Split(lines[1], ",")

	bestId := 0
	smallestWait := math.MaxInt64
	for i := range busIds {
		if busIds[i] == "x" {
			// "Out of service". Sure.
			continue
		}
		id := common.MustAtoi(busIds[i])
		wait := id - (arrivalTime % id)
		if wait < smallestWait {
			bestId = id
			smallestWait = wait
		}
	}
	println(bestId * smallestWait)

	skip := 1
	pinned := 0
	compTime := 0
	for pinned < len(busIds) {
		compTime += skip
		pinned = 0
		skip = 1

		for i := range busIds {
			if busIds[i] == "x" {
				pinned++
				continue
			}
			id := common.MustAtoi(busIds[i])
			if (compTime + i) % id == 0 {
				pinned++
				skip *= id
			} else {
				break
			}
		}
	}

	println(compTime)
}
