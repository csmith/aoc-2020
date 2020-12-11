package main

import (
	"github.com/csmith/aoc-2020/common"
)

const Floor common.Tile = '.'
const Empty common.Tile = 'L'
const Occupied common.Tile = '#'
const Null common.Tile = '💩'

var isOccupied = func(tile common.Tile) bool {
	return tile == Occupied
}

var isNotFloor = func(tile common.Tile) bool {
	return tile != Floor
}

func main() {
	original := common.ReadFileAsMap("11/input.txt")
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

				occupied := common.CountTiles(state.Neighbours(y, x, Floor), isOccupied)
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
			println(state.Count(isOccupied))
			break
		}
	}

	state = original
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

				occupied := common.CountTiles(state.Starburst(y, x, Null, isNotFloor), isOccupied)
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
			println(state.Count(isOccupied))
			break
		}
	}

}
