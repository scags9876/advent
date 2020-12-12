package main

import (
	"fmt"
	"github.com/scags9876/adventOfCode/lib"
	"strings"
)

const inputFilename = "input.txt"

func main() {
	input := lib.GetInputStrings(inputFilename)
	solvePuzzle(input)
}

func solvePuzzle(input []string) {
	validPasswordCount1 := 0
	validPasswordCount2 := 0
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		validPass1, validPass2 := checkLine(line)
		if validPass1 {
			validPasswordCount1++
		}
		if validPass2 {
			validPasswordCount2++
		}
	}
	fmt.Printf("\tFound %d valid passwords from the first algorithm\n", validPasswordCount1)
	fmt.Printf("\tFound %d valid passwords from the second algorithm\n", validPasswordCount2)
}

func checkLine(line string) (validPass1, validPass2 bool) {
	var (
		min, max int
		letter rune
		pw string
	)
	_, err := fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &letter, &pw)
	if err != nil {
		panic(err)
	}

	count := strings.Count(pw, string(letter))
	positionCount := 0
	if len(pw) >= min && rune(pw[min-1]) == letter {
		positionCount++
	}
	if len(pw) >= max && rune(pw[max-1]) == letter {
		positionCount++
	}

	if count >= min && count <= max {
		validPass1 = true
	}
	if positionCount == 1 {
		validPass2 = true
	}
	return
}