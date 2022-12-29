package main

import (
	"fmt"
	"strings"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"
const expectedTestResultPart1 = 95437
const expectedTestResultPart2 = 24933642

const dirMaxSize = 100000
const totalDiskSize = 70000000
const neededSize = 30000000

const verbose = false

type file struct {
	name string
	size int
}

type dir struct {
	name     string
	parent   *dir
	children []*dir
	files    []file
	size     int
}

func main() {
	testInput := lib.GetInputStrings(testInputFilename)
	input := lib.GetInputStrings(inputFilename)

	testDirList := parsefilesystem(testInput)

	fmt.Print("Part 1:\n")
	testResult := part1(testDirList)
	if verbose {
		fmt.Printf("Part 1 testInput result: %d\n", testResult)
	}
	if testResult == expectedTestResultPart1 {
		fmt.Printf("testfile returned expected result! (%d)\n", testResult)
	} else {
		panic(fmt.Errorf("expected %d but got %d", expectedTestResultPart1, testResult))
	}

	dirList := parsefilesystem(input)

	result := part1(dirList)
	fmt.Printf("Part 1 result: %d\n", result)

	fmt.Print("\nPart 2: \n")
	testResult = part2(testDirList)
	if verbose {
		fmt.Printf("Part 2 testInput result: %d\n", testResult)
	}
	if testResult == expectedTestResultPart2 {
		fmt.Printf("testfile returned expected result! (%d)\n", testResult)
	} else {
		panic(fmt.Errorf("expected %d but got %d", expectedTestResultPart2, testResult))
	}

	result = part2(dirList)
	fmt.Printf("Part 2 result: %d\n", result)
}

func parsefilesystem(input []string) []*dir {
	rootDir := dir{
		name:     "/",
		children: make([]*dir, 0),
		files:    make([]file, 0),
	}
	listingMode := false
	currentDir := &rootDir

	dirList := []*dir{&rootDir}

	for i, line := range input {
		if verbose {
			fmt.Printf("line %2d: %s\n", i, line)
		}
		if strings.HasPrefix(line, "$") {
			if listingMode {
				listingMode = false
			}
			if line == "$ ls" {
				listingMode = true
				continue
			}
			var newDirName string
			_, err := fmt.Sscanf(line, "$ cd %s", &newDirName)
			if err != nil {
				panic(err)
			}
			if newDirName == ".." {
				if currentDir.parent == nil {
					panic("root dir does not have parent directory")
				}
				currentDir = currentDir.parent
			} else if newDirName == "/" {
				// skip the root dir
			} else {
				i := findDirIdx(currentDir.children, newDirName)
				currentDir = currentDir.children[i]
			}
		} else {
			if !listingMode {
				continue
			}
			var newDirName string
			_, err := fmt.Sscanf(line, "dir %s", &newDirName)
			if err == nil {
				newDir := &dir{
					name:     newDirName,
					parent:   currentDir,
					children: make([]*dir, 0),
					files:    make([]file, 0),
				}
				currentDir.children = append(currentDir.children, newDir)
				dirList = append(dirList, newDir)
			} else {
				var newFileSize int
				var newFilename string
				_, err = fmt.Sscanf(line, "%d %s", &newFileSize, &newFilename)

				file := file{
					name: newFilename,
					size: newFileSize,
				}
				currentDir.files = append(currentDir.files, file)
				currentDir.size += newFileSize
			}
		}
	}
	return dirList
}

func findDirIdx(list []*dir, name string) int {
	for i, directory := range list {
		if directory.name == name {
			return i
		}
	}
	return -1
}

func (d *dir) totalSize() int {
	sz := d.size
	for _, childDir := range d.children {
		sz += childDir.totalSize()
	}
	return sz
}

func part1(dirList []*dir) int {
	var result int

	for _, dir := range dirList {
		sz := dir.totalSize()
		if verbose {
			fmt.Printf("directory %s has size %d\n", dir.name, sz)
		}
		if sz < dirMaxSize {
			result += sz
		}
	}

	return result
}

func part2(dirList []*dir) int {
	rootDirIdx := findDirIdx(dirList, "/")
	rootDirSize := dirList[rootDirIdx].totalSize()

	currentFreeSpace := totalDiskSize - rootDirSize
	necessarySpace := neededSize - currentFreeSpace

	if verbose {
		fmt.Printf("current free space: %d\n", currentFreeSpace)
		fmt.Printf("to get to %d, must free at least %d\n", neededSize, necessarySpace)
	}
	var result int
	for _, dir := range dirList {
		sz := dir.totalSize()
		if verbose {
			fmt.Printf("directory %s has size %d\n", dir.name, sz)
		}
		if sz >= necessarySpace {
			if result == 0 || sz < result {
				result = sz
			}
		}
	}

	return result
}
