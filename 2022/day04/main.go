package main

import (
	"fmt"
	"strings"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 2
const expectedTestResultPart2 = 4

const verbose = true

func main() {
	testInput := lib.GetInputStrings(testInputFilename)
	input := lib.GetInputStrings(inputFilename)

	fmt.Print("Part 1:\n")
	testResult := part1(testInput)
	if verbose {
		fmt.Printf("Part 1 testInput result: %d\n", testResult)
	}
	if testResult == expectedTestResultPart1 {
		fmt.Printf("testfile returned expected result! (%d)\n", testResult)
	} else {
		panic(fmt.Errorf("expected %d but got %d", expectedTestResultPart1, testResult))
	}

	result := part1(input)
	fmt.Printf("Part 1 result: %d\n", result)

	fmt.Print("\nPart 2: \n")
	testResult = part2(testInput)
	if verbose {
		fmt.Printf("Part 2 testInput result: %d\n", testResult)
	}
	if testResult == expectedTestResultPart2 {
		fmt.Printf("testfile returned expected result! (%d)\n", testResult)
	} else {
		panic(fmt.Errorf("expected %d but got %d", expectedTestResultPart2, testResult))
	}

	result = part2(input)
	fmt.Printf("Part 2 result: %d\n", result)
}

func part1(input []string) int {
	var result int
	for i, line := range input {
		ranges := strings.Split(line, ",")
		range1 := lib.StringToInts(ranges[0], "-")
		range2 := lib.StringToInts(ranges[1], "-")
		if isSubset(range1, range2) {
			if verbose {
				fmt.Printf("line %2d: range 1 (%d-%d) is a subset of range 2 (%d-%d)\n", i, range1[0], range1[1], range2[0], range2[1])
			}
			result++
		} else if isSubset(range2, range1) {
			if verbose {
				fmt.Printf("line %2d: range 2 (%d-%d) is a subset of range 1 (%d-%d)\n", i, range2[0], range2[1], range1[0], range1[1])
			}
			result++
		}
	}
	return result
}

func part2(input []string) int {
	var result int
	for i, line := range input {
		ranges := strings.Split(line, ",")
		range1 := lib.StringToInts(ranges[0], "-")
		range2 := lib.StringToInts(ranges[1], "-")
		if anyOverlap(range1, range2) {
			if verbose {
				fmt.Printf("line %2d: range 1 (%d-%d) overlaps with range 2 (%d-%d)\n", i, range1[0], range1[1], range2[0], range2[1])
			}
			result++
		}
	}
	return result
}

func isSubset(a, b []int) bool {
	return a[0] <= b[0] && a[1] >= b[1]
}

func anyOverlap(a, b []int) bool {
	return (a[0] <= b[0] && a[1] >= b[0]) ||
		(a[0] <= b[1] && a[1] >= b[1]) ||
		isSubset(a, b) ||
		isSubset(b, a)
}
