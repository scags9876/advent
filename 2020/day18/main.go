package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"

const verbose = false

func main() {
	input := lib.GetInputStrings(inputFilename)
	part1(input)
	part2(input)
}

func part1(input []string) {
	result := int64(0)
	for _, line := range input {
		r := calculate(strings.Join(strings.Fields(line), ""))
		fmt.Printf("%s = %d new total: %d\n", line, r, result+r)
		result += r
	}
	fmt.Printf("Part 1: %d\n", result)
}

func part2(input []string) {
	result := int64(0)
	for _, line := range input {
		r := calculatev2(strings.Join(strings.Fields(line), ""))
		fmt.Printf("%s = %d new total: %d\n", line, r, result+r)
		result += r
	}
	fmt.Printf("Part 2: %d\n", result)
}

const (
	add  = '+'
	mult = '*'
)

func calculate(expr string) int64 {
	position := 0
	total := int64(0)
	op := add
	for {
		if position == len(expr) {
			break
		}
		token := rune(expr[position])

		if n, ok := isANumber(expr, position); ok {
			if op == add {
				total += n
			} else {
				total *= n
			}
			position += len(fmt.Sprintf("%d", n))
		} else if token == '(' {
			subExpr := subExprStartingAt(expr, position)
			n := calculate(subExpr)
			if op == add {
				total += n
			} else {
				total *= n
			}
			position += len(subExpr) + 2
		} else if token == add || token == mult {
			if token == add {
				op = add
			} else if token == mult {
				op = mult
			}
			position++
		} else {
			panic("unexpected token")
		}
	}
	return total
}

func calculatev2(expr string) int64 {
	if verbose {
		fmt.Printf("%s =>\n", expr)
	}
	newExpr := calcSubexpr(expr)
	if verbose {
		fmt.Printf("subExpr(%s) => %s\n", expr, newExpr)
	}
	newExpr = calcAdditions(newExpr)
	if verbose {
		fmt.Printf("additions(%s) => %s\n", expr, newExpr)
	}
	newExpr = calcMultiplications(newExpr)
	if verbose {
		fmt.Printf("multiplications(%s) => %s\n", expr, newExpr)
	}
	total, err := strconv.ParseInt(newExpr, 10, 64)
	if err != nil {
		panic(err)
	}
	if verbose {
		fmt.Printf("%s => %d\n", expr, total)
	}
	return total
}

func calcSubexpr(expr string) string {
	var newExpr strings.Builder

	position := 0
	for {
		if position == len(expr) {
			break
		}
		token := expr[position]

		if token == '(' {
			subExpr := subExprStartingAt(expr, position)
			n := calculatev2(subExpr)
			if verbose {
				fmt.Printf("subExpr %s = %d\n", subExpr, n)
			}
			newExpr.WriteString(fmt.Sprintf("%d", n))
			position += len(subExpr) + 2
		} else {
			newExpr.WriteByte(token)
			position++
		}
	}
	return newExpr.String()
}

func calcAdditions(expr string) string {
	newExpr := expr

	addRegExp := regexp.MustCompile(`\d+\+\d+`)
	for {
		matches := addRegExp.FindAllIndex([]byte(newExpr), -1)
		if len(matches) == 0 {
			break
		}
		matchLoc := matches[0]
		match := newExpr[matchLoc[0]:matchLoc[1]]
		var a, b int64
		n, err := fmt.Sscanf(match, "%d+%d", &a, &b)
		if err != nil || n != 2 {
			panic(err)
		}
		//fmt.Printf("newExpr %s found %s, %d+%d=%d => ", newExpr, match, a,b,a+b)

		newExpr = fmt.Sprintf("%s%d%s", newExpr[:matchLoc[0]], a+b, newExpr[matchLoc[1]:])
	}
	return newExpr
}

func calcMultiplications(expr string) string {
	newExpr := expr

	addRegExp := regexp.MustCompile(`\d+\*\d+`)
	for {
		match := addRegExp.FindString(newExpr)
		if match == "" {
			break
		}
		var a, b int64
		n, err := fmt.Sscanf(match, "%d*%d", &a, &b)
		if err != nil || n != 2 {
			panic(err)
		}
		newExpr = strings.Replace(newExpr, match, fmt.Sprintf("%d", a*b), -1)
	}
	return newExpr
}

var numRegExp = regexp.MustCompile(`^\d+`)

func isANumber(expr string, position int) (int64, bool) {
	match := numRegExp.FindString(expr[position:])
	if match == "" {
		return 0, false
	}
	n, err := strconv.ParseInt(match, 10, 64)
	if err != nil {
		return 0, false
	}
	return n, true
}

func subExprStartingAt(expr string, position int) string {
	var subExpr strings.Builder
	openParenCount := 1
	for p := position + 1; p < len(expr); p++ {
		switch expr[p] {
		case '(':
			openParenCount++
		case ')':
			openParenCount--
		}
		if openParenCount == 0 {
			break
		}
		subExpr.WriteByte(expr[p])
	}
	if openParenCount != 0 {
		panic("unclosed paren!")
	}
	return subExpr.String()
}
