package main

import (
	"github.com/csmith/aoc-2020/common"
	"strconv"
	"strings"
)

var (
	containedIn = make(map[string][]string)
	contents    = make(map[string][]string)
	sums        = make(map[string]int)
	found       = make(map[string]bool)
)

func check(bag string) int {
	count := 0
	for _, b := range containedIn[bag] {
		if !found[b] {
			found[b] = true
			count += 1 + check(b)
		}
	}
	return count
}

func sum(bag string) int {
	if cached, ok := sums[bag]; ok {
		return cached
	}

	res := 1
	for _, b := range contents[bag] {
		parts := strings.SplitN(b, " ", 2)
		count, _ := strconv.Atoi(parts[0])
		res += count * sum(parts[1])
	}
	sums[bag] = res
	return res
}

func main() {
	lines := common.ReadFileAsStrings("07/input.txt")

	for i := range lines {
		line := lines[i]
		parts := strings.SplitN(strings.TrimSuffix(line, "."), " contain ", 2)
		bag := strings.TrimRight(parts[0], "s")
		inner := strings.Split(parts[1], ", ")
		for _, c := range inner {
			name := strings.TrimRight(strings.TrimLeft(c, "0123456789 "), "s")
			containedIn[name] = append(containedIn[name], bag)
			if c != "no other bags" {
				contents[bag] = append(contents[bag], strings.TrimRight(c, "s"))
			}
		}
	}

	println(check("shiny gold bag"))
	println(sum("shiny gold bag") - 1)
}
