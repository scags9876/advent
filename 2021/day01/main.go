package main

import (
	"fmt"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 7
const expectedTestResultPart2 = 5

const verbose = false

func main() {
	testInput := lib.GetInputInts(testInputFilename)
	input := lib.GetInputInts(inputFilename)

	fmt.Print("Part 1:\n")
	testResult := part1(testInput)
	if testResult == expectedTestResultPart1 {
		fmt.Printf("testfile returned expected result! (%d)\n", testResult)
	} else {
		panic(fmt.Errorf("expected %d but got %d", expectedTestResultPart1, testResult))
	}

	result := part1(input)
	fmt.Printf("result: %d\n", result)

	fmt.Print("\nPart 2: \n")
	testResult = part2(testInput)
	if testResult == expectedTestResultPart2 {
		fmt.Printf("testfile returned expected result! (%d)\n", testResult)
	} else {
		panic(fmt.Errorf("expected %d but got %d", expectedTestResultPart2, testResult))
	}

	result = part2(input)
	fmt.Printf("Part 2 result: %d\n", result)
}

func part1(input []int) int {
	var currentDepth int
	var increases int
	for _, depth := range input {
		if currentDepth != 0 {
			if depth > currentDepth {
				increases++
			}
		}
		currentDepth = depth
	}
	return increases
}

func part2(input []int) int {
	var increases int
	for i := range input {
		if i < 3 {
			continue
		}
		sum1 := input[i-1] + input[i-2] + input[i-3]
		sum2 := input[i]   + input[i-1] + input[i-2]
		if sum2 > sum1 {
			increases++
		}
	}
	return increases
}
