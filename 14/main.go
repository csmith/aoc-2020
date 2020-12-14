package main

import (
	"fmt"
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
	mask := ""
	for i := range tokens {
		if tokens[i][0] == "mask" {
			mask = tokens[i][1]
		} else if tokens[i][0] == "mem" {
			addresses := [][]rune{[]rune(fmt.Sprintf("%036b", int64(common.MustAtoi(tokens[i][1]))))}
			for n, m := range mask {
				if m == '0' {
					continue
				}

				// Set the nth byte in all addresses to 1
				for a := range addresses {
					addresses[a][n] = '1'
				}

				if m == 'X' {
					// For Xs, also create a copy with the nth byte set to 0
					length := len(addresses)
					for a := 0; a < length; a++ {
						address := make([]rune, 36)
						copy(address, addresses[a])
						address[n] = '0'
						addresses = append(addresses, address)
					}
				}
			}

			for a := range addresses {
				addr, _ := strconv.ParseInt(string(addresses[a]), 2, 64)
				memory[addr] = common.MustAtoi(tokens[i][2])
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
