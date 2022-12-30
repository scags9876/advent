package main

import (
	"fmt"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 21
const expectedTestResultPart2 = 8

const verbose = true

func main() {
	testInput := lib.GetInputStrings(testInputFilename)
	input := lib.GetInputStrings(inputFilename)

	testTrees := parseTrees(testInput)
	trees := parseTrees(input)

	fmt.Print("Part 1:\n")
	testResult := part1(testTrees)
	if verbose {
		fmt.Printf("Part 1 testInput result: %d\n", testResult)
	}
	if testResult == expectedTestResultPart1 {
		fmt.Printf("testfile returned expected result! (%d)\n", testResult)
	} else {
		panic(fmt.Errorf("expected %d but got %d", expectedTestResultPart1, testResult))
	}

	result := part1(trees)
	fmt.Printf("Part 1 result: %d\n", result)

	fmt.Print("\nPart 2: \n")
	testResult = part2(testTrees)
	if verbose {
		fmt.Printf("Part 2 testInput result: %d\n", testResult)
	}
	if testResult == expectedTestResultPart2 {
		fmt.Printf("testfile returned expected result! (%d)\n", testResult)
	} else {
		panic(fmt.Errorf("expected %d but got %d", expectedTestResultPart2, testResult))
	}

	result = part2(trees)
	fmt.Printf("Part 2 result: %d\n", result)
}

func parseTrees(input []string) [][]int {
	trees := make([][]int, len(input))
	for i, line := range input {
		if verbose {
			fmt.Printf("line %2d\n", i)
		}
		treeline := lib.StringToInts(line, "")
		trees[i] = treeline
	}
	return trees
}

func part1(trees [][]int) int {
	var result int
	for i, treeline := range trees {
		if verbose {
			fmt.Printf("line %2d\n", i)
		}
		for j := range treeline {
			if treeIsVisible(trees, i, j) {
				result++
			}
		}

	}
	return result
}

func treeIsVisible(trees [][]int, i, j int) bool {
	maxRow := len(trees) - 1
	maxCol := len(trees[0]) - 1

	// If it's on the edge, it's visible
	if i == 0 || j == 0 || i == maxRow || j == maxCol {
		return true
	}

	treeHeight := trees[i][j]
	// check to the left
	visibleLeft := true
	for k := i - 1; k >= 0; k-- {
		if trees[k][j] >= treeHeight {
			visibleLeft = false
			break
		}
	}
	if visibleLeft == true {
		return true
	}
	// check to the right
	visibleRight := true
	for k := i + 1; k <= maxCol; k++ {
		if trees[k][j] >= treeHeight {
			visibleRight = false
			break
		}
	}
	if visibleRight == true {
		return true
	}
	// check to the top
	visibleTop := true
	for l := j - 1; l >= 0; l-- {
		if trees[i][l] >= treeHeight {
			visibleTop = false
			break
		}
	}
	if visibleTop == true {
		return true
	}
	// check to the bottom
	visibleBottom := true
	for l := j + 1; l <= maxRow; l++ {
		if trees[i][l] >= treeHeight {
			visibleBottom = false
			break
		}
	}
	if visibleBottom == true {
		return true
	}
	return false
}

func part2(trees [][]int) int {
	var highScenicScore int
	for i, treeline := range trees {
		for j := range treeline {
			scenicScore := scenicScore(trees, i, j)
			if scenicScore > highScenicScore {
				if verbose {
					fmt.Printf("found new high scenic score of %d at [%d,%d]\n", scenicScore, i, j)
				}
				highScenicScore = scenicScore
			}
		}
	}
	return highScenicScore
}

func scenicScore(trees [][]int, i, j int) int {
	maxRow := len(trees) - 1
	maxCol := len(trees[0]) - 1

	var leftScore, rightScore, upScore, downScore int

	treeHeight := trees[i][j]
	// check to the left
	if i == 0 {
		leftScore = 0
	} else {
		for k := i - 1; k >= 0; k-- {
			leftScore++
			if trees[k][j] >= treeHeight {
				break
			}
		}
	}

	// check to the right
	if i == maxCol {
		rightScore = 0
	} else {
		for k := i + 1; k <= maxCol; k++ {
			rightScore++
			if trees[k][j] >= treeHeight {
				break
			}
		}
	}

	// check to the top
	if j == 0 {
		upScore = 0
	} else {
		for l := j - 1; l >= 0; l-- {
			upScore++
			if trees[i][l] >= treeHeight {
				break
			}
		}
	}

	// check to the bottom
	if j == maxRow {
		downScore = 0
	} else {
		for l := j + 1; l <= maxRow; l++ {
			downScore++
			if trees[i][l] >= treeHeight {
				break
			}
		}
	}

	return leftScore * rightScore * upScore * downScore
}
