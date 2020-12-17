package main

import (
	"bufio"
	"fmt"
	"os"
)

const inputFilename = "input.txt"

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	answers := map[rune]bool{}
	total := 0
	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			total += len(answers)
			answers = map[rune]bool{}
			continue
		}
		for _, char := range line {
			answers[char] = true
		}
	}
	total += len(answers)

	fmt.Printf("Part 1: total answer sum: %d\n", total)
}

func part2() {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	answers := map[rune]int{}
	groupSize := 0
	total := 0
	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			total += answerCountPart2(answers, groupSize)
			answers = map[rune]int{}
			groupSize = 0
			continue
		}
		groupSize++
		for _, char := range line {
			if n, ok := answers[char]; ok {
				answers[char] = n + 1
			} else {
				answers[char] = 1
			}
		}
	}
	total += answerCountPart2(answers, groupSize)

	fmt.Printf("Part 2: total answer sum: %d\n", total)

}

func answerCountPart2(answers map[rune]int, groupSize int) int {
	answerCount := 0
	for _, count := range answers {
		if count == groupSize {
			answerCount++
		}
	}
	return answerCount
}
