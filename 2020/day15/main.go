package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	testInput1 = "0,3,6"
	testInput2 = "1,3,2"
	testInput3 = "3,1,2"
	input      = "0,1,5,10,3,12,19"
)

func main() {
	part1FinalPosition := 2020
	part2FinalPosition := 30000000
	if solve(testInput1, part1FinalPosition) != 436 {
		panic("wrong")
	}
	if solve(testInput2, part1FinalPosition) != 1 {
		panic("wrong")
	}
	if solve(testInput3, part1FinalPosition) != 1836 {
		panic("wrong")
	}
	result := solve(input, part1FinalPosition)
	fmt.Printf("Part 1: %d\n", result)

	result = solve(input, part2FinalPosition)
	fmt.Printf("Part 2: %d\n", result)
}

func solve(input string, finalPosition int) int {
	sequence := make([]int, 0)
	lastSeen := make(map[int]int)

	for _, num := range strings.Split(input, ",") {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		sequence = append(sequence, n)
	}

	for position := 1; position < finalPosition; position++ {
		n := sequence[position-1]
		nextNum := 0
		if numSeenAtPosition, ok := lastSeen[n]; ok {
			nextNum = position - numSeenAtPosition
			lastSeen[n] = position
		} else {
			lastSeen[n] = position
		}

		// only append if it's the last element in the list
		if len(sequence) == position {
			sequence = append(sequence, nextNum)
		}
	}

	result := sequence[len(sequence)-1]
	return result
}
