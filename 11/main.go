package main

import (
	"github.com/csmith/aoc-2020/common"
)

const Floor common.Tile = '.'
const Empty common.Tile = 'L'
const Occupied common.Tile = '#'

func main() {
	original := common.ReadFileAsMap("11/input.txt")
	part1 := 0
	state := original

	for {
		newState := common.Map{}
		changed := false
		for y := range state {
			newState = append(newState, []common.Tile{})
			for x := range state[y] {
				if state[y][x] == Floor {
					newState[y] = append(newState[y], state[y][x])
					continue
				}

				occupied := 0
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if dy == 0 && dx == 0 {
							continue
						}
						if y+dy >= 0 && x+dx >= 0 && y+dy < len(state) && x+dx < len(state[y+dy]) && state[y+dy][x+dx] == Occupied {
							occupied++
						}
					}
				}
				if state[y][x] == Empty && occupied == 0 {
					newState[y] = append(newState[y], Occupied)
					changed = true
				} else if state[y][x] == Occupied && occupied >= 4 {
					newState[y] = append(newState[y], Empty)
					changed = true
				} else {
					newState[y] = append(newState[y], state[y][x])
				}
			}
		}
		state = newState
		if !changed {
			for y := range state {
				for x := range state[y] {
					if state[y][x] == Occupied {
						part1++
					}
				}
			}
			println(part1)
			break
		}
	}

	// -------

	state = original
	part2 := 0
	for {
		newState := common.Map{}
		changed := false
		for y := range state {
			newState = append(newState, []common.Tile{})
			for x := range state[y] {
				if state[y][x] == Floor {
					newState[y] = append(newState[y], state[y][x])
					continue
				}

				occupied := 0
				dirs := [][]int{
					{-1, -1},
					{-1, 0},
					{-1, 1},
					{0, -1},
					{0, 1},
					{1, -1},
					{1, 0},
					{1, 1},
				}
				for _, d := range dirs {
					nx := x + d[0]
					ny := y + d[1]
					for nx >= 0 && ny >= 0 && ny < len(state) && nx < len(state[ny]) {
						if state[ny][nx] == Occupied {
							occupied++
							break
						} else if state[ny][nx] == Empty {
							break
						} else {
							nx += d[0]
							ny += d[1]
						}
					}
				}

				if state[y][x] == Empty && occupied == 0 {
					newState[y] = append(newState[y], Occupied)
					changed = true
				} else if state[y][x] == Occupied && occupied >= 5 {
					newState[y] = append(newState[y], Empty)
					changed = true
				} else {
					newState[y] = append(newState[y], state[y][x])
				}
			}
		}
		state = newState
		if !changed {
			for y := range state {
				for x := range state[y] {
					if state[y][x] == Occupied {
						part2++
					}
				}
			}
			println(part2)
			break
		}
	}

}
