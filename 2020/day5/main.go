package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFilename = "input.txt"

func main() {
	input := getInput()
	solvePuzzle(input)
}

func getInput() []string {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	var input []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		input = append(input, sc.Text())
	}
	return input
}

func solvePuzzle(input []string) {
	highestID := findHighestSeatID(input)
	fmt.Printf("Part 1: Highest seat ID: %d\n", highestID)

	//fmt.Printf("Part 2: Multiplying all your trees gets you %d ... Ouch!\n", total)
}

func findHighestSeatID(input []string) int {
	highestSeatID := 0
	for _, pass := range input {
		seatID := calcSeatID(pass)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	return highestSeatID
}

func calcSeatID(pass string) int {
	binPass := strings.Replace(pass, "F", "0", -1)
	binPass = strings.Replace(binPass, "B", "1", -1)
	binPass = strings.Replace(binPass, "L", "0", -1)
	binPass = strings.Replace(binPass, "R", "1", -1)

	binRow, binCol := binPass[:7], binPass[7:]

	fmt.Printf("binPass: %s binRow: %s binCol: %s\n", binPass, binRow, binCol)

	row, err := strconv.ParseInt(binRow, 2, 32)
	if err != nil {
		panic(err)
	}
	col, err := strconv.ParseInt(binCol, 2, 32)
	if err != nil {
		panic(err)
	}
	seatID, err := strconv.ParseInt(binPass, 2, 32)
	if err != nil {
		panic(err)
	}

	//seatID := (row * 8) + col

	fmt.Printf("boarding pass %s is row %d and col %d .. seatID: %d\n", pass, row, col, seatID)
	return int(seatID)
}