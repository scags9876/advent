package main

import (
	"fmt"
	"sort"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 24000
const expectedTestResultPart2 = 45000

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
	elves := make([][]int, 1)
	elfIdx := 0
	elfTotal := 0
	var result int
	for _, val := range input {
		if len(val) > 0 {
			cal := lib.ToInt(val)
			elfTotal += cal
			elves[elfIdx] = append(elves[elfIdx], cal)
		} else {
			if verbose {
				fmt.Printf("Elf %d finished! Total Calories: %d\n", elfIdx+1, elfTotal)
			}
			if result == 0 || result < elfTotal {
				result = elfTotal
			}
			elfTotal = 0
			elfIdx++
			elves = append(elves, make([]int, 1))
		}
	}
	return result
}

func part2(input []string) int {
	elves := make([]int, 1)
	elfIdx := 0
	for _, val := range input {
		if len(val) > 0 {
			cal := lib.ToInt(val)
			elves[elfIdx] += cal
		} else {
			if verbose {
				fmt.Printf("Elf %d finished! Total Calories: %d\n", elfIdx+1, elves[elfIdx])
			}
			elfIdx++
			elves = append(elves, 0)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	result := lib.SumInts(elves[:3])
	return result
}
