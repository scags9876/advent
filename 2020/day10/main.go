package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const inputFilename = "input.txt"

func main() {
	input := getInput()
	part1(input)
	part2(input)
}

func getInput() []int {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	input := make([]int, 0)

	for sc.Scan() {
		line := sc.Text()

		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		input = SortedInsert(input, num)
	}

	return input
}

func SortedInsert(ss []int, s int) []int {
	i := sort.SearchInts(ss, s)
	ss = append(ss, 0)
	copy(ss[i+1:], ss[i:])
	ss[i] = s
	return ss
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

//func part2(input []int) {
//	var (
//		finalJoltage int
//	)
//
//	finalJoltage = input[len(input)-1]+3
//
//	ctx, cancelFunc := context.WithCancel(context.Background())
//	defer cancelFunc()
//
//	fmt.Printf("input: %v, finalJoltage: %d", input, finalJoltage)
//
//	comboChan := make(chan combo, len(input)*len(input)*len(input))
//	resultChan := make(chan bool)
//
//	var wg sync.WaitGroup
//
//	for i := 0; i < 5; i++ {
//		go findCombos(ctx, finalJoltage, comboChan, resultChan, &wg)
//	}
//
//	comboCount := 0
//	deadEndCount := 0
//	go func() {
//		for {
//			select {
//			case <-ctx.Done():
//				return
//			case result := <- resultChan:
//				//fmt.Printf("got %d on resultChan\n", count)
//				if result {
//					comboCount++
//				} else {
//					deadEndCount++
//				}
//				// this branch is done, so remove it from the waitgroup
//				wg.Done()
//			}
//		}
//	}()
//
//	initialCombo := combo{
//		combo:          make([]int, 0),
//		remainingInput: input,
//		currentOutput:  0,
//	}
//	// add one to the waitgroup for the first branch
//	wg.Add(1)
//
//	comboChan <- initialCombo
//
//	wg.Wait()
//
//	fmt.Printf("Part 2: There are %d combinations (and %d dead ends) to get to %d joltage.\n", comboCount, deadEndCount, finalJoltage)
//}
//
//func traverseCombinations(input []int, finalJoltage, currentOutput int, combo []int) int {
//	if currentOutput+3 > finalJoltage {
//		//fmt.Printf("Output too high.. combo %v is a dud at output %d\n", combo, currentOutput)
//		return 0
//	}
//	if currentOutput+3 == finalJoltage {
//		//fmt.Printf("We found a combination! Combo %v gives output %d\n", combo, currentOutput+3)
//		return 1
//	}
//	// if there is no more input to try, then return a failure
//	adaptersLeft := len(input)
//	if adaptersLeft == 0 {
//		//fmt.Printf("no adapters left.. combo %v is a dud at output %d\n", combo, currentOutput)
//		return 0
//	}
//	currentCount := 0
//	for offset := 0; offset < adaptersLeft && input[offset] <= currentOutput+3; offset++ {
//		output := input[offset]
//		combo := append(combo, input[offset])
//		//fmt.Printf("adding %d to the combo %v and looking ahead\n", input[offset], combo)
//		workingCombosFromHere := traverseCombinations(input[offset+1:], finalJoltage, output, combo)
//		currentCount += workingCombosFromHere
//		//fmt.Printf("workingCombos from combo %v: %d, currentCount: %d\n", combo, workingCombosFromHere, currentCount)
//	}
//	//fmt.Printf("combo %d resulted in a count of %d\n", combo, currentCount)
//	return currentCount
//}
//
//func findCombos(ctx context.Context, finalJoltage int, comboChan chan combo, resultChan chan bool, wg *sync.WaitGroup) {
//	for {
//		select {
//		case <-ctx.Done():
//			return
//		case c := <-comboChan:
//
//		}
//	}
//}
