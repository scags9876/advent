package main

import (
	"fmt"
	"github.com/scags9876/adventOfCode/lib"
	"regexp"
	"strconv"
	"strings"
)

const inputFilename = "input.txt"

func main() {
	input := lib.GetInputStrings(inputFilename)
	part1(input)
	part2(input)
}

var memRegex = regexp.MustCompile(`^mem\[(\d+)] = (\d+)$`)

func part1(input []string) {
	mem := make(map[int]int64)
	var maskOR, maskAND int64
	for _, line := range input {
		if strings.HasPrefix(line, "mask") {
			var mask string
			_, err := fmt.Sscanf(line,"mask = %s", &mask)
			if err != nil {
				panic(fmt.Errorf("invalid mask input: %s", line))
			}
			maskORStr := strings.Replace(mask, "X", "0", -1)
			maskANDStr := strings.Replace(mask, "X", "1", -1)

			maskOR, err = strconv.ParseInt(maskORStr, 2, 64)
			if err != nil {
				panic(err)
			}
			maskAND, err = strconv.ParseInt(maskANDStr, 2, 64)
			if err != nil {
				panic(err)
			}
			fmt.Printf("new mask parsed: %s\n", mask)

		} else if strings.HasPrefix(line, "mem") {
			matches := memRegex.FindStringSubmatch(line)
			if len(matches) == 0 {
				panic(fmt.Errorf("invalid mem input: %s", line))
			}
			memKey, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			memValue, err := strconv.ParseInt(matches[2], 10, 64)
			if err != nil {
				panic(err)
			}

			memValue &= maskAND
			memValue |= maskOR

			fmt.Printf("mem operation: mem[%d] = %d\n", memKey, memValue)

			// keep mem small if we are zeroing it out
			if memValue == 0 {
				delete(mem, memKey)
			} else {
				mem[memKey] = memValue
			}
		} else {
			panic("unknown input")
		}
	}

	result := int64(0)
	for _, value := range mem {
		result += value
	}
	fmt.Printf("Part 1: %d\n", result)
}

func part2(input []string) {
	mem := make(map[int64]int64)
	var mask string
	for lineN, line := range input {
		if strings.HasPrefix(line, "mask") {
			_, err := fmt.Sscanf(line,"mask = %s", &mask)
			if err != nil {
				panic(fmt.Errorf("invalid mask input: %s", line))
			}

			fmt.Printf("%d new mask parsed: %s\n", lineN, mask)

		} else if strings.HasPrefix(line, "mem") {
			matches := memRegex.FindStringSubmatch(line)
			if len(matches) == 0 {
				panic(fmt.Errorf("invalid mem input: %s", line))
			}
			memKey, err := strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				panic(err)
			}
			memValue, err := strconv.ParseInt(matches[2], 10, 64)
			if err != nil {
				panic(err)
			}

			memKeys := []int64{memKey}

			for i := 0; i < len(mask); i++ {
				maskChar := mask[i]
				position := 35-i
				switch maskChar {
				case '1':
					for j, k := range memKeys {
						//fmt.Printf("char %c at %d, key %d becomes %d\n", maskChar, position, memKeys[j], k | 1<<position)
						memKeys[j] = k | 1<<position
					}
				case 'X':
					newMemKeys := []int64{}
					for _, k := range memKeys {
						optionA := setBit(k, position)
						optionB := clearBit(k, position)
						//fmt.Printf("char %c at %d, key %d becomes %d and %d\n", maskChar, position, k, optionA, optionB)
						newMemKeys = append(newMemKeys, optionA, optionB)
					}
					memKeys = newMemKeys
				}
			}
			fmt.Printf("%d mem operation: mem[%d] = %d results in %d ops\n", lineN, memKey, memValue, len(memKeys))

			for _, key := range memKeys {
				// keep mem small if we are zeroing it out
				if memValue == 0 {
					delete(mem, key)
				} else {
					mem[key] = memValue
				}
			}
		} else {
			panic("unknown input")
		}
	}

	result := int64(0)
	for _, value := range mem {
		result += value
	}
	fmt.Printf("\nPart 2: %d\n", result)
}
// Sets the bit at pos in the integer n.
func setBit(n int64, pos int) int64 {
	mask := int64(1 << pos)
	n |= mask
	return n
}
// Clears the bit at pos in n.
func clearBit(n int64, pos int) int64 {
	mask := int64(^(1 << pos))
	n &= mask
	return n
}
