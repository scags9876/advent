package main

import (
	"bufio"
	"fmt"
	"os"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"

func main() {
	input := getInput(inputFilename)
	part1(input)
	//part2(input)
}

const (
	active rune = '#'
	inactive = '.'
)

func getInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("%s file not found", inputFilename))
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	input := make([][]rune, 0)

	for sc.Scan() {
		line := sc.Text()
		y := make([]rune, len(line))
		for x, val := range line {
			y[x] = val
		}

		input = append(input, y)
	}

	return input
}

type pocket struct {
	cubeState [][][]rune
	zOffset int
}

func part1(input [][]rune) {
	var activeCount int
	pocket, activeCubes := createPocketDimension(input)
	fmt.Printf("Initial state, %d active cubes:\n", activeCubes)
	printPocket(pocket)
	for round := 0; round < 6; round++ {
		pocket, activeCount = flipCubes(pocket)

		fmt.Printf("After round %d, %d active cubes:\n", round+1, activeCount)
		printPocket(pocket)
	}
	fmt.Printf("Part 1: After 6 rounds, %d active cubes\n", activeCount)
}

func createPocketDimension(input [][]rune) (pocket, int) {
	cubeState := make([][][]rune, 1)
	activeCubes := 0

	cubeState[0] = make([][]rune, len(input))

	for y, col := range input {
		cubeState[0][y] = make([]rune, len(col))
		for x, state := range col {
			cubeState[0][y][x] = state
			if state == active {
				activeCubes++
			}
		}
	}

	pocket := pocket{
		cubeState: cubeState,
		zOffset: 0,
	}

	return pocket, activeCubes
}

func printPocket(pocket pocket) {
	for z, plane := range pocket.cubeState {
		fmt.Printf("z=%d:\n", z-pocket.zOffset)
		for _, y := range plane {
			for _, state := range y {
				fmt.Printf("%c", state)
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}


func flipCubes(p pocket) (pocket, int) {
	activeCubes := 0

	cubeState := make([][][]rune, 0)

	for z := -1; z <= len(p.cubeState); z++ {
		newPlane := make([][]rune, 0)
		for y := -1; y <= len(p.cubeState[0]); y++ {
			newRow := make([]rune, 0)
			for x := -1; x <= len(p.cubeState[0][0]); x++ {
				newState := inactive
				if z >= 0 && z < len(p.cubeState) && y >= 0 && y < len(p.cubeState[z]) && x >=0 && x < len(p.cubeState[z][y]) {
					newState = p.cubeState[z][y][x]
				}
				
				adjacentCount := adjacentCount(p, x, y, z)
				if newState == active {
					if adjacentCount != 2 && adjacentCount != 3 {
						newState = inactive
					}
				} else {
					if adjacentCount == 3 {
						newState = active
					}
				}
				
				newRow = append(newRow, newState)

				if newState == active {
					activeCubes++
				}
			}
			newPlane = append(newPlane, newRow)
		}
		cubeState = append(cubeState, newPlane)
	}
	
	newPocket := pocket{
		cubeState: cubeState,
		zOffset: p.zOffset+1,
	}
	return newPocket, activeCubes
}

func adjacentCount(p pocket, x, y, z int) int {
	
  count := 0
	for vector := 0; vector < 26; vector++ {
		xCheck, yCheck, zCheck := x, y, z
		switch vector {
		case 0, 1, 2, 3, 4, 5, 6, 7, 8: // up
			zCheck = z+1
		case 17, 18, 19, 20, 21, 22, 23, 24, 25: // down
			zCheck = z-1
		}
		switch vector {
		case 0, 1, 2, 9, 10, 11, 17, 18, 19: // behind
			yCheck = y-1
		case 6, 7, 8, 14, 15, 16, 23, 24, 25: // in front of
			yCheck = y+1
		}
		switch vector {
		case 0, 3, 6, 9, 12, 14, 17, 20, 23: // to the left
			xCheck = x-1
		case 2, 5, 8, 11, 13, 16, 19, 22, 25: // to the right
			xCheck = x+1
		}

		//fmt.Printf("[%d,%d],[%d,%d], offset %d, vector %d, check [%d,%d]\n", ySize,xSize, i,j,offset,vector,yCheck,xCheck)
		zSize := len(p.cubeState)
		ySize := len(p.cubeState[0])
		xSize := len(p.cubeState[0][0])
		if zCheck < 0 || zCheck >= zSize || yCheck < 0 || yCheck >= ySize || xCheck < 0 || xCheck >= xSize {
			continue
		} else if p.cubeState[zCheck][yCheck][xCheck] == active {
			count++
		}
	}
	return count
}

