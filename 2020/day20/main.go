package main

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const testInput2Filename = "testinput2.txt"

const verbose = false

const (
	on  byte = '#'
	off byte = '.'
)

type tile struct {
	id        int
	grid      [][]byte
	edges     [4]string // top, right, bottom, left
	rotations int
	flipped   bool
}

func main() {
	input := lib.GetInputStrings(inputFilename)
	pic := part1(input)
	part2(pic)
}

func part1(input []string) [][]*tile {
	tiles := parseInput(input)

	if verbose {
		fmt.Printf("input: %s", spew.Sdump(tiles))
	}

	pic := arrangeTiles(tiles)

	picDimension := len(pic)
	result := pic[0][0].id * pic[0][picDimension-1].id * pic[picDimension-1][0].id * pic[picDimension-1][picDimension-1].id

	fmt.Printf("Part 1: %d\n", result)
	return pic
}

func part2(tiledPic [][]*tile) {
	pic := stitchPic(tiledPic)
	pic, monsterCount := countMonsters(pic)
	result := 0
	if monsterCount > 0 {
		result = countWaves(pic)
	}
	fmt.Printf("Part 2: %d\n", result)
}

func stitchPic(tiledPic [][]*tile) [][]byte {
	picSize := len(tiledPic)
	tileSize := len(tiledPic[0][0].grid) - 2
	pic := make([][]byte, tileSize*picSize)
	for tileI, row := range tiledPic {
		iOffset := tileI * tileSize
		for _, tile := range row {
			for i := 1; i < len(tile.grid)-1; i++ {
				for j := 1; j < len(tile.grid)-1; j++ {
					pic[i-1+iOffset] = append(pic[i-1+iOffset], tile.grid[i][j])
				}
			}
		}
	}
	fmt.Printf("Stitched pic:\n")
	printGrid(pic)
	return pic
}

func countMonsters(pic [][]byte) ([][]byte, int) {
	//01234567890123456789
	//.#...#.###...#.##.O#
	//O.##.OO#.#.OO.##.OOO
	//#O.#O#.O##O..O.#O##.
	monsterPairs := [][]int{
		{-1, -1},
		{0, -1},
		{0, -2},
		{1, -3},
		{1, -6},
		{0, -7},
		{0, -8},
		{1, -9},
		{1, -12},
		{0, -13},
		{0, -14},
		{1, -15},
		{1, -18},
		{0, -19},
	}

	size := len(pic)
	monsterCount := 0

	for flips := 0; flips < 2; flips++ {
		for rotations := 0; rotations < 4; rotations++ {
			fmt.Printf("trying with %d flips, %d rotations\n", flips, rotations)
			printGrid(pic)

			// find the nose of the monster
			for i := 1; i < size-1; i++ {
				for j := 19; j < size; j++ {
					if pic[i][j] == on {
						foundMonster := true
						// look for the body of the monster with relative pixels
						for _, coord := range monsterPairs {
							if pic[i+coord[0]][j+coord[1]] != on {
								foundMonster = false
								break
							}
						}
						if foundMonster {
							monsterCount++
							pic[i][j] = 'o'
							for _, coord := range monsterPairs {
								pic[i+coord[0]][j+coord[1]] = 'O'
							}
						}
					}
				}
			}
			if monsterCount > 0 {
				fmt.Printf("found %d monsters!\n", monsterCount)
				printGrid(pic)
				return pic, monsterCount
			}
			pic = rotateGrid(pic)
		}
		pic = flipGrid(pic)
	}

	return pic, monsterCount
}

func parseInput(input []string) map[int]tile {
	tiles := make(map[int]tile)
	var t tile
	for _, line := range input {
		if len(line) == 0 {
			t = calculateEdges(t)
			tiles[t.id] = t
		} else if strings.HasPrefix(line, "Tile") {
			var tileID int
			n, err := fmt.Sscanf(line, "Tile %d:", &tileID)
			if err != nil || n != 1 {
				panic(fmt.Errorf("unexpeced tile ID %v %s", err, line))
			}
			t = tile{
				id:   tileID,
				grid: make([][]byte, 0),
			}
		} else {
			row := make([]byte, 0)
			for _, p := range line {
				if string(p) == string(on) {
					row = append(row, on)
				} else {
					row = append(row, off)
				}
			}
			t.grid = append(t.grid, row)
		}
	}
	t = calculateEdges(t)
	tiles[t.id] = t

	if verbose {
		fmt.Printf("%d tiles\n", len(tiles))
	}
	return tiles
}

func calculateEdges(t tile) tile {
	newTile := t
	newTile.edges = [4]string{}
	tileSize := len(t.grid)
	maxOffset := tileSize - 1
	for offset := 0; offset < tileSize; offset++ {
		newTile.edges[0] = newTile.edges[0] + string(newTile.grid[0][offset])         // top edge
		newTile.edges[1] = newTile.edges[1] + string(newTile.grid[offset][maxOffset]) // right edge
		newTile.edges[2] = newTile.edges[2] + string(newTile.grid[maxOffset][offset]) // bottom edge
		newTile.edges[3] = newTile.edges[3] + string(newTile.grid[offset][0])         // left edge
	}
	return newTile
}

func arrangeTiles(tiles map[int]tile) [][]*tile {
	numTiles := len(tiles)
	picSize := int(math.Sqrt(float64(numTiles)))
	if (picSize * picSize) != numTiles {
		panic("not a square")
	}

	fmt.Printf("making a %d x %d pic from %d tiles\n", picSize, picSize, numTiles)

	fullPic := make([][]*tile, picSize)
	for i := 0; i < picSize; i++ {
		fullPic[i] = make([]*tile, picSize)
	}

	pic, err := fillPic(fullPic, tiles, []int{})
	if err != nil {
		panic(err)
	}

	fmt.Print("fully arranged tiles: \n")
	printPic(pic)
	return pic
}

func fillPic(startPic [][]*tile, tiles map[int]tile, startUsedTiles []int) ([][]*tile, error) {
	if verbose {
		fmt.Printf("filling pic, used tiles: %#v: \n", startUsedTiles)
		printPic(startPic)
	}
	pic := startPic
	picSize := len(pic)
	for i := 0; i < picSize; i++ {
		for j := 0; j < picSize; j++ {
			if pic[i][j] != nil {
				continue
			}
			for id, t := range tiles {
				if containsInt(startUsedTiles, id) {
					continue
				}
				for flips := 0; flips < 2; flips++ {
					flippedTile := flipTile(t, flips)
					for rotation := 0; rotation < 4; rotation++ {
						rotatedTile := rotateTile(flippedTile, rotation)

						if verbose {
							fmt.Printf("tile %d, flipped %d, rotated %d times has left edge %s\n", rotatedTile.id, flips, rotation, rotatedTile.edges[3])
						}
						if i != 0 {
							if rotatedTile.edges[0] != pic[i-1][j].edges[2] {
								continue
							}
						}
						if j != 0 {
							if rotatedTile.edges[3] != pic[i][j-1].edges[1] {
								continue
							}
						}
						if verbose {
							fmt.Printf("\ttile %d, flipped %d, rotated %d times is a fit in position [%d,%d]\n", rotatedTile.id, flips, rotation, i, j)
						}
						pic[i][j] = &rotatedTile
						usedTiles := append(startUsedTiles, t.id)
						if len(usedTiles) == len(tiles) {
							return pic, nil
						}
						pic, err := fillPic(pic, tiles, usedTiles)
						if err != nil {
							if verbose {
								fmt.Printf("\ttile %d, flipped %d, rotated %d times is not in position [%d,%d]\n", t.id, flips, rotation, i, j)
							}
							pic[i][j] = nil
							continue
						}
						if verbose {
							fmt.Printf("tile %d, flipped %d, rotated %d times is in position [%d][%d]\n", t.id, flips, rotation, i, j)
						}
						return pic, nil
					}
				}
			}
			return pic, errors.New("no tile found")
		}
	}
	return pic, nil
}

func rotateTile(t tile, rotations int) tile {
	newTile := t
	for r := 1; r <= rotations; r++ {
		newTile.grid = rotateGrid(newTile.grid)
		newTile = calculateEdges(newTile)
		newTile.rotations++
	}
	newTile.rotations = newTile.rotations % 4
	return newTile
}

func flipTile(t tile, n int) tile {
	newTile := t
	flips := n % 2
	switch flips {
	case 0:
		break
	case 1:
		newTile.grid = flipGrid(newTile.grid)
		newTile = calculateEdges(newTile)
		newTile.flipped = !newTile.flipped
	}
	return newTile
}

func rotateGrid(g [][]byte) [][]byte {
	size := len(g) // assumes grid is a square
	max := size - 1
	newGrid := make([][]byte, size)
	for i := 0; i < size; i++ {
		newGrid[i] = make([]byte, size)
		for j := 0; j < size; j++ {
			newGrid[i][j] = g[max-j][i]
		}
	}
	return newGrid
}

func flipGrid(g [][]byte) [][]byte {
	size := len(g) // assumes grid is a square
	max := size - 1
	newGrid := make([][]byte, size)
	for i := 0; i < size; i++ {
		newGrid[i] = make([]byte, size)
		for j := 0; j < size; j++ {
			newGrid[i][j] = g[i][max-j]
		}
	}
	return newGrid
}

func containsInt(ss []int, s int) bool {
	for _, n := range ss {
		if n == s {
			return true
		}
	}
	return false
}

func printPic(pic [][]*tile) {
	picSize := len(pic)
	for i := 0; i < picSize; i++ {
		for j := 0; j < picSize; j++ {
			var id int
			if pic[i][j] != nil {
				id = pic[i][j].id
			}
			fmt.Printf("%d ", id)
		}
		fmt.Print("\n")
	}
}

func (t tile) String() string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("Tile %d:\n", t.id))
	for _, row := range t.grid {
		str.Write(row)
		str.WriteString("\n")
	}
	str.WriteString(
		fmt.Sprintf(
			"   topEdge: %s\n rightEdge: %s\nbottomEdge: %s\n  leftEdge: %s\n",
			t.edges[0],
			t.edges[1],
			t.edges[2],
			t.edges[3],
		),
	)

	return str.String()
}

func printGrid(grid [][]byte) {
	for i, row := range grid {
		fmt.Printf("%2d ", i)
		for _, pixel := range row {
			fmt.Printf("%c", pixel)
		}
		fmt.Println()
	}
}

func countWaves(grid [][]byte) int {
	count := 0
	for _, row := range grid {
		for _, pixel := range row {
			if pixel == on {
				count++
			}
		}
	}
	return count
}
