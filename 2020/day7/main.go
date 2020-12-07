package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const inputFilename = "input.txt"

func main() {
	part1()
	part2()
}

const myBag = "shiny gold"

type bag struct {
	desc string
	canHold map[string]int
	flatCanHold map[string]int
	bagCount *int
}

func part1() {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	rLine := regexp.MustCompile(`^(\w+\s\w+) bags contain (no other bags.|.+)+$`)
	rBags := regexp.MustCompile(`(?:(\d+) (\w+\s\w+) bags?.\s?)`)
	total := 0
	bags := make(map[string]bag)
	for sc.Scan() {
		line := sc.Text()
		matches := rLine.FindAllStringSubmatch(line, -1)
		//fmt.Printf("matches is %#v\n", matches)
		thisBag := bag{
			desc: matches[0][1],
		}
		if matches[0][2] == "no other bags." {
			// leaf
		} else {
			bagsMatches := rBags.FindAllStringSubmatch(matches[0][2], -1)
			//fmt.Printf("bagsMatches is %#v\n", bagsMatches)

			thisBag.canHold = make(map[string]int)
			for _, container := range bagsMatches {
				thisBag.canHold[container[2]], err = strconv.Atoi(container[1])
				if err != nil {
					panic(err)
				}
			}
		}
		bags[thisBag.desc] = thisBag
	}

	fmt.Printf("found %d total bags\n", len(bags))
	canHold := bagCanHold(myBag, bags)
	total = len(canHold)

	// now get the flattened
	for desc, thisBag := range bags {
		//fmt.Printf("\tflattening %s\n", desc)
		canHold := bagCanHold(desc, bags)
		thisBag.flatCanHold = canHold
		bags[desc] = thisBag
		if count, ok := thisBag.flatCanHold[myBag]; ok && count > 0 {
			//fmt.Printf("%s bag can hold %s bag: (%#v)\n", desc, myBag, canHold)
			total++
		}
	}

	fmt.Printf("Part 1: Number of bags that can hold my %s bag: %d\n", myBag, total)
}

func bagCanHold(desc string, bags map[string]bag) map[string]int {
	//fmt.Printf("looking at %s bag\n", desc)
	thisBag, ok := bags[desc]
	if !ok {
		panic("can't find bag")
	}
	if thisBag.canHold == nil {
		//fmt.Printf("%s bag doesn't hold anything\n", desc)
		return nil
	}

	if thisBag.flatCanHold != nil {
		return thisBag.flatCanHold
	}

	canHold := make(map[string]int)
	for innerBag, thisCount := range thisBag.canHold {
		canHold[innerBag] = thisCount
	}
	//fmt.Printf("%s bag intially can hold %#v\n", desc, canHold)

	for innerBag, _ := range thisBag.canHold {
		bagCanHold := bagCanHold(innerBag, bags)

		for innerInnerBag, innerInnerBagCount := range bagCanHold {
			if existingCount, ok := canHold[innerInnerBag]; ok {
				canHold[innerInnerBag] = existingCount + innerInnerBagCount
			} else {
				canHold[innerInnerBag] = innerInnerBagCount
			}
		}
		//fmt.Printf("%s bag after looking at %s bag can hold %#v\n", desc, innerBag, canHold)
	}

	return canHold
}

func part2() {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	rLine := regexp.MustCompile(`^(\w+\s\w+) bags contain (no other bags.|.+)+$`)
	rBags := regexp.MustCompile(`(?:(\d+) (\w+\s\w+) bags?.\s?)`)

	bags := make(map[string]bag)
	for sc.Scan() {
		line := sc.Text()
		matches := rLine.FindAllStringSubmatch(line, -1)
		//fmt.Printf("matches is %#v\n", matches)
		thisBag := bag{
			desc: matches[0][1],
		}
		if matches[0][2] == "no other bags." {
			// leaf
		} else {
			bagsMatches := rBags.FindAllStringSubmatch(matches[0][2], -1)
			//fmt.Printf("bagsMatches is %#v\n", bagsMatches)

			thisBag.canHold = make(map[string]int)
			for _, container := range bagsMatches {
				thisBag.canHold[container[2]], err = strconv.Atoi(container[1])
				if err != nil {
					panic(err)
				}
			}
		}
		bags[thisBag.desc] = thisBag
	}

	fmt.Printf("found %d total bags\n", len(bags))
	bagCount := getBagCount(myBag, bags)

	//fmt.Printf("%s bag: %#v\n", myBag, bags[myBag])

	fmt.Printf("Part 2: A %s bag must contain %d bags\n", myBag, bagCount)
}

func getBagCount(bagName string, bags map[string]bag) int {
	//fmt.Printf("looking at %s bag\n", desc)
	thisBag, ok := bags[bagName]
	if !ok {
		panic("can't find bag")
	}
	if thisBag.canHold == nil {
		//fmt.Printf("%s bag doesn't hold anything\n", desc)
		return 0
	}

	if thisBag.bagCount != nil {
		return *thisBag.bagCount
	}

	bagCount := 0
	for innerBagName, count := range thisBag.canHold {
		innerBag := getBagCount(innerBagName, bags)

		bagCount += count + (count*innerBag)
	}

	thisBag.bagCount = &bagCount

	bags[bagName] = thisBag

	//fmt.Printf("%s bag holds %d bags\n", bagName, bagCount)

	return bagCount
}
