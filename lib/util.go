package lib

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
		input = append(input, ToInt(sc.Text()))
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
		input = SortedInsertInt(input, ToInt(sc.Text()))
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

func ToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func ToIntOk(s string) (int, bool) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return n, true
}

func JoinInts(si []int, s string) string {
	ss := make([]string, len(si))
	for i, n := range si {
		ss[i] = strconv.Itoa(n)
	}
	return strings.Join(ss, s)
}

func StringInSlice(set []string, s string) bool {
	for _, el := range set {
		if el == s {
			return true
		}
	}
	return false
}

func IntInSlice(set []int, s int) bool {
	for _, el := range set {
		if el == s {
			return true
		}
	}
	return false
}