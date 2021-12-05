package main

import (
	"fmt"
	"github.com/scags9876/adventOfCode/lib"
	"strings"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 4512
const expectedTestResultPart2 = 1924

const verbose = false
const boardSize = 5

type board [][]int

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
	var numList []int
	for _, n := range strings.Split(input[0], ",") {
		numList = append(numList, lib.ToInt(n))
	}

	boards := make([]board, 0)

	type coords []int // board, x, y
	numMaps := make(map[int][]coords)

	var boardNum int
	var thisBoard [][]int
	var boardRow int
	for inputLine, line := range input {
		if inputLine == 0 {
			continue // skip first line
		}
		if len(line) == 0 {
			if len(thisBoard) > 0 {
				boards = append(boards, thisBoard)
				boardNum++
			}
			thisBoard = make([][]int, boardSize)
			boardRow = 0
		} else {
			thisBoard[boardRow] = make([]int, boardSize)
			boardCol := 0
			for _, n := range strings.Split(line, " ") {
				if len(n) == 0 {
					continue
				}
				num := lib.ToInt(n)
				thisBoard[boardRow][boardCol] = num

				if numMaps[num] == nil {
					numMaps[num] = make([]coords, 0)
				}
				numMaps[num] = append(numMaps[num], []int{boardNum, boardRow, boardCol})

				boardCol++
			}
			boardRow++
		}
	}
	boards = append(boards, thisBoard)

	if verbose {
		fmt.Printf("found %d input numbers and %d boards\n", len(numList), len(boards))
		for i, b := range boards {
			fmt.Printf("Board %d:\n", i)
			printBoard(b)
		}
	}

	for _, num := range numList {
		if numMaps[num] == nil {
			if verbose {
				fmt.Printf("%d not found in any board\n", num)
			}
			continue
		}
		if verbose {
			fmt.Printf("Calling number: %d\n", num)
		}
		for _, foundNum := range numMaps[num] {
			boardNum, x, y := foundNum[0], foundNum[1], foundNum[2]
			if verbose {
				fmt.Printf("Found on board %d[%d][%d]\n", boardNum, x, y)
				printBoard(boards[boardNum])
			}
			boards[boardNum][x][y] = boards[boardNum][x][y] * -1 // flip to negative to mark it
			if checkBoard(boards[boardNum], x, y) {
				if verbose {
					fmt.Printf("Board %d is the winner!\n", boardNum)
				}
				return calcScore(num, boards[boardNum])
			}
		}
	}

	return 0
}

func checkBoard(b board, x, y int) bool {
	horizWin := true
	for i := 0; i < boardSize; i++ {
		if b[x][i] > 0 {
			horizWin = false
		}	
	}	
	if horizWin {
		return true
	}
	vertWin := true
	for i := 0; i < boardSize; i++ {
		if b[i][y] > 0 {
			vertWin = false
		}
	}
	if vertWin {
		return true
	}
	return false
}

func calcScore(num int, b board) int {
	sum := 0
	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			if b[x][y] > 0 {
				sum = sum + b[x][y]
			}
		}
	}
	return sum * num
}

func printBoard(b board) {
	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			fmt.Printf("%4d", b[x][y])
		}
		fmt.Print("\n")
	}
}

func part2(input []string) int {
	var numList []int
	for _, n := range strings.Split(input[0], ",") {
		numList = append(numList, lib.ToInt(n))
	}

	boards := make([]board, 0)

	type coords []int // board, x, y
	numMaps := make(map[int][]coords)

	var boardNum int
	var thisBoard [][]int
	var boardRow int
	for inputLine, line := range input {
		if inputLine == 0 {
			continue // skip first line
		}
		if len(line) == 0 {
			if len(thisBoard) > 0 {
				boards = append(boards, thisBoard)
				boardNum++
			}
			thisBoard = make([][]int, boardSize)
			boardRow = 0
		} else {
			thisBoard[boardRow] = make([]int, boardSize)
			boardCol := 0
			for _, n := range strings.Split(line, " ") {
				if len(n) == 0 {
					continue
				}
				num := lib.ToInt(n)
				thisBoard[boardRow][boardCol] = num

				if numMaps[num] == nil {
					numMaps[num] = make([]coords, 0)
				}
				numMaps[num] = append(numMaps[num], []int{boardNum, boardRow, boardCol})

				boardCol++
			}
			boardRow++
		}
	}
	boards = append(boards, thisBoard)

	if verbose {
		fmt.Printf("found %d input numbers and %d boards\n", len(numList), len(boards))
		for i, b := range boards {
			fmt.Printf("Board %d:\n", i)
			printBoard(b)
		}
	}

	boardsLeft := make([]int, len(boards))
	for i := range boards {
		boardsLeft[i] = i
	}
	for _, num := range numList {
		if numMaps[num] == nil {
			if verbose {
				fmt.Printf("%d not found in any board\n", num)
			}
			continue
		}
		if verbose {
			fmt.Printf("Calling number: %d\n", num)
		}
		for _, foundNum := range numMaps[num] {
			boardNum, x, y := foundNum[0], foundNum[1], foundNum[2]
			if verbose {
				fmt.Printf("Found on board %d[%d][%d]\n", boardNum, x, y)
				printBoard(boards[boardNum])
			}
			boards[boardNum][x][y] = boards[boardNum][x][y] * -1 // flip to negative to mark it
			if !lib.IntInSlice(boardsLeft, boardNum) {
				if verbose {
					fmt.Printf("Board %d is already a winner, skipping check\n", boardNum)
				}
				continue
			}
			if checkBoard(boards[boardNum], x, y) {
				if verbose {
					fmt.Printf("Board %d is a winner!\n", boardNum)
				}
				// remove this boardNum from the list of boards left
				newBoardsLeft := make([]int, 0)
				for _, b := range boardsLeft {
					if boardNum != b {
						newBoardsLeft = append(newBoardsLeft, b)
					}
				}
				boardsLeft = newBoardsLeft
				if len(boardsLeft) == 0 {
					if verbose {
						fmt.Printf("Board %d is the last winner!\n", boardNum)
					}
					return calcScore(num, boards[boardNum])
				}

			}
		}
		if len(boardsLeft) == 1 {
			if verbose {
				fmt.Printf("One board left! (board %d)", boardsLeft[0])
			}
		}
	}

	return 0
}
