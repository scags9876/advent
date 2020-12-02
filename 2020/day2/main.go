package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const inputFilename = "input.txt"

var re = regexp.MustCompile(`(\d+)-(\d+)\s(\w):\s(\w+)`)

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
	matches := re.FindStringSubmatch(line)
	min, _ := strconv.Atoi(matches[1])
	max, _ := strconv.Atoi(matches[2])
	letter := matches[3]
	pw := matches[4]

	count := 0
	positionCount := 0
	for i, char := range pw {
		if string(char) == letter {
			count++
			if i+1 == min {
				positionCount++
			}
			if i+1 == max {
				positionCount++
			}
		}
	}

	if count >= min && count <= max {
		validPass1 = true
	}
	if positionCount == 1 {
		validPass2 = true
	}
	return
}