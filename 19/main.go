package main

import (
	"fmt"
	"github.com/csmith/aoc-2020/common"
	"regexp"
	"strings"
)

var (
	rawRules    = make(map[int][]string)
	parsedRules = make(map[int]string)
)

func Matcher(rule int) string {
	if p, ok := parsedRules[rule]; ok {
		return p
	}

	if rawRules[rule][0][0] == '"' {
		value := string(rawRules[rule][0][1])
		parsedRules[rule] = value
		return value
	}

	builder := strings.Builder{}
	builder.WriteRune('(')
	for p := range rawRules[rule] {
		part := rawRules[rule][p]
		if part == "|" {
			builder.WriteRune('|')
		} else {
			builder.WriteString(Matcher(common.MustAtoi(part)))
		}
	}
	builder.WriteRune(')')
	value := builder.String()

	parsedRules[rule] = value
	return value
}

func main() {
	sections := common.ReadFileAsSectionedStrings("19/input.txt")

	// Parse the raw rules out of the input
	for r := range sections[0] {
		parts := strings.Split(sections[0][r], " ")
		number := common.MustAtoi(strings.TrimSuffix(parts[0], ":"))
		rawRules[number] = parts[1:]
	}

	// Build up a giant regex for rule 0
	matcher1 := regexp.MustCompile(fmt.Sprintf("^%s$", Matcher(0)))

	// Throw the cached values away and force some weird loops in
	parsedRules = make(map[int]string)
	parsedRules[8] = fmt.Sprintf("%s+", Matcher(42))

	rule11 := strings.Builder{}
	rule11.WriteRune('(')
	for i := 1; i < 10; i++ {
		// Go doesn't support recursive submatches, so just bodge manual matches for up to 10 repeats.
		// Experimentally more than 4-5 doesn't seem to affect the result.
		if rule11.Len() > 1 {
			rule11.WriteRune('|')
		}
		rule11.WriteString(Matcher(42))
		rule11.WriteString(fmt.Sprintf("{%d}", i))
		rule11.WriteString(Matcher(31))
		rule11.WriteString(fmt.Sprintf("{%d}", i))
	}
	rule11.WriteRune(')')
	parsedRules[11] = rule11.String()

	matcher2 := regexp.MustCompile(fmt.Sprintf("^%s$", Matcher(0)))

	// Check which messages match
	matching1, matching2 := 0, 0
	for i := range sections[1] {
		if matcher1.MatchString(sections[1][i]) {
			matching1++
		}
		if matcher2.MatchString(sections[1][i]) {
			matching2++
		}
	}
	println(matching1)
	println(matching2)
}
