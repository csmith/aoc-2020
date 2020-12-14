package main

import (
	"github.com/csmith/aoc-2020/common"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func part1(tokens [][]string) int64 {
	memory := make(map[int]int64)
	maskOn := int64(0)
	maskOff := int64(1)
	for i := range tokens {
		if tokens[i][0] == "mask" {
			maskOn, _ = strconv.ParseInt(strings.ReplaceAll(tokens[i][1], "X", "0"), 2, 64)
			maskOff, _ = strconv.ParseInt(strings.ReplaceAll(tokens[i][1], "X", "1"), 2, 64)
		} else if tokens[i][0] == "mem" {
			memory[common.MustAtoi(tokens[i][1])] = (int64(common.MustAtoi(tokens[i][2])) | maskOn) & maskOff
		} else {
			log.Panicf("Unknown instruction: %s", tokens[i])
		}
	}

	sum := int64(0)
	for _, v := range memory {
		sum += v
	}
	return sum
}

func part2(tokens [][]string) int {
	// Give up on using bitmasks how they're meant to be used, and do some horrific string munging instead. Sad.
	memory := make(map[int64]int)
	maskOn := int64(0)
	var floaters []int

	for i := range tokens {
		if tokens[i][0] == "mask" {
			maskOn, _ = strconv.ParseInt(strings.ReplaceAll(tokens[i][1], "X", "0"), 2, 64)
			floaters = []int{}
			for k, v := range tokens[i][1] {
				if v == 'X' {
					floaters = append(floaters, 35-k)
				}
			}

		} else if tokens[i][0] == "mem" {
			rawAddr := common.MustAtoi(tokens[i][1])
			value := common.MustAtoi(tokens[i][2])

			for m := int64(0); m < 1<<len(floaters); m++ {
				addr := int64(rawAddr) | maskOn
				for j := range floaters {
					addr = (addr & ^(1 << floaters[j])) | (((m >> j) & 1) << floaters[j])
					//      ^^^^^^^^^^^^^^^^^^^^^^^^^^^   ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
					//      Sets the bit of the address   Gets the desired state of the
					//      covered by the floater to 0   floater (the j'th bit of m) and
					//      so we can replace it with     shifts it into the desired
					//      its new value                 position
				}
				memory[addr] = value
			}
		} else {
			log.Panicf("Unknown instruction: %s", tokens[i])
		}
	}

	sum := 0
	for _, v := range memory {
		sum += v
	}
	return sum
}

func main() {
	tokens := common.TokeniseLines(
		common.ReadFileAsStrings("14/input.txt"),
		regexp.MustCompile(`(mask) = ([10X]+)|(mem)\[(\d+)] = (\d+)`),
	)

	println(part1(tokens))
	println(part2(tokens))
}
