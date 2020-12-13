package main

import (
	"github.com/csmith/aoc-2020/common"
	"log"
	"regexp"
)

func main() {
	instr := common.TokeniseLines(
		common.ReadFileAsStrings("12/input.txt"),
		regexp.MustCompile(`(.)(\d+)`),
	)

	println(part1(instr))
	println(part2(instr))
}

func part1(instr [][]string) int64 {
	bearing := 90
	north := 0
	east := 0
	for i := range instr {
		value := common.MustAtoi(instr[i][1])
		switch instr[i][0] {
		case "N":
			north += value
		case "S":
			north -= value
		case "E":
			east += value
		case "W":
			east -= value
		case "L":
			bearing = ((360 + bearing) - value) % 360
		case "R":
			bearing = (bearing + value) % 360
		case "F":
			switch bearing {
			case 0:
				north += value
			case 90:
				east += value
			case 180:
				north -= value
			case 270:
				east -= value
			default:
				log.Panicf("Unsupported bearing: %d", bearing)
			}
		default:
			log.Panicf("Unsupported instruction %s at line %d", instr[i][0], i)
		}
	}
	return common.Abs(int64(north)) + common.Abs(int64(east))
}

func part2(instr [][]string) int64 {
	waypointNorth := 1
	waypointEast := 10
	shipNorth := 0
	shipEast := 0
	for i := range instr {
		value := common.MustAtoi(instr[i][1])
		switch instr[i][0] {
		case "N":
			waypointNorth += value
		case "S":
			waypointNorth -= value
		case "E":
			waypointEast += value
		case "W":
			waypointEast -= value
		case "L":
			value = 360 - value
			fallthrough
		case "R":
			for value > 0 {
				value -= 90
				waypointEast, waypointNorth = waypointNorth, -waypointEast
			}
		case "F":
			shipNorth += value * waypointNorth
			shipEast += value * waypointEast
		default:
			log.Panicf("Unsupported instruction %s at line %d", instr[i][0], i)
		}
	}
	return common.Abs(int64(shipNorth)) + common.Abs(int64(shipEast))
}
