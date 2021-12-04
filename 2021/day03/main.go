package main

import (
	"fmt"
	"github.com/scags9876/adventOfCode/lib"
	"strconv"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 198
const expectedTestResultPart2 = 230

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
	bits := len(input[0])
	bitCounts := make([]map[rune]int, bits)
	for i := 0; i < bits; i++ {
		bitCounts[i] = make(map[rune]int)
	}
	for _, line := range input {
		for i, c := range line {
			bitCounts[i][c]++
		}
	}
	if verbose {
		fmt.Printf("After reading input, bitCounts is: %v\n", bitCounts)
	}

	var gammaStr, epsilonStr string
	for i := 0; i < bits; i++ {
		if bitCounts[i]['0'] > bitCounts[i]['1'] {
			gammaStr = gammaStr + "0"
			epsilonStr = epsilonStr + "1"
		} else {
			gammaStr = gammaStr + "1"
			epsilonStr = epsilonStr + "0"
		}
	}
	gamma, err := strconv.ParseInt(gammaStr, 2, 64)
	if err != nil {
		panic("can't parse binary")
	}
	epsilon, err := strconv.ParseInt(epsilonStr, 2, 64)
	if err != nil {
		panic("can't parse binary")
	}
	if verbose {
		fmt.Printf("Gamma: %s (%d)  Epsilon: %s (%d)\n", gammaStr, gamma, epsilonStr, epsilon)
	}
	return int(gamma * epsilon)
}

func part2(input []string) int {
	bits := len(input[0])
	filteredList1 := make([]string, len(input))
	filteredList2 := make([]string, len(input))
	for j, line := range input {
		filteredList1[j] = line
		filteredList2[j] = line
	}

	var o2GenRateStr, co2ScrubRateStr string
	for i := 0; i < bits; i++ {
		bitCounts := map[rune]int{}
		for _, line := range filteredList1 {
			bitCounts[rune(line[i])]++
		}
		if verbose {
			fmt.Printf("bitCounts of remaining values in the list is: %v\n", bitCounts)
		}

		var newFilteredList1 []string
		var filterBit rune
		if bitCounts['0'] > bitCounts['1'] {
			filterBit = '0'
		} else {
			filterBit = '1'
		}
		for _, line := range filteredList1 {
			if rune(line[i]) == filterBit {
				newFilteredList1 = append(newFilteredList1, line)
			}
		}
		if verbose {
			fmt.Printf("After filtering bit %d to only '%c' values, %d values remain\n", i+1, filterBit, len(newFilteredList1))
		}
		if len(newFilteredList1) == 0 {
			panic("all values filtered!")
		}
		if len(newFilteredList1) == 1 {
			o2GenRateStr = newFilteredList1[0]
			break
		}
		filteredList1 = newFilteredList1
	}
	for i := 0; i < bits; i++ {
		bitCounts := map[rune]int{}
		for _, line := range filteredList2 {
			bitCounts[rune(line[i])]++
		}
		if verbose {
			fmt.Printf("bitCounts of remaining values in the list is: %v\n", bitCounts)
		}

		var newFilteredList2 []string
		var filterBit rune
		if bitCounts['0'] <= bitCounts['1'] {
			filterBit = '0'
		} else {
			filterBit = '1'
		}
		for _, line := range filteredList2 {
			if rune(line[i]) == filterBit {
				newFilteredList2 = append(newFilteredList2, line)
			}
		}
		if verbose {
			fmt.Printf("After filtering bit %d to only '%c' values, %d values remain\n", i+1, filterBit, len(newFilteredList2))
		}
		if len(newFilteredList2) == 0 {
			panic("all values filtered!")
		}
		if len(newFilteredList2) == 1 {
			co2ScrubRateStr = newFilteredList2[0]
			break
		}
		filteredList2 = newFilteredList2
	}

	o2GenRate, err := strconv.ParseInt(o2GenRateStr, 2, 64)
	if err != nil {
		panic("can't parse binary")
	}
	co2ScrubRate, err := strconv.ParseInt(co2ScrubRateStr, 2, 64)
	if err != nil {
		panic("can't parse binary")
	}
	if verbose {
		fmt.Printf("o2GenRate: %s (%d)  co2ScrubRate: %s (%d)\n", o2GenRateStr, o2GenRate, co2ScrubRateStr, co2ScrubRate)
	}
	return int(o2GenRate * co2ScrubRate)
}
