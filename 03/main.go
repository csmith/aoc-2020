package main

import "github.com/csmith/aoc-2020/common"

var data common.Map

const Tree common.Tile = '#'

func count(rowDelta, colDelta int) int {
	row := 0
	col := 0
	trees := 0
	for row < len(data)-rowDelta {
		row += rowDelta
		col += colDelta
		if data.TileAt(row, col) == Tree {
			trees++
		}
	}
	return trees
}

func main() {
	data = common.ReadFileAsMap("03/input.txt")
	part1 := count(1, 3)
	println(part1)
	println(part1 * count(1, 1) * count(1, 5) * count(1, 7) * count(2, 1))
}
