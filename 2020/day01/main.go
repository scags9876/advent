package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputFilename = "input.txt"

func main() {
	input := getInput()
	solvePuzzle(input)
}

func getInput() []int {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	var input []int
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		input = append(input, num)
	}
	return input
}

func solvePuzzle(input []int) {
	var product1, product2 int
	inputSize := len(input)
	for ia, a := range input {
		for ib := ia + 1; ib < inputSize; ib++ {
			b := input[ib]
			if product1 == 0 {
				if a+b == 2020 {
					product1 = a * b
					fmt.Printf("%d + %d = 2020, %d * %d = %d\n", a, b, a, b, product1)
					continue
				}
			}
			if product2 == 0 && a+b < 2020 {
				for ic := ib + 1; ic < inputSize; ic++ {
					c := input[ic]
					if a+b+c == 2020 {
						product2 = a * b * c
						fmt.Printf(
							"%d + %d + %d = 2020, %d * %d * %d = %d\n",
							a, b, c,
							a, b, c,
							product2,
						)
						break
					}

				}
			}
		}
		if product1 != 0 && product2 != 0 {
			break
		}
	}
}
