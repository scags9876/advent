package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputFilename = "input.txt"

const preamble = 25
const lookback = 25

func main() {
	magicNum := part1()
	part2(magicNum)
}

func part1() int {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	lookbackSet := make([]int, lookback)
	i := 0
	var answer int
	for sc.Scan() {
		line := sc.Text()

		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if i < preamble {
			lookbackSet[i] = num
		} else {
			if !sumExists(lookbackSet, num) {
				answer = num
				break
			} else {
				lookbackSet = append(lookbackSet[1:], num)
			}
		}
		i++
	}
	fmt.Printf("Part 1: first number not a sum of 2 previous %d numbers: %d\n", lookback, answer)
	return answer
}

func sumExists(set []int, sum int) bool {
	for i := 0; i < len(set) - 1; i++ {
		a := set[i]
		for j := i+1; j < len(set); j++ {
			b := set[j]
			if a+b == sum {
				return true
			}
		}
	}
	return false
}

func part2(magicNum int) {
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
		input = append(input, num)
	}

	set := make([]int, 0)
	sum := 0
	for _, num := range input {
		set = append(set, num)
		sum = setSum(set)
		//fmt.Printf("set is %v, sum is %d\n", set, sum)
		if sum == magicNum {
			break
		} else if sum > magicNum {
			for sum > magicNum {
				set = set[1:]
				sum = setSum(set)
				//fmt.Printf("after removing first element, set is %v, sum is %d\n", set, sum)
			}
			if sum == magicNum {
				break
			}
		}
	}
	smallest, largest := edges(set)
	answer := smallest + largest
	fmt.Printf("Part 2: contiguous set is %+v.  First and last added together is %d\n", set, answer)
}

func setSum(set []int) int {
	sum := 0
	for _, n := range set {
		sum += n
	}
	return sum
}

func edges(set []int) (int, int) {
	smallest, largest := set[0], set[len(set)-1]
	for _, n := range set {
		if n > largest {
			largest = n
		}
		if n < smallest {
			smallest = n
		}
	}
	return smallest, largest
}