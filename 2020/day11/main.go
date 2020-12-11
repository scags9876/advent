package main

import (
	"bufio"
	"fmt"
	"os"
)

const inputFilename = "input.txt"

func main() {
	input := getInput()
	part1(input)
	part2(input)
}

const (
	empty rune = 'L'
	occupied = '#'
	floor = '.'
)

func getInput() [][]rune {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	input := make([][]rune, 0)

	for sc.Scan() {
		line := sc.Text()
		row := make([]rune, len(line))
		for col, val := range line {
			row[col] = val
		}

		input = append(input, row)
	}

	return input
}

func part1(seatChart [][]rune) {
	var rounds, changeCount int
	for {
		rounds++
		seatChart, changeCount = flipSeatsPt1(seatChart)

		fmt.Printf("Round %2d: %3d seats changed\n", rounds, changeCount)
		if changeCount == 0 {
			break
		}
	}

	occupiedSeatCount := countOccupiedSeats(seatChart)


	fmt.Printf("Part 1: Stabilization after %d rounds.  %d occupied seats\n", rounds, occupiedSeatCount)
}

func flipSeatsPt1(seatChart [][]rune) ([][]rune, int) {
	newSeatChart := make([][]rune, len(seatChart))
	changCount := 0
	for i, row := range seatChart {
		newRow := make([]rune, len(row))
		for j, seat := range row {
			switch seat {
			case floor:
				newRow[j] = floor
			case empty:
				if adjacentCountPt1(seatChart, i, j) == 0 {
					newRow[j] = occupied
					changCount++
				} else {
					newRow[j] = empty
				}
			case occupied:
				if adjacentCountPt1(seatChart, i, j) >= 4 {
					newRow[j] = empty
					changCount++
				} else {
					newRow[j] = occupied
				}
			}
		}
		newSeatChart[i] = newRow
	}
	return newSeatChart, changCount
}

func adjacentCountPt1(seatChart [][]rune, i, j int) int {
	count := 0

	if i > 0 {
		row := seatChart[i-1]
		if j > 0 && row[j-1] == occupied {
			count++
		}
		if row[j] == occupied {
			count++
		}
		if j < len(row)-1 && row[j+1] == occupied {
			count++
		}
	}
	{
		row := seatChart[i]
		if j > 0 && row[j-1] == occupied {
			count++
		}
		if j < len(row)-1 && row[j+1] == occupied {
			count++
		}
	}
	if i < len(seatChart)-1 {
		row := seatChart[i+1]
		if j > 0 && row[j-1] == occupied {
			count++
		}
		if row[j] == occupied {
			count++
		}
		if j < len(row)-1 && row[j+1] == occupied {
			count++
		}
	}

	return count
}

func countOccupiedSeats(seatChart [][]rune) int {
	count := 0
	for _, row := range seatChart {
		for _, seat := range row {
			if seat == occupied {
				count++
			}
		}
	}
	return count
}

func part2(seatChart [][]rune) {
	var rounds, changeCount int
	fmt.Printf("Start part2 ... \n")
	printSeatChart(seatChart)
	for {
		rounds++
		seatChart, changeCount = flipSeatsPt2(seatChart)

		fmt.Printf("Round %2d: %3d seats changed\n", rounds, changeCount)
		//printSeatChart(seatChart)
		if changeCount == 0 {
			break
		}
	}

	occupiedSeatCount := countOccupiedSeats(seatChart)

	fmt.Printf("Part 2: Stabilization after %d rounds.  %d occupied seats\n", rounds, occupiedSeatCount)
}

func flipSeatsPt2(seatChart [][]rune) ([][]rune, int) {
	newSeatChart := make([][]rune, len(seatChart))
	changCount := 0
	for i, row := range seatChart {
		newRow := make([]rune, len(row))
		for j, seat := range row {
			switch seat {
			case floor:
				newRow[j] = floor
			case empty:
				if adjacentCountPt2(seatChart, i, j) == 0 {
					newRow[j] = occupied
					changCount++
				} else {
					newRow[j] = empty
				}
			case occupied:
				if adjacentCountPt2(seatChart, i, j) >= 5 {
					newRow[j] = empty
					changCount++
				} else {
					newRow[j] = occupied
				}
			}
		}
		newSeatChart[i] = newRow
	}
	return newSeatChart, changCount
}

type vectorStatus int

const (
	 unspecified vectorStatus = iota
	 emptySeat
	 occupiedSeat
	 edge
)

func adjacentCountPt2(seatChart [][]rune, i, j int) int {
	// direction counts, starting up going clockwise
	vectors := make([]vectorStatus, 8)

	rowSize := len(seatChart)
	colSize := len(seatChart[0])

	count := 0
	vectorsDoneCount := 0
	offset := 0
	for {
		offset++
		vectorsDoneCount = 0
		for vector := 0; vector < 8; vector++ {
			if vectors[vector] == unspecified {
				rowCheck, colCheck := i, j
				switch vector {
				case 0, 1, 7: // up, up to the right, up to the left
					rowCheck = i - offset
				case 3, 4, 5: // down, down to the right, down to the left
					rowCheck = i + offset
				}
				switch vector {
				case 1, 2, 3: // up to the right, right, down to the right
					colCheck = j + offset
				case 5, 6, 7: // down to the left, left, up to the left
					colCheck = j - offset
				}
				//fmt.Printf("[%d,%d],[%d,%d], offset %d, vector %d, check [%d,%d]\n", rowSize,colSize, i,j,offset,vector,rowCheck,colCheck)
				if rowCheck < 0 || rowCheck >= rowSize || colCheck < 0 || colCheck >= colSize {
					vectors[vector] = edge
				} else if seatChart[rowCheck][colCheck] == occupied {
					vectors[vector] = occupiedSeat
					count++
				} else if seatChart[rowCheck][colCheck] == empty {
					vectors[vector] = emptySeat
				}
			} else {
				vectorsDoneCount++
			}
		}
		if vectorsDoneCount == 8 {
			break
		}
	}
	//fmt.Printf("[%d,%d] adjacentCount: %d\n", i,j,count)
	return count
}

func printSeatChart(seatChart [][]rune) int {
	count := 0
	fmt.Print("seat chart:\n")
	for _, row := range seatChart {
		for _, seat := range row {
			fmt.Print(string(seat))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")

	return count
}

