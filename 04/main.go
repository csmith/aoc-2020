package main

import (
	"github.com/csmith/aoc-2020/common"
	"regexp"
)

func parse(chunks [][]string) []map[string]string {
	var passports []map[string]string
	for i := range chunks {
		details := make(map[string]string)
		for j := 0; j < len(chunks[i]); j += 2 {
			details[chunks[i][j]] = chunks[i][j+1]
		}
		passports = append(passports, details)
	}
	return passports
}

func main() {
	passports := parse(common.TokeniseLines(
		common.ReadFileAsStringChunks("04/input.txt"),
		regexp.MustCompile(`(\w{3}):(\S+)`),
	))

	rules := map[string]*regexp.Regexp{
		"byr": regexp.MustCompile(`^(19[0-9]{2}|200[012])$`),
		"iyr": regexp.MustCompile(`^20(1\d|20)$`),
		"eyr": regexp.MustCompile(`^20(2\d|30)$`),
		"hgt": regexp.MustCompile(`^(1([5-8]\d|9[0-3])cm|(59|6\d|7[0-6])in)$`),
		"hcl": regexp.MustCompile(`^#[\da-f]{6}$`),
		"ecl": regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
		"pid": regexp.MustCompile(`^\d{9}$`),
	}

	presentCount := 0
	validatedCount := 0
	for i := range passports {
		present := true
		validated := true
		for field, regex := range rules {
			val, ok := passports[i][field]
			if !ok {
				present = false
				break
			}
			if !regex.MatchString(val) {
				validated = false
			}
		}
		if present {
			presentCount++
			if validated {
				validatedCount++
			}
		}
	}

	println(presentCount)
	println(validatedCount)
}
