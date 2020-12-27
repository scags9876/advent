package main

import (
	"fmt"
	"github.com/scags9876/adventOfCode/lib"
	"strconv"
	"strings"
)

const input     = "219347865"
const testInput = "389125467"

const (
	lowValue = 1
	highValue = 9
	moves = 100
	giantCupsSize = 1_000_000
	//giantCupsSize = 20
	giantMoves = 10_000_000
	//giantMoves = 10
)

const verbose = false

func main() {
	fmt.Println("start")
	cups := parseInput(input)
	part1(cups)
	fmt.Println()
	part2(cups)
}

func part1(initialCups []int) {
	fmt.Printf("Starting Part 1, making %d moves with %d cups\n\n", moves, len(initialCups))

	currentCup := 0
	cups := make([]int, len(initialCups))
	copy(cups, initialCups)
	for turn := 1; turn <= moves; turn++ {
		fmt.Printf("-- move %d --\n", turn)
		cups, currentCup = makeMove(cups, currentCup)
	}
	// reorder the list so that 1 is last
	cups = reorderCups(cups, 1)
	// exclude the last element from the result, since it is the one
	result := lib.JoinInts(cups[:len(cups)-1], "")
	fmt.Printf("Part 1: %s\n", result)
}

func makeMove(cups []int, currentCup int) ([]int, int) {
	size := len(cups)
	if verbose {
		cupList := cupListStr(cups, currentCup)
		fmt.Printf("cups: %s\n", cupList)
	}

	var pickUpList []int
	for i := 0; i < 3; i++ {
		idx := currentCup+1+i
		if idx >= size {
			idx -= size
		}
		pickUpList = append(pickUpList, cups[idx])
	}

	destValue := cups[currentCup] - 1
	if destValue < lowValue {
		destValue = highValue
	}
	for {
		if !lib.IntInSlice(pickUpList, destValue) {
			break
		}
		destValue--
		if destValue < lowValue {
			destValue = highValue
		}
	}
	destCup := findDestCup(cups, destValue)

	for i := 1; i <= size; i++ {
		dest := currentCup+i
		if dest >= size {
			dest -= size
		}
		from := dest+3
		if from >= size {
			from -= size
		}
		cups[dest] = cups[from]
		if cups[dest] == destValue {
			for i, n := range pickUpList {
				pickUpDest := dest+i+1
				if pickUpDest >= size {
					pickUpDest -= size
				}
				cups[pickUpDest] = n
			}
			break
		}
	}

	if verbose {
		fmt.Printf("pick up: %s\n", lib.JoinInts(pickUpList, ", "))
		fmt.Printf("destination: %d (%d)\n", destValue, destCup)
		fmt.Printf("result: %s\n", cupListStr(cups, currentCup))
		fmt.Println()
	}
	currentCup++
	if currentCup >= size {
		currentCup -= size
	}
	return cups, currentCup
}

func part2(initialCups []int) {
	fmt.Printf("Starting Part 2, making %d moves with %d cups\n\n", giantMoves, giantCupsSize)

	// make a linked list where the index is the number and value stored there is the index of the next number in the list
	cups := make([]int, giantCupsSize+1)
	size := len(initialCups)

	firstNum := initialCups[0]
	var lastNum int

	for i := 0; i < size; i++ {
		if i < size-1 {
			cups[initialCups[i]] = initialCups[i+1]
		} else {
			lastNum = initialCups[i]
		}
	}
	for i := size+1; i <= giantCupsSize; i++ {
		cups[lastNum] = i
		lastNum = i
	}
	// close the circle
	cups[lastNum] = firstNum
	size = len(cups)
	// max is 1 less than size because zero is not used
	max := size - 1

	currentCup := firstNum

	if verbose {
		cupList := cupLinkedListStr(cups, currentCup)
		fmt.Printf("Initial cups: %s\n", cupList)
		fmt.Printf("max: %d\n", max)
	}

	for turn := 1; turn <= giantMoves; turn++ {
		if verbose {
			fmt.Printf("-- move %d --\n", turn)
		}

		if verbose {
			cupList := cupLinkedListStr(cups, currentCup)
			fmt.Printf("cups: %s\n", cupList)
		}

		// pick up the first 3 cups after the currentCup
		pickUp1 := cups[currentCup]
		pickUp2 := cups[pickUp1]
		pickUp3 := cups[pickUp2]
		pickUpEnd := cups[pickUp3]

		//relink the current cup to after the pickup list
		cups[currentCup] = pickUpEnd

		// find the destination cup
		destCup := currentCup-1
		if destCup < 1 {
			destCup = max
		}
		for {
			if destCup == pickUp1 || destCup == pickUp2 || destCup == pickUp3 {
				destCup--
				if destCup < 1 {
					destCup = max
				}
				continue
			}
			break
		}

		// insert our pickups into the links
		pickUpEnd = cups[destCup]
		cups[destCup] = pickUp1
		cups[pickUp3] = pickUpEnd

		if verbose {
			fmt.Printf("pick up: %d, %d, %d\n", pickUp1, pickUp2, pickUp3)
			fmt.Printf("destination: %d\n", destCup)
			fmt.Printf("result: %s\n", cupLinkedListStr(cups, currentCup))
			fmt.Println()
		}

		currentCup = cups[currentCup]
	}
	a := cups[1]
	b := cups[cups[1]]
	result := a * b
	fmt.Printf("Part 2: %d * %d = %d\n", a, b, result)
}

func cupListStr(cups []int, currentCup int) string {
	var list strings.Builder
	for i := 0; i < len(cups); i++ {
		if i == currentCup {
			list.WriteString(fmt.Sprintf("(%d)", cups[i]))
		} else {
			list.WriteString(fmt.Sprintf(" %d ", cups[i]))
		}
	}
	return list.String()
}

func cupLinkedListStr(cups []int, currentCup int) string {
	var list strings.Builder
	cup := currentCup
	for i := 0; i < len(cups)-1; i++ {
		if cup == currentCup {
			list.WriteString(fmt.Sprintf("(%d)", cup))
		} else {
			list.WriteString(fmt.Sprintf(" %d ", cup))
		}
		cup = cups[cup]
	}
	return list.String()
}

func findDestCup(cups []int, destValue int) int {
	for i := 0; i < len(cups); i++ {
		if cups[i] == destValue {
			return i
		}
	}
	return -1
}

func reorderCups(cups []int, lastCup int) []int {
	targetIdx := 0
	for i, n := range cups {
		if n == lastCup {
			targetIdx = i
			break
		}
	}
	if targetIdx == len(cups) - 1 {
		// target is already last
		return cups
	}
	reorderedCups := append(cups[targetIdx+1:], cups[:targetIdx+1]...)
	return reorderedCups
}

func parseInput(input string) []int {
	cups := make([]int, len(input))
	for i, n := range input {
		cup, err := strconv.Atoi(string(n))
		if err != nil {
			panic(err)
		}
		cups[i] = cup
	}
	return cups
}
