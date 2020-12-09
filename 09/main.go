package main

import (
	"github.com/csmith/aoc-2020/common"
	"math"
)

func main() {
	lines := common.ReadFileAsInts("09/input.txt")

	target := 0
	for i := 25; i < len(lines); i++ {
		var found = false
		for j := i - 25; j < i - 1; j++ {
			for k := j + 1; k < i; k++ {
				if lines[j] + lines[k] == lines[i] && lines[j] != lines[k] {
					found = true
				}
			}
		}
		if !found {
			target = lines[i]
			break
		}
	}

	println(target)

	start := 0
	count := 0
	for i := range lines {
		count += lines[i]
		for count > target {
			count -= lines[start]
			start++
		}
		if count == target && (i - start) >= 2 {
			min := math.MaxInt64
			max := 0
			for j := start; j <= i; j++ {
				min = common.Min(min, lines[j])
				max = common.Max(max, lines[j])
			}
			println(min + max)
			break
		}
	}
}
