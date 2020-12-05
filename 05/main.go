package main

import (
	"github.com/csmith/aoc-2020/common"
	"strconv"
	"strings"
)

func main() {
	replacer := strings.NewReplacer("B", "1", "F", "0", "R", "1", "L", "0")
	input := common.ReadFileAsStrings("05/input.txt")
	max := 0
	var seats [1024]bool
	for i := range input {
		line := input[i]
		row, _ := strconv.ParseInt(replacer.Replace(line[0:7]), 2, 64)
		col, _ := strconv.ParseInt(replacer.Replace(line[7:10]), 2, 64)
		id := row*8 + col
		max = common.Max(max, int(id))
		seats[id] = true
	}
	println(max)

	for i := range seats {
		if i > 0 && !seats[i] && seats[i-1] && seats[i+1] {
			println(i)
			break
		}
	}
}
