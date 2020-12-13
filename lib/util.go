package lib

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func GetInputStrings(inputFilename string) []string {
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

func GetInputInts(inputFilename string) []int {
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


func GetInputSortedInts(inputFilename string) []int {
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
		input = SortedInsertInt(input, num)
	}

	return input
}

func SortedInsertInt(ss []int, s int) []int {
	i := sort.SearchInts(ss, s)
	ss = append(ss, 0)
	copy(ss[i+1:], ss[i:])
	ss[i] = s
	return ss
}