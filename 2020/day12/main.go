package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"

func main() {
	input := lib.GetInputStrings(inputFilename)
	part1(input)
	part2(input)
}

const (
	East  = 0
	South = 1
	West  = 2
	North = 3
)

var directions = []string{
	"E",
	"S",
	"W",
	"N",
}

func part1(input []string) {
	currentX := 0
	currentY := 0
	currentDirection := East
	for i, instr := range input {
		operation := string(instr[0])
		value, err := strconv.Atoi(instr[1:])
		if err != nil {
			panic(err)
		}

		if operation == "F" {
			operation = directions[currentDirection]
		}

		switch operation {
		case "N": // Action N means to move north by the given value.
			currentY += value
		case "S": // Action S means to move south by the given value.
			currentY -= value
		case "E": // Action E means to move east by the given value.
			currentX += value
		case "W": // Action W means to move west by the given value.
			currentX -= value
		case "L": // Action L means to turn left the given number of degrees.
			switch value {
			case 90:
				currentDirection--
			case 180:
				currentDirection -= 2
			case 270:
				currentDirection -= 3
			default:
				panic(fmt.Errorf("unexpected turn: %s", instr))
			}
			if currentDirection < 0 {
				currentDirection = 4 + currentDirection
			}
		case "R": // Action R means to turn left the given number of degrees.
			switch value {
			case 90:
				currentDirection++
			case 180:
				currentDirection += 2
			case 270:
				currentDirection += 3
			default:
				panic(fmt.Errorf("unexpected turn: %s", instr))
			}
			if currentDirection > 3 {
				currentDirection = currentDirection - 4
			}
		}

		fmt.Printf("%3d. After %4s at position [%3d,%3d], facing %s\n", i, instr, currentX, currentY, directions[currentDirection])
	}

	result := math.Abs(float64(currentX)) + math.Abs(float64(currentY))

	fmt.Printf("Part 1: %d\n", int(result))
}

type coord struct {
	x, y int
}

func part2(input []string) {
	var ship, waypoint coord

	waypoint.x = 10
	waypoint.y = 1

	for i, instr := range input {
		operation := string(instr[0])
		value, err := strconv.Atoi(instr[1:])
		if err != nil {
			panic(err)
		}

		//#X#|###
		//###|###
		//###|###
		//---+---
		//###|###
		//x##|###
		//###|#X#
		//3,2 -> 2, -3
		//2, -3 -> -3, -2
		//2, -3
		switch operation {
		case "N": // Action N means to move the waypoint north by the given value.
			waypoint.y += value
		case "S": // Action S means to move the waypoint south by the given value.
			waypoint.y -= value
		case "E": // Action E means to move the waypoint east by the given value.
			waypoint.x += value
		case "W": // Action W means to move the waypoint west by the given value.
			waypoint.x -= value
		case "F": // Action F means to move forward to the waypoint a number of times equal to the given value.
			ship.x = ship.x + (waypoint.x * value)
			ship.y = ship.y + (waypoint.y * value)
		case "L", "R":
			// Action L means to rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
			// Action R means to rotate the waypoint around the ship right (clockwise) the given number of degrees.
			turns := value / 90
			xTransform, yTransform := 1, 1
			if operation == "L" {
				yTransform = -1
			} else {
				xTransform = -1
			}
			for i := 1; i <= turns; i++ {
				newX := waypoint.y * yTransform
				newY := waypoint.x * xTransform
				waypoint.x = newX
				waypoint.y = newY
			}
		}

		fmt.Printf("%3d. After %4s at position [%3d,%3d], waypoint at [%3d,%3d]\n", i, instr, ship.x, ship.y, waypoint.x, waypoint.y)
	}

	result := math.Abs(float64(ship.x)) + math.Abs(float64(ship.y))

	fmt.Printf("Part 2: %d\n", int(result))
}
