package main

import (
	"bufio"
	"fmt"
	"os"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"

const maxCycles = 6
const verbose = false

func main() {
	input := getInput(inputFilename)
	part1(input)
	part2(input)
}

const (
	active   rune = '#'
	inactive      = '.'
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
	cubeState [][][][]rune
	zOffset   int
	wSize     int
}

func part1(input [][]rune) {
	var activeCount int
	pocket, activeCubes := createPocketDimension(input)
	fmt.Printf("Initial state, %d active cubes:\n", activeCubes)
	printPocket(pocket, false)
	for round := 0; round < maxCycles; round++ {
		pocket, activeCount = flipCubes(pocket, false)

		fmt.Printf("After round %d, %d active cubes:\n", round+1, activeCount)
		printPocket(pocket, false)
	}
	fmt.Printf("Part 1: After %d rounds, %d active cubes\n", maxCycles, activeCount)
}

func part2(input [][]rune) {
	var activeCount int
	pocket, activeCubes := createPocketDimension(input)
	fmt.Printf("Initial state, %d active cubes:\n", activeCubes)
	printPocket(pocket, true)
	for round := 0; round < maxCycles; round++ {
		pocket, activeCount = flipCubes(pocket, true)

		fmt.Printf("After round %d, %d active cubes:\n", round+1, activeCount)
		printPocket(pocket, true)
	}
	fmt.Printf("Part 2: After %d rounds, %d active cubes\n", maxCycles, activeCount)
}

func createPocketDimension(input [][]rune) (pocket, int) {
	cubeState := make([][][][]rune, 1)
	activeCubes := 0

	cubeState[0] = make([][][]rune, len(input))

	for y, col := range input {
		cubeState[0][y] = make([][]rune, len(col))
		for x, state := range col {
			cubeState[0][y][x] = []rune{state}
			if state == active {
				activeCubes++
			}
		}
	}

	wSize := 1
	pocket := pocket{
		cubeState: cubeState,
		zOffset:   0,
		wSize:     wSize,
	}

	return pocket, activeCubes
}

func printPocket(p pocket, use4D bool) {
	if verbose != true {
		return
	}
	for w := 0; w < p.wSize; w++ {
		for z, plane := range p.cubeState {
			fmt.Printf("z=%d", z-p.zOffset)
			if use4D {
				fmt.Printf(", w=%d", w-p.zOffset)
			}
			fmt.Println(":")
			for _, y := range plane {
				for _, x := range y {
					state := x[w]
					fmt.Printf("%c", state)
				}
				fmt.Print("\n")
			}
			fmt.Print("\n")
		}
	}
}

func flipCubes(p pocket, use4D bool) (pocket, int) {
	activeCubes := 0

	cubeState := make([][][][]rune, 0)

	minW := -1
	maxW := p.wSize
	if !use4D {
		minW = 0
		maxW = 0
	}
	for z := -1; z <= len(p.cubeState); z++ {
		newPlane := make([][][]rune, 0)
		for y := -1; y <= len(p.cubeState[0]); y++ {
			newRow := make([][]rune, 0)
			for x := -1; x <= len(p.cubeState[0][0]); x++ {
				newPoint := make([]rune, 0)
				for w := minW; w <= maxW; w++ {

					newState := inactive
					if z >= 0 && z < len(p.cubeState) &&
						y >= 0 && y < len(p.cubeState[z]) &&
						x >= 0 && x < len(p.cubeState[z][y]) &&
						w >= 0 && w < len(p.cubeState[z][y][x]) {
						newState = p.cubeState[z][y][x][w]
					}

					adjacentCount := adjacentCount(p, w, x, y, z, use4D)
					if newState == active {
						if adjacentCount != 2 && adjacentCount != 3 {
							newState = inactive
						}
					} else {
						if adjacentCount == 3 {
							newState = active
						}
					}
					newPoint = append(newPoint, newState)

					if newState == active {
						activeCubes++
					}
				}
				newRow = append(newRow, newPoint)
			}
			newPlane = append(newPlane, newRow)
		}
		cubeState = append(cubeState, newPlane)
	}

	wSize := p.wSize
	if use4D {
		wSize += 2
	}

	newPocket := pocket{
		cubeState: cubeState,
		zOffset:   p.zOffset + 1,
		wSize:     wSize,
	}
	return newPocket, activeCubes
}

func adjacentCount(p pocket, w, x, y, z int, use4D bool) int {
	count := 0
	zSize := len(p.cubeState)
	ySize := len(p.cubeState[0])
	xSize := len(p.cubeState[0][0])

	minW := -1
	maxW := 1
	if !use4D {
		minW = 0
		maxW = 0
	}

	for zOffset := -1; zOffset <= 1; zOffset++ {
		zCheck := z + zOffset

		if zCheck < 0 || zCheck >= zSize {
			continue
		}
		for yOffset := -1; yOffset <= 1; yOffset++ {
			yCheck := y + yOffset

			if yCheck < 0 || yCheck >= ySize {
				continue
			}
			for xOffset := -1; xOffset <= 1; xOffset++ {
				xCheck := x + xOffset

				if xCheck < 0 || xCheck >= xSize {
					continue
				}
				for wOffset := minW; wOffset <= maxW; wOffset++ {
					wCheck := w + wOffset
					wSize := len(p.cubeState[zCheck][yCheck][xCheck])

					if wCheck < 0 || wCheck >= wSize {
						continue
					}

					// skip if trying to check the point that we are inquiring about
					if zCheck == z && yCheck == y && xCheck == x && wCheck == w {
						continue
					}
					if p.cubeState[zCheck][yCheck][xCheck][wCheck] == active {
						count++
					}
				}
			}
		}
	}
	return count
}
