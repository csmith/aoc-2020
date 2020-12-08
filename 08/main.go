package main

import (
	"github.com/csmith/aoc-2020/common"
	"regexp"
)

var tokens [][]string

func run(munge int) (terminated bool, acc int) {
	pc := 0
	seen := make(map[int]bool)
	for {
		if seen[pc] {
			return
		} else if pc >= len(tokens) {
			terminated = true
			return
		}

		seen[pc] = true
		instr := tokens[pc][0]

		if munge == pc {
			if instr == "nop" {
				instr = "jmp"
			} else if instr == "jmp" {
				instr = "nop"
			}
		}

		switch instr {
		case "nop":
			pc++
			continue
		case "jmp":
			pc += common.MustAtoi(tokens[pc][1])
			continue
		case "acc":
			acc += common.MustAtoi(tokens[pc][1])
			pc++
			continue
		}
	}
}

func main() {
	tokens = common.TokeniseLines(
		common.ReadFileAsStrings("08/input.txt"),
		regexp.MustCompile(`^(\S+) ([-+]\d+)$`),
	)

	_, acc := run(-1)
	println(acc)

	for i := range tokens {
		if tokens[i][0] != "acc" {
			if t, acc := run(i); t {
				println(acc)
				break
			}
		}
	}
}
