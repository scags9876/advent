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

type instruction struct {
	op string
	value int
	visited bool
	changed bool
}
func part1() {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	instructions := make([]instruction, 0)
	for sc.Scan() {
		line := sc.Text()

		var operation string
		var value int
		_, err := fmt.Sscanf(line, "%s %d", &operation, &value)
		if err != nil {
			panic(err)
		}

		instr := instruction{
			op:    operation,
			value: value,
		}
		instructions = append(instructions, instr)
	}
	//fmt.Printf("instructions: %+v", instructions)

	accumulator := 0
	position := 0
	for {
		instr := instructions[position]
		//fmt.Printf("insruction %3d: %v (accumulator: %d)\n", position, instr, accumulator)
		if instr.visited {
			break
		}
		instructions[position].visited = true
		switch instr.op {
		case "nop":
			position++
		case "acc":
			accumulator += instr.value
			position++
		case "jmp":
			position += instr.value
		default:
			panic(fmt.Errorf("unknown operation: %s", instr.op))
		}
	}

	fmt.Printf("Part 1: loop starts at instruction %d with accumulator value: %d\n", position, accumulator)
}

func part2() {
	file, err := os.Open(inputFilename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	instructions := make([]instruction, 0)
	for sc.Scan() {
		line := sc.Text()

		var operation string
		var value int
		_, err := fmt.Sscanf(line, "%s %d", &operation, &value)
		if err != nil {
			panic(err)
		}

		instr := instruction{
			op:    operation,
			value: value,
		}
		instructions = append(instructions, instr)
	}
	//fmt.Printf("instructions: %+v", instructions)

	accumulator := 0
	donePosition := len(instructions)
	changeInstruction := donePosition - 1
	var done bool
	for done != true {
		var changeFromOp, changeToOp string
		for i := changeInstruction; i > 0; i-- {
			instr := instructions[i]
			if !instr.changed && instr.op != "acc" {
				instr.changed = true
				if instr.op == "jmp" {
					changeFromOp = "jmp"
					changeToOp = "nop"
				} else if instr.op == "nop" {
					changeFromOp = "nop"
					changeToOp = "jmp"
				}
				instr.op = changeToOp
				instructions[i] = instr
				changeInstruction = i
				fmt.Printf("Changing instruction %3d from %s to %s: %+v\n", changeInstruction, changeFromOp, changeToOp, instr)
				break
			}
		}

		accumulator = 0
		position := 0
		for {
			instr := instructions[position]
			//fmt.Printf("insruction %3d: %v (accumulator: %d)\n", position, instr, accumulator)
			if instr.visited {
				break
			}
			instructions[position].visited = true
			switch instr.op {
			case "nop":
				position++
			case "acc":
				accumulator += instr.value
				position++
			case "jmp":
				position += instr.value
			default:
				panic(fmt.Errorf("unknown operation: %s", instr.op))
			}

			if position >= donePosition {
				done = true
				break
			}
		}

		if !done {
			instr := instructions[changeInstruction]
			instr.op = changeFromOp
			instructions[changeInstruction] = instr
			for i := range instructions {
				instructions[i].visited = false
			}
		}
	}
	fmt.Printf("Part 2: Fixed the program!  Accumulator is %d\n", accumulator)
}
