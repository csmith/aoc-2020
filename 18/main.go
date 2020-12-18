package main

import (
	"github.com/csmith/aoc-2020/common"
	"regexp"
	"strconv"
	"strings"
)

var bracketRegexp = regexp.MustCompile(`\([^()]+\)`)
var sumRegexp = regexp.MustCompile(`(\d+) \+ (\d+)`)

func evalLeftToRight(equation string) int {
	for strings.Contains(equation, "(") {
		equation = bracketRegexp.ReplaceAllStringFunc(equation, func(s string) string {
			return strconv.Itoa(evalLeftToRight(s[1 : len(s)-1]))
		})
	}

	parts := strings.Split(equation, " ")
	ans := common.MustAtoi(parts[0])
	for i := 1; i < len(parts)-1; i += 2 {
		switch parts[i] {
		case "+":
			ans += common.MustAtoi(parts[i+1])
		case "*":
			ans *= common.MustAtoi(parts[i+1])
		}
	}
	return ans
}

func evalAdditionThenMultiplication(equation string) int {
	for strings.Contains(equation, "(") {
		equation = bracketRegexp.ReplaceAllStringFunc(equation, func(s string) string {
			return strconv.Itoa(evalAdditionThenMultiplication(s[1 : len(s)-1]))
		})
	}

	for strings.Contains(equation, "+") {
		equation = sumRegexp.ReplaceAllStringFunc(equation, func(s string) string {
			parts := strings.Split(s, " ")
			return strconv.Itoa(common.MustAtoi(parts[0]) + common.MustAtoi(parts[2]))
		})
	}

	parts := strings.Split(equation, " ")
	ans := common.MustAtoi(parts[0])
	for i := 1; i < len(parts)-1; i += 2 {
		ans *= common.MustAtoi(parts[i+1])
	}
	return ans
}

func main() {
	lines := common.ReadFileAsStrings("18/input.txt")
	sum1 := 0
	sum2 := 0
	for i := range lines {
		sum1 += evalLeftToRight(lines[i])
		sum2 += evalAdditionThenMultiplication(lines[i])
	}
	println(sum1)
	println(sum2)
}
