package main

import (
	"fmt"
	"github.com/scags9876/adventOfCode/lib"
	"regexp"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 5
const expectedTestResultPart2 = 12

const verbose = false

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
	re := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

	// [y][x]numLines
	board := make(map[int]map[int]int)
	var maxX, maxY int
	var numCrossPoints int

	for i, line := range input {
		matches := re.FindStringSubmatch(line)
		x1, y1, x2, y2 := lib.ToInt(matches[1]), lib.ToInt(matches[2]), lib.ToInt(matches[3]), lib.ToInt(matches[4])
		if verbose {
			fmt.Printf("line %d goes from %d,%d to %d,%d\n", i, x1, y1, x2, y2)
		}
		if x1 > maxX {
			maxX = x1
		}
		if y1 > maxY {
			maxY = y1
		}
		
		if x1 != x2 && y1 != y2 {
			if verbose {
				fmt.Printf("line %d is a diagonal, skipping to next\n", i)
			}
			continue
		}
		currX, currY := x1, y1
		if board[currY] == nil {
			board[currY] = make(map[int]int)
		}
		board[currY][currX] = board[currY][currX] + 1
		if board[currY][currX] == 2 {
			numCrossPoints++
		}
		for !(currX == x2 && currY == y2) {
			if currX < x2 {
				currX++
			} else if currX > x2 {
				currX--
			} else if currY < y2 {
				currY++
			} else if currY > y2 {
				currY--
			} else {
				panic("nothing moved from last time!")
			}
			if board[currY] == nil {
				board[currY] = make(map[int]int)
			}
			board[currY][currX] = board[currY][currX] + 1
			if board[currY][currX] == 2 {
				numCrossPoints++
			}
		}
		if verbose {
			fmt.Printf("After line %d, there are %d cross points\n", i, numCrossPoints)
			if maxX < 80 {
				fmt.Printf("After line %d, board is: \n", i)
				printBoard(board, maxX, maxY)
			}
		}
	}

	return numCrossPoints
}

func printBoard(b map[int]map[int]int, maxX, maxY int) {
	for y := 0; y <= maxY; y++ {
		if b[y] == nil {
			b[y] = make(map[int]int)
		}
		for x := 0; x <= maxX; x++ {
			if b[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(b[y][x])
			}
		}
		fmt.Print("\n")
	}
}

func part2(input []string) int {
	re := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

	// [y][x]numLines
	board := make(map[int]map[int]int)
	var maxX, maxY int
	var numCrossPoints int

	for i, line := range input {
		matches := re.FindStringSubmatch(line)
		x1, y1, x2, y2 := lib.ToInt(matches[1]), lib.ToInt(matches[2]), lib.ToInt(matches[3]), lib.ToInt(matches[4])
		if verbose {
			fmt.Printf("line %d goes from %d,%d to %d,%d\n", i, x1, y1, x2, y2)
		}
		if x1 > maxX {
			maxX = x1
		}
		if y1 > maxY {
			maxY = y1
		}

		currX, currY := x1, y1
		if board[currY] == nil {
			board[currY] = make(map[int]int)
		}
		board[currY][currX] = board[currY][currX] + 1
		if board[currY][currX] == 2 {
			numCrossPoints++
		}
		for !(currX == x2 && currY == y2) {
			if currX < x2 {
				currX++
			} else if currX > x2 {
				currX--
			}
			if currY < y2 {
				currY++
			} else if currY > y2 {
				currY--
			}
			if board[currY] == nil {
				board[currY] = make(map[int]int)
			}
			board[currY][currX] = board[currY][currX] + 1
			if board[currY][currX] == 2 {
				numCrossPoints++
			}
		}
		if verbose {
			fmt.Printf("After line %d, there are %d cross points\n", i, numCrossPoints)
			if maxX < 80 {
				fmt.Printf("After line %d, board is: \n", i)
				printBoard(board, maxX, maxY)
			}
		}
	}

	return numCrossPoints
}

