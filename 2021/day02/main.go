package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 150
const expectedTestResultPart2 = 900

const verbose = true

func main() {
	testInput := lib.GetInputStrings(testInputFilename)
	input := lib.GetInputStrings(inputFilename)

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

func part1(input []string) int {
	var x, depth int
	for _, line := range input {
		fields := strings.Split(line, " ")
		command := fields[0]
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			panic("could not convert to number")
		}
		switch command {
		case "forward":
			x = x + n
		case "up":
			depth = depth - n
		case "down":
			depth = depth + n
		}
		if verbose {
			fmt.Printf("After %s, position: x=%d, depth=%d\n", line, x, depth)
		}
	}
	if verbose {
		fmt.Printf("Final position: x=%d, depth=%d\n", x, depth)
	}
	return x * depth
}

func part2(input []string) int {
	var x, depth, aim int
	for _, line := range input {
		fields := strings.Split(line, " ")
		command := fields[0]
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			panic("could not convert to number")
		}
		switch command {
		case "forward":
			x = x + n
			depth = depth + (aim * n)
		case "up":
			aim = aim - n
		case "down":
			aim = aim + n
		}
		if verbose {
			fmt.Printf("After %s, position: x=%d, depth=%d, aim=%d\n", line, x, depth, aim)
		}
	}
	if verbose {
		fmt.Printf("Final position: x=%d, depth=%d, aim=%d\n", x, depth, aim)
	}
	return x * depth
}
