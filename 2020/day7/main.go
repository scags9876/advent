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
	desc         string
	contents     map[string]int
	flatContents map[string]int
	bagCount     *int
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

			thisBag.contents = make(map[string]int)
			for _, innerBag := range bagsMatches {
				thisBag.contents[innerBag[2]], err = strconv.Atoi(innerBag[1])
				if err != nil {
					panic(err)
				}
			}
		}
		bags[thisBag.desc] = thisBag
	}

	fmt.Printf("found %d total bags\n", len(bags))

	contents := bagContents(myBag, bags)
	total = len(contents)

	// now get the flattened contents
	for desc, _ := range bags {
		//fmt.Printf("\t flattening %s\n", desc)
		contents := bagContents(desc, bags)
		if count, ok := contents[myBag]; ok && count > 0 {
			//fmt.Printf("%s bag can hold %s bag: (%#v)\n", desc, myBag, contents)
			total++
		}
	}

	fmt.Printf("Part 1: Number of bags that can hold my %s bag: %d\n", myBag, total)
}

func bagContents(desc string, bags map[string]bag) map[string]int {
	//fmt.Printf("looking at %s bag\n", desc)
	thisBag, ok := bags[desc]
	if !ok {
		panic("can't find bag")
	}
	if thisBag.contents == nil {
		//fmt.Printf("%s bag doesn't hold anything\n", desc)
		return nil
	}

	if thisBag.flatContents != nil {
		return thisBag.flatContents
	}

	contents := make(map[string]int)
	for innerBag, thisCount := range thisBag.contents {
		contents[innerBag] = thisCount
	}
	//fmt.Printf("%s bag initially can hold %#v\n", desc, contents)

	for innerBag := range thisBag.contents {
		innerBagContents := bagContents(innerBag, bags)

		for innerInnerBag, innerInnerBagCount := range innerBagContents {
			if existingCount, ok := contents[innerInnerBag]; ok {
				contents[innerInnerBag] = existingCount + innerInnerBagCount
			} else {
				contents[innerInnerBag] = innerInnerBagCount
			}
		}
		//fmt.Printf("%s bag after looking at %s bag can hold %#v\n", desc, innerBag, contents)
	}

	thisBag.flatContents = contents
	bags[desc] = thisBag

	return contents
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

			thisBag.contents = make(map[string]int)
			for _, container := range bagsMatches {
				thisBag.contents[container[2]], err = strconv.Atoi(container[1])
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
	if thisBag.contents == nil {
		//fmt.Printf("%s bag doesn't hold anything\n", desc)
		return 0
	}

	if thisBag.bagCount != nil {
		return *thisBag.bagCount
	}

	bagCount := 0
	for innerBagName, count := range thisBag.contents {
		innerBag := getBagCount(innerBagName, bags)

		bagCount += count + (count*innerBag)
	}

	thisBag.bagCount = &bagCount

	bags[bagName] = thisBag

	//fmt.Printf("%s bag holds %d bags\n", bagName, bagCount)

	return bagCount
}
