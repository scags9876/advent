package main

import (
	"fmt"
	"github.com/scags9876/adventOfCode/lib"
	"sort"
	"strings"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 5934
const expectedTestResultPart2 = 26984457539

const verbose = true

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
	fishListStr := strings.Split(input[0], ",")
	fishList := make([]int, len(fishListStr))
	for i, fish := range fishListStr {
		fishList[i] = lib.ToInt(fish)
	}
	return calcFishPop(fishList, 80)
}

func part2(input []string) int {
	fishListStr := strings.Split(input[0], ",")
	fishList := make([]int, len(fishListStr))
	for i, fish := range fishListStr {
		fishList[i] = lib.ToInt(fish)
	}
	return calcFishPop(fishList, 256)
}

func calcFishPop(fishList []int, days int) int {
	if verbose {
		fmt.Printf("Initial state: %s\n", lib.JoinInts(fishList, ","))
	}
	sort.Ints(fishList)
	if verbose {
		fmt.Printf("After sorting: %s\n", lib.JoinInts(fishList, ","))
	}
	fishCounts := make([]int, 9)
	for _, fish := range fishList {
		fishCounts[fish]++
	}

	for day := 1; day <= days; day ++ {
		newFishCounts := []int{
			fishCounts[1], // 0
			fishCounts[2], // 1
			fishCounts[3], // 2
			fishCounts[4], // 3
			fishCounts[5], // 4
			fishCounts[6], // 5
			fishCounts[0] + fishCounts[7], // 6
			fishCounts[8], // 7
			fishCounts[0], // 8
		}
		fishCounts = newFishCounts
		if verbose {
			fmt.Printf("After %2d days, there are %d fish\n", day, sumFish(fishCounts))
		}
	}

	return sumFish(fishCounts)
}

func sumFish(fishCounts []int) int {
	sum := 0
	for _, fish := range fishCounts {
		sum += fish
	}
	return sum
}


