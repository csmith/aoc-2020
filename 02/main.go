package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"strings"
)

func main() {
	part1 := 0
	part2 := 0

	lines := common.ReadFileAsStrings("02/input.txt")
	for i := range lines {
		var (
			min, max int
			char string
			password string
		)
		if _, err := fmt.Sscanf(lines[i], "%d-%d %1s: %s", &min, &max, &char, &password); err != nil {
			panic(err)
		}

		if count := strings.Count(password, char); count >= min && count <= max {
			part1++
		}

		if (password[min-1] == char[0]) != (password[max-1] == char[0]) {
			part2++
		}
	}

	println(part1)
	println(part2)
}
