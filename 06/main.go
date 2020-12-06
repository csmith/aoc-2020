package main

import (
	"github.com/csmith/aoc-2020/common"
	"strings"
)

func main() {
	groups := common.ReadFileAsStringChunks("06/input.txt")
	distinct := 0
	shared := 0

	for i := range groups {
		var options [26]int
		for _, o := range groups[i] {
			if o >= 'a' && o <= 'z' {
				options[o-'a']++
			}
		}

		for _, count := range options {
			if count > 0 {
				distinct++
			}
			if count == strings.Count(groups[i], " ")+1 {
				shared++
			}
		}
	}

	println(distinct)
	println(shared)
}
