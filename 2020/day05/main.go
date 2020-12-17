package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputFilename = "input.txt"

func main() {
	part1()
	part2()
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
func part2() {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	seatIDs := []int{}
	r := strings.NewReplacer("F", "0", "L", "0", "B", "1", "R", "1")

	for sc.Scan() {
		pass := sc.Text()
		seatID, err := strconv.ParseInt(r.Replace(pass), 2, 64)
		if err != nil {
			panic(err)
		}
		seatIDs = Insert(seatIDs, int(seatID))
	}

	lastSeatID := 0
	mySeat := 0
	for _, seatID := range seatIDs {
		fmt.Printf("seatID: %d\n", seatID)
		if lastSeatID != 0 && lastSeatID+1 != seatID {
			fmt.Printf("Gap in the seats!  lastSeatID: %d, this seatID: %d\n", lastSeatID, seatID)
			mySeat = lastSeatID + 1
			break
		}
		lastSeatID = seatID
	}

	fmt.Printf("Part 2: My seatID is: %d\n", mySeat)
}

func Insert(ss []int, s int) []int {
	i := sort.SearchInts(ss, s)
	ss = append(ss, 0)
	copy(ss[i+1:], ss[i:])
	ss[i] = s
	return ss
}
