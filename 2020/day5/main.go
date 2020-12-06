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
	part1()
}

func part1() {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	highestSeatID := int64(0)
	r := strings.NewReplacer("F", "0", "L", "0", "B", "1", "R", "1")

	for sc.Scan() {
		pass := sc.Text()
		seatID, err := strconv.ParseInt(r.Replace(pass), 2, 64)
		if err != nil {
			panic(err)
		}
		if seatID > highestSeatID {
			highestSeatID = seatID
		}

		row := seatID >> 3
		col := seatID - row*8
		fmt.Printf("boarding pass %s is row %d and col %d .. seatID: %d\n", pass, row, col, seatID)
	}

	fmt.Printf("Part 1: Highest seat ID: %d\n", highestSeatID)
}