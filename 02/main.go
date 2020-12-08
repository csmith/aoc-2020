package main

import (
	"github.com/csmith/aoc-2020/common"
	"regexp"
	"strings"
)

func main() {
	part1 := 0
	part2 := 0

	lines := common.TokeniseLines(
		common.ReadFileAsStrings("02/input.txt"),
		regexp.MustCompile(`^(\d+)-(\d+) (.*?): (.*?)$`),
	)
	for i := range lines {
		var (
			min = common.MustAtoi(lines[i][0])
			max = common.MustAtoi(lines[i][1])
			char = lines[i][2]
			password = lines[i][3]
		)

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
