package main

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/scags9876/adventOfCode/lib"
	"math"
	"strings"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const testInput2Filename = "testinput2.txt"

const verbose = false

const (
	on byte = '#'
	off byte = '.'
)

type tile struct {
	id int
	grid [][]byte
	edges [4]string // top, right, bottom, left
	rotations int
	flipped bool
}

func main() {
	input := lib.GetInputStrings(inputFilename)
	pic := part1(input)
	fmt.Printf("pic size: %d\n", len(pic))
	//part2(input)
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
				id:         tileID,
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
	maxOffset := tileSize-1
	for offset := 0; offset < tileSize; offset++ {
		newTile.edges[0] = newTile.edges[0] + string(newTile.grid[0][offset]) // top edge
		newTile.edges[1] = newTile.edges[1] + string(newTile.grid[offset][maxOffset]) // right edge
		newTile.edges[2] = newTile.edges[2] + string(newTile.grid[maxOffset][offset]) // bottom edge
		newTile.edges[3] = newTile.edges[3] + string(newTile.grid[offset][0]) // left edge
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
	for r := 1; r<=rotations; r++ {
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
	max := size-1
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
	max := size-1
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
