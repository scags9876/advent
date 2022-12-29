package main

import (
	"fmt"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"

var part1Tests = map[string]int{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
	"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
	"nppdvjthqldpwncqszvftbrmjlhg":      6,
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
}

var part2Tests = map[string]int{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
	"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
	"nppdvjthqldpwncqszvftbrmjlhg":      23,
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
}

const verbose = true

func main() {
	input := lib.GetInputStrings(inputFilename)

	fmt.Print("Part 1:\n")
	for testInput, expectedResult := range part1Tests {
		testResult := part1(testInput)
		if verbose {
			fmt.Printf("Part 1 testInput (%s) result: %d\n", testInput, testResult)
		}
		if testResult == expectedResult {
			fmt.Printf("test returned expected result! (%d)\n", testResult)
		} else {
			panic(fmt.Errorf("expected %d but got %d", expectedResult, testResult))
		}
	}

	result := part1(input[0])
	fmt.Printf("Part 1 result: %d\n", result)

	fmt.Print("Part 2:\n")
	for testInput, expectedResult := range part2Tests {
		testResult := part2(testInput)
		if verbose {
			fmt.Printf("Part 2 testInput (%s) result: %d\n", testInput, testResult)
		}
		if testResult == expectedResult {
			fmt.Printf("test returned expected result! (%d)\n", testResult)
		} else {
			panic(fmt.Errorf("expected %d but got %d", expectedResult, testResult))
		}
	}

	result = part2(input[0])
	fmt.Printf("Part 2 result: %d\n", result)
}

func part1(input string) int {
	return findMarker(input, 4)
}

func part2(input string) int {
	return findMarker(input, 14)
}

func findMarker(input string, markerSize int) int {
	var result int
	inputLen := len(input)
	markerIdx := markerSize - 1
	for i := markerIdx; i < inputLen; i++ {
		if nonRepeatingChars(input[i-markerIdx : i+1]) {
			if verbose {
				fmt.Printf("non repeating sequence found at position %d: %s\n", i+1, input[i-markerIdx:i+1])
			}
			result = i + 1
			return result
		}
	}
	return result
}

func nonRepeatingChars(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
