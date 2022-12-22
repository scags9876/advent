package main

import (
	"fmt"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 157
const expectedTestResultPart2 = 70

const verbose = true

var priorities map[rune]int

func main() {
	testInput := lib.GetInputStrings(testInputFilename)
	input := lib.GetInputStrings(inputFilename)

	setPriorities()

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
		lineLength := len(line)
		ruckSize := lineLength / 2
		if verbose {
			fmt.Printf("line %2d, length: %d, rucksack size: %d\n", i, lineLength, ruckSize)
		}
		ruck1 := line[:ruckSize]
		ruck2 := line[ruckSize:]
		sharedItems := findSharedItems(ruck1, ruck2)
		if verbose {
			fmt.Printf("line %d2: number of shared items: %d\n", i, len(sharedItems))
		}
		for _, c := range sharedItems {
			result += priorities[c]
		}
	}
	return result
}

func part2(input []string) int {
	var result int
	line := 0
	groupNum := 1
	inputLength := len(input)
	if verbose {
		fmt.Printf("input has length %d, meaning there are %d groups\n", inputLength, inputLength/3)
	}
	for line < inputLength {
		group := input[line : line+3]
		commonItems := findSharedItems(group[0], group[1])
		if verbose {
			fmt.Printf("group %2d: %d common items in first 2 lines\n", groupNum, len(commonItems))
		}
		commonItems = findSharedItems(string(commonItems), group[2])
		if verbose {
			fmt.Printf("group %2d: %d common items left: %c\n", groupNum, len(commonItems), commonItems[0])
		}

		result += priorities[commonItems[0]]

		line += 3
		groupNum++
	}
	return result
}

func setPriorities() {
	priorities = make(map[rune]int, 52)

	val := 1
	for i := 'a'; i <= 'z'; i++ {
		if verbose {
			fmt.Printf("\tpriorities[%c]: %d\n", i, val)
		}

		priorities[i] = val
		val++
	}
	for i := 'A'; i <= 'Z'; i++ {
		if verbose {
			fmt.Printf("\tpriorities[%c]: %d\n", i, val)
		}
		priorities[i] = val
		val++
	}
	return
}

func findSharedItems(set1, set2 string) []rune {
	commonItemsMap := make(map[rune]int)
	for _, c1 := range set1 {
		for _, c2 := range set2 {
			if c1 == c2 {
				commonItemsMap[c1] = 1
			}
		}
	}
	commonItems := make([]rune, 0)
	for c := range commonItemsMap {
		commonItems = append(commonItems, c)
	}
	return commonItems
}
