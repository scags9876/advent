package main

import (
	"fmt"
	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"

func main() {
	input := lib.GetInputSortedInts(inputFilename)
	part1(input)
	part2(input)
}

func part1(input []int) {
	counts := map[int]int{
		1: 0,
		2: 0,
		3: 0,
	}

	currentOutput := 0
	for _, joltage := range input {
		diff := joltage - currentOutput
		if diff < 1 {
			panic("jump of less than 1 jolt!")
		} else if diff > 3 {
			panic("jump of more than 3 jolt!")
		}
		counts[diff]++
		currentOutput = joltage
	}

	// jump up to final joltage
	currentOutput += 3
	counts[3]++

	fmt.Printf("Part 1: Current joltage: %d. There are %d 1 jolt jumps, %d 2 jolt jumps and %d 3 jolt jumps.  1-jolt jumps x 3-jolt jumps = %d\n", currentOutput, counts[1], counts[2], counts[3], counts[1]*counts[3])
}

func part2(input []int) {
	fmt.Printf("input: %v\n", input)
	var (
		finalJoltage int
	)

	finalJoltage = input[len(input)-1] + 3

	// break the problem into smaller combos, the edges are anywhere that has a gap of 3
	var subsequences [][]int
	currentOutput := 0
	var currentSequence []int
	for _, joltage := range input {
		diff := joltage - currentOutput
		if diff == 3 {
			subsequences = append(subsequences, currentSequence)
			currentSequence = make([]int, 0)
		}
		currentOutput = joltage
		currentSequence = append(currentSequence, joltage)
	}
	subsequences = append(subsequences, currentSequence)

	fmt.Printf("found %d subsequences in this input\n", len(subsequences))

	var sequenceCounts []int
	comboCount := 1
	for _, sequence := range subsequences {
		initialJoltage, finalJoltage := sequenceEdges(sequence)
		sequenceCount := calcSequenceCombinations(sequence, initialJoltage, finalJoltage)
		sequenceCounts = append(sequenceCounts, sequenceCount)

		fmt.Printf("sequence %v has %d combination(s) to traverse from %d to %d\n", sequence, sequenceCount, initialJoltage, finalJoltage)

		comboCount = comboCount * sequenceCount
	}

	fmt.Printf("Part 2: There are %d combinations to get to %d joltage.\n", comboCount, finalJoltage)
}


type combo struct {
	sequence       []int
	remainingInput []int
	currentOutput  int
}

func sequenceEdges(input []int) (initialJoltage, finalJoltage int) {
	initialJoltage = input[0] - 3
	if initialJoltage < 0 {
		initialJoltage = 0
	}
	finalJoltage = input[len(input)-1] + 3

	return
}

func calcSequenceCombinations(input []int, initialJoltage, finalJoltage int) int {

	initialCombo := combo{
		sequence:       make([]int, 0),
		remainingInput: input,
		currentOutput:  initialJoltage,
	}

	comboCount := 0
	deadEndCount := 0

	stack := []combo{initialCombo}
	for {
		if len(stack) == 0 {
			break
		}
		var c combo
		//fmt.Printf(spew.Sdump(stack))
		c, stack = stack[len(stack)-1], stack[:len(stack)-1]

		//fmt.Printf("stacksize: %7d, sequenceSize: %3d, remaininginput size: %3d, sequence: %v\n", len(stack), len(c.sequence), len(c.remainingInput), c.sequence)
		if c.currentOutput+3 > finalJoltage {
			fmt.Printf("Output too high.. sequence %v is a dud at output %d\n", c.sequence, c.currentOutput)
			deadEndCount++
			continue
		}
		if c.currentOutput+3 == finalJoltage {
			//fmt.Printf("We found a combination! sequence %v gives output %d (stacksize %d)\n", c.sequence, c.currentOutput+3, len(stack))
			comboCount++
			continue
		}
		// if there is no more input to try, then return a failure
		adaptersLeft := len(c.remainingInput)
		if adaptersLeft == 0 {
			fmt.Printf("no adapters left.. sequence %v is a dud at output %d\n", c.sequence, c.currentOutput)
			deadEndCount++
			continue
		}
		for offset := 0; offset < adaptersLeft && c.remainingInput[offset] <= c.currentOutput+3; offset++ {
			newOutput := c.remainingInput[offset]

			newSequence := make([]int, len(c.sequence)+1)
			copy(newSequence, c.sequence)
			newSequence[len(c.sequence)] = newOutput

			newRemainingInput := make([]int, len(c.remainingInput)-(offset+1))
			copy(newRemainingInput, c.remainingInput[1+offset:])

			newCombo := combo{
				sequence:       newSequence,
				remainingInput: newRemainingInput,
				currentOutput:  newOutput,
			}
			//fmt.Printf("offset[%d], new output %d. putting new combo %+v in stack\n", offset, newOutput, newCombo)
			stack = append(stack, newCombo)
		}
	}

	return comboCount
}
