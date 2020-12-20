package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/scags9876/adventOfCode/lib"
	"regexp"
	"strings"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const testInput2Filename = "testinput2.txt"

const verbose = false

func main() {
	input := lib.GetInputStrings(inputFilename)
	part1(input)
	part2(input)
}

func part1(input []string) {
	rule, msgs := parseInput(input, map[int]string{}, 10)

	if verbose {
		fmt.Printf("input: %s", spew.Sdump(rule,msgs))
	}

	count := 0
	for _, msg := range msgs {
		if matchesRule(msg, rule) {
			count++
		}
	}
	fmt.Printf("Part 1: %d\n", count)
}

func part2(input []string) {
	replacements := map[int]string{
		8: "42 | 42 8",
		11: "42 31 | 42 11 31",
	}

	lastCount := 0
	// how deep to recurse before giving up, increasing this num until the given input gives a stable output
	maxDepth := 1
	for {
		rule, msgs := parseInput(input, replacements, maxDepth)

		if verbose {
			fmt.Printf("input: %s", spew.Sdump(rule,msgs))
		}

		count := 0
		for _, msg := range msgs {
			if matchesRule(msg, rule) {
				count++
			}
		}
		fmt.Printf("With maxDepth %d, got count %d\n", maxDepth, count)
		if lastCount == count {
			break
		}
		maxDepth++
		lastCount = count
	}

	fmt.Printf("Part 2: %d\n", lastCount)
}

func parseInput(input []string, replacements map[int]string, maxDepth int) (*regexp.Regexp, []string) {
	rules := make(map[int]string)
	msgs := make([]string, 0)
	mode := "rules"
	for _, line := range input {
		if len(line) == 0 {
			mode = "msgs"
		}
		switch mode {
		case "rules":
			parts := strings.Split(line, ":")
			id := lib.ToInt(parts[0])
			rule := strings.Trim(parts[1], ` ""`)
			if replaceRule, ok := replacements[id]; ok {
				rules[id] = replaceRule
			} else {
				rules[id] = rule
			}
		case "msgs":
			msgs = append(msgs, line)
		default:
			panic("unknown mode")
		}
	}

	if verbose {
		fmt.Printf("rules: %s", spew.Sdump(rules))
	}

	var reg strings.Builder
	reg.WriteRune('^')
	reg.WriteString(resolveRule(0, rules, 0, maxDepth))
	reg.WriteRune('$')

	if verbose {
		fmt.Printf("regex: %s\n", reg.String())
	}

	rule := regexp.MustCompile(reg.String())

	return rule, msgs
}


func resolveRule(ruleID int, rules map[int]string, loopDepth, maxDepth int) string {
	rule := rules[ruleID]
	var reg strings.Builder

	if strings.Contains(rule, "|") {
		reg.WriteString("((")
	}

	parts := strings.Fields(rule)
	for _, p := range parts {
		if isALetter(p) {
			reg.WriteString(p)
		} else if n, ok := lib.ToIntOk(p); ok {
			offset := 0
			if n == ruleID {
				if loopDepth > maxDepth {
					continue
				}
				offset = 1
			}
			ruleRegex := resolveRule(n, rules, loopDepth+offset, maxDepth)
			reg.WriteString(ruleRegex)
		} else if p == "|" {
			reg.WriteString(")|(")
		}
	}
	if strings.Contains(rule, "|") {
		reg.WriteString(	"))")
	}
	return reg.String()
}

func matchesRule(msg string, rule *regexp.Regexp) bool {
	return rule.MatchString(msg)
}

var letterRegExp = regexp.MustCompile(`^[ab]$`)
func isALetter(expr string) bool {
	match := letterRegExp.FindString(expr)
	if match == "" {
		return false
	}
	return true
}
