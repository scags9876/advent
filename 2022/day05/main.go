package main

import (
	"fmt"
	"strings"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = "CMZ"
const expectedTestResultPart2 = "MCD"

const verbose = true

func main() {
	testInput := lib.GetInputStrings(testInputFilename)
	input := lib.GetInputStrings(inputFilename)

	testStacks, testMoves := parseInput(testInput)
	fmt.Print("Part 1:\n")
	testResult := part1(testStacks, testMoves)
	if verbose {
		fmt.Printf("Part 1 testInput result: %s\n", testResult)
	}
	if testResult == expectedTestResultPart1 {
		fmt.Printf("testfile returned expected result! (%s)\n", testResult)
	} else {
		panic(fmt.Errorf("expected %s but got %s", expectedTestResultPart1, testResult))
	}

	stacks, moves := parseInput(input)
	result := part1(stacks, moves)
	fmt.Printf("Part 1 result: %s\n", result)

	fmt.Print("\nPart 2: \n")
	testStacks, testMoves = parseInput(testInput)
	testResult = part2(testStacks, testMoves)
	if verbose {
		fmt.Printf("Part 2 testInput result: %s\n", testResult)
	}
	if testResult == expectedTestResultPart2 {
		fmt.Printf("testfile returned expected result! (%s)\n", testResult)
	} else {
		panic(fmt.Errorf("expected %s but got %s", expectedTestResultPart2, testResult))
	}

	stacks, moves = parseInput(input)
	result = part2(stacks, moves)
	fmt.Printf("Part 2 result: %s\n", result)
}

func part1(stacks []stack, moves []string) string {
	var result string
	for i, move := range moves {
		var quantity, start, dest int
		_, err := fmt.Sscanf(move, "move %d from %d to %d", &quantity, &start, &dest)
		if err != nil {
			panic(err)
		}
		if verbose {
			fmt.Printf("move %2d: move %d from %d to %d\n", i, quantity, start, dest)
		}
		for q := quantity; q > 0; q-- {
			c, ok := stacks[start-1].Pop()
			if !ok {
				panic(fmt.Errorf("stack %d was empty", start-1))
			}
			stacks[dest-1].Push(c)
		}
	}

	result = readTopContainers(stacks)
	return result
}

func part2(stacks []stack, moves []string) string {
	var result string
	for i, move := range moves {
		var quantity, start, dest int
		_, err := fmt.Sscanf(move, "move %d from %d to %d", &quantity, &start, &dest)
		if err != nil {
			panic(err)
		}
		if verbose {
			fmt.Printf("move %2d: move %d from %d to %d\n", i, quantity, start, dest)
		}
		containers, ok := stacks[start-1].PopN(quantity)
		if !ok {
			panic(fmt.Errorf("stack %d didn't have enough containers", start-1))
		}
		stacks[dest-1].PushN(containers)

		if verbose {
			fmt.Printf("stack %d is now %s, stack %d is now %s\n", start-1, stacks[start-1].String(), dest-1, stacks[dest-1].String())
		}
	}

	result = readTopContainers(stacks)
	return result
}

func readTopContainers(stacks []stack) (result string) {
	for i, s := range stacks {
		c, ok := s.Pop()
		if !ok {
			panic(fmt.Errorf("stack %d was empty", i))
		}
		fmt.Printf("stack %d top container was %c\n", i, c)
		result += string(c)
	}
	return result
}

type stack struct {
	stack []rune
}

func (s *stack) String() string {
	ret := ""
	for _, c := range s.stack {
		ret += fmt.Sprintf("[%c] ", c)
	}
	return ret
}

func (s *stack) Push(c rune) {
	s.stack = append(s.stack, c)
}

func (s *stack) Pop() (rune, bool) {
	var c rune
	stackSize := len(s.stack)
	if stackSize < 1 {
		return c, false
	} else {
		c = s.stack[stackSize-1]
		s.stack = s.stack[:stackSize-1]
	}

	return c, true
}

func (s *stack) PushN(cn []rune) {
	for _, c := range cn {
		s.stack = append(s.stack, c)
	}
}

func (s *stack) PopN(n int) ([]rune, bool) {
	var c []rune
	stackSize := len(s.stack)
	if stackSize < n {
		return c, false
	} else {
		c = s.stack[stackSize-n:]
		s.stack = s.stack[:stackSize-n]
	}

	return c, true
}

func parseInput(input []string) (stacks []stack, moves []string) {
	var movesStart int
	for i, line := range input {
		if strings.HasPrefix(line, "move") {
			movesStart = i
			break
		}
	}
	if verbose {
		fmt.Printf("moves start on line %d\n", movesStart)
	}

	moves = input[movesStart:]
	if verbose {
		fmt.Printf("stacks: '%s'\n", input[movesStart-2])
	}

	stackNames := strings.Split(input[movesStart-2], "  ")
	numStacks := len(stackNames)

	if verbose {
		fmt.Printf("there are %d stacks and %d moves\n", numStacks, len(moves))
	}

	stacks = make([]stack, numStacks)

	for i := movesStart - 3; i >= 0; i-- {
		line := input[i]
		lineLen := len(line)
		if verbose {
			fmt.Printf("containers on level %d: '%s'\n", (i-(movesStart-3))*-1, line)
		}
		for j := 0; j < numStacks; j++ {
			if lineLen < (j * 4) {
				break
			}
			c := line[(j*4)+1]
			if verbose {
				fmt.Printf("level %d, container %d: '%c'\n", (i-(movesStart-3))*-1, j, c)
			}
			if c == ' ' {
				continue
			}
			stacks[j].Push(rune(c))
		}
	}

	if verbose {
		for i, s := range stacks {
			fmt.Printf("stack %d: %s\n", i, s.String())
		}
	}

	return stacks, moves

}
