package main

import (
	"fmt"
	"strings"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 15
const expectedTestResultPart2 = 12

const verbose = true
const (
	win      = 6
	draw     = 3
	loss     = 0
	rock     = 1
	paper    = 2
	scissors = 3
)

var key1 = map[string]int{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var key2 = map[string]int{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": loss,
	"Y": draw,
	"Z": win,
}

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
		outcome := 0
		moves := strings.Split(line, " ")
		opponentMove := key1[moves[0]]
		myMove := key1[moves[1]]
		if verbose {
			fmt.Printf("line %2d, opponent: %d, me: %d\n", i, opponentMove, myMove)
		}
		switch opponentMove {
		case rock:
			switch myMove {
			case rock:
				outcome = draw
			case paper:
				outcome = win
			case scissors:
				outcome = loss
			}
		case paper:
			switch myMove {
			case rock:
				outcome = loss
			case paper:
				outcome = draw
			case scissors:
				outcome = win
			}
		case scissors:
			switch myMove {
			case rock:
				outcome = win
			case paper:
				outcome = loss
			case scissors:
				outcome = draw
			}
		}
		if verbose {
			fmt.Printf("line %2d, outcome: %d\n", i, outcome)
		}
		outcome += myMove
		result += outcome
	}
	return result
}

func part2(input []string) int {
	var result int
	for i, line := range input {
		moves := strings.Split(line, " ")
		opponentMove := key2[moves[0]]
		outcome := key2[moves[1]]
		myMove := 0
		if verbose {
			fmt.Printf("line %2d, opponent: %d, outcome: %d\n", i, opponentMove, outcome)
		}
		switch opponentMove {
		case rock:
			switch outcome {
			case loss:
				myMove = scissors
			case draw:
				myMove = rock
			case win:
				myMove = paper
			}
		case paper:
			switch outcome {
			case loss:
				myMove = rock
			case draw:
				myMove = paper
			case win:
				myMove = scissors
			}
		case scissors:
			switch outcome {
			case loss:
				myMove = paper
			case draw:
				myMove = scissors
			case win:
				myMove = rock
			}
		}
		if verbose {
			fmt.Printf("line %2d, myMove: %d\n", i, myMove)
		}
		outcome += myMove
		result += outcome
	}
	return result
}
