package main

import (
	"github.com/csmith/aoc-2020/common"
	"regexp"
	"sort"
	"strings"
)

type Rule struct {
	Name           string
	Ranges         [][2]int
	PossibleFields []bool
}

func (r Rule) Accepts(value int) bool {
	for i := range r.Ranges {
		if value >= r.Ranges[i][0] && value <= r.Ranges[i][1] {
			return true
		}
	}
	return false
}

func (r Rule) PossibleFieldCount() int {
	count := 0
	for i := range r.PossibleFields {
		if r.PossibleFields[i] {
			count++
		}
	}
	return count
}

var ruleRegex = regexp.MustCompile(`^(.*?): (\d+)-(\d+) or (\d+)-(\d+)$`)

func main() {
	var (
		rules         []Rule
		nearbyTickets [][]int
	)

	sections := common.ReadFileAsSectionedStrings("16/input.txt")

	// First section - rules
	for l := range sections[0] {
		tokens := common.TokeniseLine(sections[0][l], ruleRegex)
		rules = append(rules, Rule{
			Name: tokens[0],
			Ranges: [][2]int{
				{common.MustAtoi(tokens[1]), common.MustAtoi(tokens[2])},
				{common.MustAtoi(tokens[3]), common.MustAtoi(tokens[4])},
			},
		})
	}

	// Second section - my ticket. First line is a header.
	myTicket := ParseTicket(sections[1][1])

	// Third section - nearby tickets. First line is a header.
	for l := range sections[2] {
		if l > 0 {
			nearbyTickets = append(nearbyTickets, ParseTicket(sections[2][l]))
		}
	}

	// Initialise the possible fields slice to the right length
	for r := range rules {
		rules[r].PossibleFields = make([]bool, len(myTicket))
		for i := range rules[r].PossibleFields {
			rules[r].PossibleFields[i] = true
		}
	}

	errorRate := 0
	for t := range nearbyTickets {
		// Loop through all the fields and make sure they're permitted by _some_ rule. The ones that aren't are
		// summed to get the "error rate".
		invalid := false
		for f := range nearbyTickets[t] {
			found := false
			for r := range rules {
				if rules[r].Accepts(nearbyTickets[t][f]) {
					found = true
					break
				}
			}
			if !found {
				invalid = true
				errorRate += nearbyTickets[t][f]
			}
		}

		// If the ticket doesn't contain any completely invalid fields, use it to ratchet down the set of possible
		// field orderings.
		if !invalid {
			for f := range nearbyTickets[t] {
				for r := range rules {
					if !rules[r].Accepts(nearbyTickets[t][f]) {
						rules[r].PossibleFields[f] = false
					}
				}
			}
		}
	}

	// At this point we should have a rule with one possible field, another with two possible fields, and so on and
	// so forth. If we sort them by possible field count we should be able to do a single pass and assign them all.
	sort.SliceStable(rules, func(i, j int) bool {
		return rules[i].PossibleFieldCount() < rules[j].PossibleFieldCount()
	})

	pinned := make([]bool, len(myTicket))
	part2 := 1
	for r := range rules {
		for p := range pinned {
			if rules[r].PossibleFields[p] && !pinned[p] {
				pinned[p] = true
				if strings.HasPrefix(rules[r].Name, "departure") {
					part2 *= myTicket[p]
				}
				break
			}
		}
	}

	println(errorRate)
	println(part2)
}

func ParseTicket(line string) []int {
	var res []int
	fields := strings.Split(line, ",")
	for i := range fields {
		res = append(res, common.MustAtoi(fields[i]))
	}
	return res
}
