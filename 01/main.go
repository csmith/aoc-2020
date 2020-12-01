package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
)

func main() {
	const target = 2020
	entries := common.ReadFileAsInts("01/input.txt")

	var part1, part2 int
	for i := 0; i < len(entries); i++ {
		for ip := i + 1; ip  < len(entries); ip++ {
			if entries[i] + entries[ip] == target {
				part1 = entries[i] * entries[ip]
			}

			for ipp := ip+1; ipp < len(entries); ipp++ {
				if entries[i] + entries[ip] + entries[ipp] == target {
					part2 = entries[i] * entries[ip] * entries[ipp]
				}
			}
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
