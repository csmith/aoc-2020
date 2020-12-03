package main

import "github.com/csmith/aoc-2019/common"

var lines []string

func count(rowDelta, colDelta int) int {
	row := 0
	col := 0
	trees := 0
	for row < len(lines)-rowDelta {
		row += rowDelta
		col = (col + colDelta) % len(lines[0])
		if lines[row][col] == '#' {
			trees++
		}
	}
	return trees
}

func main() {
	lines = common.ReadFileAsStrings("03/input.txt")
	part1 := count(1, 3)
	println(part1)
	println(part1 * count(1, 1) * count(1, 5) * count(1, 7) * count(2, 1))
}
