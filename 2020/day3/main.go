package main

import (
	"bufio"
	"fmt"
	"os"
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
	treeCount := traverseWithSlope(input, 3, 1)
	fmt.Printf("Part 1: You hit %d trees in your slope\n", treeCount)

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	total := 1
	for _, slope := range slopes {
		treeCount = traverseWithSlope(input, slope[0], slope[1])
		total = treeCount * total
	}
	fmt.Printf("Part 2: Multiplying all your trees gets you %d ... Ouch!\n", total)
}

func traverseWithSlope(input []string, xShift, yShift int) int {
	x := 0
	treeCount := 0
	for y := 0; y < len(input); y += yShift {
		line := input[y]
		if line[x] == '#' {
			treeCount++
		}
		x += xShift
		if x >= len(line) {
			x = x % len(line)
		}
	}

	fmt.Printf("Using a slope of right %d down %d, you hit %d trees\n", xShift, yShift, treeCount)
	return treeCount
}