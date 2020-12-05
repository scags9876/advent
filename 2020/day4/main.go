package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const inputFilename = "input.txt"

func main() {
	input := getInput()
	solvePuzzle(input)
}

func getInput() []map[string]string {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	var input []map[string]string
	currentPassport := make(map[string]string)
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			input = append(input, currentPassport)
			currentPassport = make(map[string]string)
			continue
		}
		for _, tuple := range strings.Split(line, " ") {
			fields := strings.Split(tuple, ":")
			currentPassport[fields[0]] = fields[1]
		}
	}
	input = append(input, currentPassport)
	return input
}

func solvePuzzle(input []map[string]string) {
	validPassportsCount := countValidPassports(input)
	fmt.Printf("Part 1: There are %d valid passports\n", validPassportsCount)

	validPassportsCount2 := countValidPassports2(input)
	fmt.Printf("Part 2: There are %d valid passports\n", validPassportsCount2)
}

func countValidPassports(input []map[string]string) int {
	count := 0
	requiredFields := []string{
		"byr", // (Birth Year)
		"iyr", // (Issue Year)
		"eyr", // (Expiration Year)
		"hgt", // (Height)
		"hcl", // (Hair Color)
		"ecl", // (Eye Color)
		"pid", // (Passport ID)
		//"cid", // (Country ID) optional
	}
	fmt.Printf("looking at %d passports\n", len(input))
	for _, passport := range input {
		valid := true
		for _, field := range requiredFields {
			if _, found := passport[field]; !found {
				valid = false
				fmt.Printf("passport %v is missing field %s!\n", passport, field)
				break
			}
		}
		if valid {
			fmt.Printf("passport %v is valid!\n", passport)
			count++
		}
	}
	return count
}

var hgtRegex = regexp.MustCompile(`^(\d+)(\w+)$`)
var hclRegex = regexp.MustCompile("^#[0-9a-f]{6}$")
var pidRegex = regexp.MustCompile(`^[1-9][0-9]{8}$`)

func countValidPassports2(input []map[string]string) int {
	count := 0
	requiredFields := []string{
		"byr", // (Birth Year)
		"iyr", // (Issue Year)
		"eyr", // (Expiration Year)
		"hgt", // (Height)
		"hcl", // (Hair Color)
		"ecl", // (Eye Color)
		"pid", // (Passport ID)
		//"cid", // (Country ID) optional
	}
	fmt.Printf("looking at %d passports\n", len(input))
	for _, passport := range input {
		valid := true
		for _, field := range requiredFields {
			if _, found := passport[field]; !found {
				valid = false
				//fmt.Printf("passport %v is missing field %s!\n", passport, field)
				break
			}
		}
		if valid == false {
			continue
		}
		//fmt.Printf("looking at passport %v\n", passport)

		passport["cid"] = "123"

		byr, _ := strconv.Atoi(passport["byr"])
		if byr < 1920 || byr > 2002 {
			valid = false
			//fmt.Printf("passport %v has invalid field %s!\n", passport, "byr")
			continue
		}
		iyr, _ := strconv.Atoi(passport["iyr"])
		if iyr < 2010 || iyr > 2020 {
			valid = false
			//fmt.Printf("passport %v has invalid field %s!\n", passport, "iyr")
			continue
		}
		eyr, _ := strconv.Atoi(passport["eyr"])
		if eyr < 2020 || eyr > 2030 {
			valid = false
			//fmt.Printf("passport %v has invalid field %s!\n", passport, "eyr")
			continue
		}
		var height int
		var unit string
		matches := hgtRegex.FindAllStringSubmatch(passport["hgt"], 1)
		if matches == nil {
			valid = false
			//fmt.Printf("passport %v has invalid field %s!\n", passport, "hgt")
			continue
		}
		height, _ = strconv.Atoi(matches[0][1])
		unit = matches[0][2]
		if unit == "cm" {
			if height < 150 || height > 193 {
				valid = false
				//fmt.Printf("passport %v has invalid field %s!\n", passport, "hgt")
				continue
			}
		} else if unit == "in" {
			if height < 59 || height > 76 {
				valid = false
				//fmt.Printf("passport %v has invalid field %s!\n", passport, "hgt")
				continue
			}
		} else {
			valid = false
			//fmt.Printf("passport %v has invalid field %s!\n", passport, "hgt")
			continue
		}
		if !hclRegex.MatchString(passport["hcl"]) {
			valid = false
			//fmt.Printf("passport %v has invalid field %s!\n", passport, "hcl")
			continue
		}
		ecl := passport["ecl"]
		if ecl != "amb" && ecl != "blu" && ecl != "brn" &&
			ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
			valid = false
			//fmt.Printf("passport %v has invalid field %s!\n", passport, "ecl")
			continue
		}
		pid, err := strconv.Atoi(passport["pid"])
		if err != nil || len(passport["pid"]) != 9 || pid <= 0 || pid > 999999999 {
			valid = false
			fmt.Printf("passport %v has invalid field %s!\n", passport, "pid")
			continue
		}

		//if !pidRegex.MatchString(passport["pid"]) {
		//	valid = false
		//	fmt.Printf("passport %v has invalid field %s!\n", passport, "pid")
		//	continue
		//}

		if valid {
			fmt.Printf("passport %v is valid!\n", passport)
			count++
		}
	}
	return count
}