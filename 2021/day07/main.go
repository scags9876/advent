package main

import (
	"fmt"
	"github.com/scags9876/adventOfCode/lib"
	"math"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 37
const expectedTestResultPart2 = 168

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
	crabList := lib.StringToSortedInts(input[0], ",")

	var minFuelScore, minFuelPosition int
	for p := crabList[0]; p <= crabList[len(crabList) - 1]; p++ {
		fuelScore := calcTotalFuel(crabList, p)
		if minFuelScore == 0 || fuelScore < minFuelScore {
			minFuelScore = fuelScore
			minFuelPosition = p
		}
	}
	if verbose {
		fmt.Printf("min fuel of %d found at position %d\n", minFuelScore, minFuelPosition)
	}

	return minFuelScore
}

func calcTotalFuel(crabList []int, p int) int {
	var fuel int
	for _, crab := range crabList {
		cost := int(math.Abs(float64(p - crab)))
		fuel = fuel + cost
	}
	return fuel
}

func part2(input []string) int {
	crabList := lib.StringToSortedInts(input[0], ",")

	var minFuelScore, minFuelPosition int
	for p := crabList[0]; p <= crabList[len(crabList) - 1]; p++ {
		fuelScore := calcTotalProgressiveFuel(crabList, p)

		if minFuelScore == 0 || fuelScore < minFuelScore {
			minFuelScore = fuelScore
			minFuelPosition = p
		}
	}
	if verbose {
		fmt.Printf("min fuel of %d found at position %d\n", minFuelScore, minFuelPosition)
	}

	return minFuelScore

}


func calcTotalProgressiveFuel(crabList []int, p int) int {
	var fuel int
	for _, crab := range crabList {
		cost := int(math.Abs(float64(p - crab)))
		cost = progressiveFuelScore(cost)
		fuel = fuel + cost
	}
	return fuel
}

var scores = map[int]int{
	1: 1,
	2: 3,
}

func progressiveFuelScore(f int) int {
	if f == 0 {
		return 0
	} else if f < 0 {
		panic("negative fuel!")
	}
	if scores[f] != 0 {
		return scores[f]
	}
	score := f + progressiveFuelScore(f-1)
	scores[f] = score
	return score
}

