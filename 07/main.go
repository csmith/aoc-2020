package main

import (
	"github.com/csmith/aoc-2020/common"
	"regexp"
)

type Quantity struct {
	name  string
	count int
}

var (
	containedIn = make(map[string][]string)
	contents    = make(map[string][]Quantity)
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
		res += b.count * sum(b.name)
	}
	sums[bag] = res
	return res
}

func main() {
	tokens := common.TokeniseLines(
		common.ReadFileAsStrings("07/input.txt"),
		regexp.MustCompile(`(^.*?) bags?|(\d+) (.*?) bags?`),
	)

	for i := range tokens {
		line := tokens[i]
		bag := line[0]
		for i := 1; i < len(line); i += 2 {
			name := line[i+1]
			containedIn[name] = append(containedIn[name], bag)
			contents[bag] = append(contents[bag], Quantity{
				name:  name,
				count: common.MustAtoi(line[i]),
			})
		}
	}

	println(check("shiny gold"))
	println(sum("shiny gold") - 1)
}
