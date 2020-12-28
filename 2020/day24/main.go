package main

import (
	"fmt"
	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"

const verbose = false

const maxDays = 100

func main() {
	input := lib.GetInputStrings(inputFilename)
	floor := part1(input)
	part2(floor)
}

func part1(input []string) map[string]bool {
	// the floor is laid out in a hex grid, represented by 3 coordinate system.  ref: https://www.redblobgames.com/grids/hexagons/
	floor := make(map[string]bool)

	for _, line := range input {
		size := len(line)
		i := 0
		x, y, z := 0, 0, 0
		//fmt.Printf("reading instrustions %s\n", line)

		for {
			directionFound := false
			var dir string
			if i < size-1 {
				dir = line[i:i+2]
				switch dir {
				case "ne":
					x++
					z--
					directionFound = true
				case "nw":
					y++
					z--
					directionFound = true
				case "se":
					y--
					z++
					directionFound = true
				case "sw":
					x--
					z++
					directionFound = true
				}
			}
			if directionFound {
				i += 2
			} else {
				dir = string(line[i])
				switch dir {
				case "w":
					x--
					y++
				case "e":
					x++
					y--
				default:
					panic(fmt.Sprintf("unknown direction: %s ",dir))
				}
				i++
			}

			//fmt.Printf("direction %s results in %d,%d,%d (next i: %d)\n", dir, x,y,z, i)
			if i == len(line) {
				break
			}
		}
		coord := lib.JoinInts([]int{x,y,z}, ",")
		if floor[coord] {
			if verbose {
				fmt.Printf("Flipping tile %s to white (%s)\n", coord, line)
			}
			delete(floor, coord)
		} else {
			if verbose {
				fmt.Printf("Flipping tile %s to black (%s)\n", coord, line)
			}
			floor[coord] = true
		}
	}

	count := countBlackTiles(floor)

	fmt.Printf("Part 1: %d black tiles remain\n", count)
	return floor
}

func countBlackTiles(floor map[string]bool) int {
	count := 0
	for coord, value := range floor {
		if value {
			count++
			if verbose {
				fmt.Printf("Tile %s is black\n", coord)
			}
		} else if verbose {
			fmt.Printf("Tile %s is white\n", coord)
		}
	}
	return count
}

var cubeDirections = [][3]int{
	{+1, -1, 0},
	{+1, 0, -1},
	{0, +1, -1},
	{-1, +1, 0},
	{-1, 0, +1},
	{0, -1, +1},
}

func part2(initialFloor map[string]bool) {

	blackCount := countBlackTiles(initialFloor)
	fmt.Printf("\nStarting part 2 with %d black tiles\n", blackCount)
	floor := initialFloor

	for day := 1; day <= maxDays; day++ {
		newFloor := make(map[string]bool)
		minX,minY,_,maxX,maxY,_ := getFloorBounds(floor)
		for x := minX-1; x <= maxX+1; x++ {
			for y := minY-1; y <= maxY+1; y++ {
				// always with cube coordinates the sum of x,y,z must equal zero
				z := 0-(x+y)
				coord := lib.JoinInts([]int{x,y,z}, ",")
				count := 0
				for _, direction := range cubeDirections {
					adjacentX := x+direction[0]
					adjacentY := y+direction[1]
					adjacentZ := z+direction[2]
					adjacentCoord := lib.JoinInts([]int{adjacentX,adjacentY,adjacentZ}, ",")
					if adjacentValue, adjacentOK := floor[adjacentCoord]; adjacentOK && adjacentValue {
						count++
					}
				}
				if value, ok := floor[coord]; ok && value {
					// tile is black, it flips to white with zero or more than 2 adjacent black tiles,
					// otherwise keep it black
					if !(count == 0 || count > 2) {
						newFloor[coord] = true
					}
				} else {
					// tile is white, it flips to black with exactly 2 black tiles adjacent,
					// otherwise keep it while
					if count == 2 {
						newFloor[coord] = true
					}
				}
			}
		}
		floor = newFloor
		if verbose {
			blackCount := countBlackTiles(floor)
			fmt.Printf("Day %3d: %d\n", day, blackCount)
		}
	}

	blackCount = countBlackTiles(floor)
	fmt.Printf("\nEnding part 2 with %d black tiles\n", blackCount)

}

func getFloorBounds(floor map[string]bool) (minX,minY,minZ,maxX,maxY,maxZ int) {
	for coord, value := range floor {
		if !value {
			continue
		}
		var x, y, z int
		_, err := fmt.Sscanf(coord, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			panic(err)
		}
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
		if z < minZ {
			minZ = z
		}
		if z > maxZ {
			maxZ = z
		}
	}
	return
}
