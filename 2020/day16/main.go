package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const testInputFilename2 = "testinput2.txt"

type rule struct {
	field             string
	lowRange          []int
	highRange         []int
	possiblePositions []int
	position          *int
}

func main() {
	input := lib.GetInputStrings(inputFilename)
	rules, myTicket, nearbyTickets := processInput(input)
	validTickets := part1(rules, myTicket, nearbyTickets)
	part2(rules, myTicket, validTickets)
}

const (
	rulesMode         = 1
	myTicketMode      = 2
	nearbyTicketsMode = 3
)

func processInput(input []string) ([]rule, []int, [][]int) {
	rules := make([]rule, 0)
	myTicket := make([]int, 0)
	nearbyTickets := make([][]int, 0)
	mode := rulesMode
	ruleRex := regexp.MustCompile(`^([^:]+): (\d+)-(\d+) or (\d+)-(\d+)$`)

	for _, line := range input {
		if len(line) == 0 {
			mode++
			continue
		}
		switch mode {
		case rulesMode:
			matches := ruleRex.FindStringSubmatch(line)
			r := rule{
				field:     matches[1],
				lowRange:  []int{lib.ToInt(matches[2]), lib.ToInt(matches[3])},
				highRange: []int{lib.ToInt(matches[4]), lib.ToInt(matches[5])},
			}
			rules = append(rules, r)

		case myTicketMode:
			if line == "your ticket:" {
				continue
			}
			for _, s := range strings.Split(line, ",") {
				myTicket = append(myTicket, lib.ToInt(s))
			}

		case nearbyTicketsMode:
			if line == "nearby tickets:" {
				continue
			}
			ticket := make([]int, 0)
			for _, s := range strings.Split(line, ",") {
				ticket = append(ticket, lib.ToInt(s))
			}
			nearbyTickets = append(nearbyTickets, ticket)

		}

	}
	return rules, myTicket, nearbyTickets
}

func part1(rules []rule, myTicket []int, nearbyTickets [][]int) [][]int {
	validValues := map[int]bool{}
	for _, rule := range rules {
		for i := rule.lowRange[0]; i <= rule.lowRange[1]; i++ {
			validValues[i] = true
		}
		for i := rule.highRange[0]; i <= rule.highRange[1]; i++ {
			validValues[i] = true
		}
	}
	validSet := make([]int, 0)
	for value, _ := range validValues {
		validSet = lib.SortedInsertInt(validSet, value)
	}

	invalidValues := make([]int, 0)
	for _, value := range myTicket {
		if !intInSet(validSet, value) {
			invalidValues = append(invalidValues, value)
		}
	}
	fmt.Printf("found %d invalid values in my ticket: %#v\n", len(invalidValues), invalidValues)

	validTickets := make([][]int, 0)
	for _, ticket := range nearbyTickets {
		ticketIsValid := true
		for _, value := range ticket {
			if !intInSet(validSet, value) {
				ticketIsValid = false
				invalidValues = append(invalidValues, value)
			}
		}
		if ticketIsValid {
			validTickets = append(validTickets, ticket)
		}
	}
	fmt.Printf("found %d invalid values in nearby tickets: %#v\n", len(invalidValues), invalidValues)

	result := 0
	for _, n := range invalidValues {
		result += n
	}
	fmt.Printf("Part 1: ticket scanning error rate: %d\n", result)

	return validTickets
}

func intInSet(set []int, n int) bool {
	i := sort.SearchInts(set, n)
	return i < len(set) && set[i] == n
}

func part2(rules []rule, myTicket []int, validTickets [][]int) {
	// fill out every rule to show that it can be in any position to start
	rules = initRulePossibilities(rules, len(myTicket))

	for _, ticket := range validTickets {
		for position, value := range ticket {
			// find out which rule this value can't be a part of
			for i, rule := range rules {
				if !intInSet(rule.possiblePositions, position) {
					continue
				}
				if value < rule.lowRange[0] || value > rule.highRange[1] || (value > rule.lowRange[1] && value < rule.highRange[0]) {
					newPossiblePositions := removeIntFromSet(rule.possiblePositions, position)
					rule.possiblePositions = newPossiblePositions
					rules[i] = rule
				}
			}
		}
	}

	// iterate until every rule has a fixed position
	done := false
	iterationCount := 0
	for !done {
		iterationCount++
		// go through each position
		for i := 0; i < len(myTicket); i++ {
			rulesThatCanFitThisPosition := make([]rule, 0)
			for _, rule := range rules {
				if intInSet(rule.possiblePositions, i) {
					rulesThatCanFitThisPosition = append(rulesThatCanFitThisPosition, rule)
				}
			}
			// if only 1 rule can fit this postion, assign that rule to this position
			if len(rulesThatCanFitThisPosition) == 1 {
				position := i
				newRule := rulesThatCanFitThisPosition[0]
				newRule.position = &position
				// when assigning a rule to a position, make sure to remove all other possible positions this rule can apply to
				// this is how the list gets whittled down
				newRule.possiblePositions = []int{i}
				for i, rule := range rules {
					if rule.field == newRule.field {
						rules[i] = newRule
					}
				}
			}
			//fmt.Printf("position %d is in rules %#v\n", i, rulesThatCanFitThisPosition)
		}
		done = true
		for _, rule := range rules {
			if rule.position == nil {
				done = false
				break
			}
		}
	}

	fmt.Printf("After %d iterations, my Ticket:\n", iterationCount)

	result := 1
	for _, rule := range rules {
		if rule.position != nil {
			fmt.Printf("%25s: %d\n", rule.field, myTicket[*rule.position])
			if strings.HasPrefix(rule.field, "departure") {
				result *= myTicket[*rule.position]
			}
			continue
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}

func initRulePossibilities(rules []rule, numFields int) []rule {
	newRules := make([]rule, 0)
	for _, rule := range rules {
		newRule := rule
		possibleFields := make([]int, numFields)
		for i := 0; i < numFields; i++ {
			possibleFields[i] = i
		}
		newRule.possiblePositions = possibleFields
		newRules = append(newRules, newRule)
	}
	return newRules
}

func removeIntFromSet(set []int, n int) []int {
	i := sort.SearchInts(set, n)
	if i == len(set) {
		return set
	}
	return append(set[:i], set[i+1:]...)
}
