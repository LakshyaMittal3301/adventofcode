package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func setBit(num *int, bit int) {
	*num |= (1 << bit)
}

// . -> 0
// # -> 1

// .#
// #.
// 0110

// num =  1010
// num >> 1 | ((num & 1) << 3)

// .#.
// #.#
// ##.

// 876
// 105
// 234

// num = 01010111'0'
// num >> 1

func encode2(grid [][]string) (encoding int) {
	if grid[0][0] == "#" {
		setBit(&encoding, 3)
	}
	if grid[0][1] == "#" {
		setBit(&encoding, 2)
	}
	if grid[1][1] == "#" {
		setBit(&encoding, 1)
	}
	if grid[1][0] == "#" {
		setBit(&encoding, 0)
	}
	return
}

func encode3(grid [][]string) (encoding int) {
	if grid[1][1] == "#" {
		setBit(&encoding, 8)
	}
	if grid[0][0] == "#" {
		setBit(&encoding, 7)
	}
	if grid[0][1] == "#" {
		setBit(&encoding, 6)
	}
	if grid[0][2] == "#" {
		setBit(&encoding, 5)
	}
	if grid[1][2] == "#" {
		setBit(&encoding, 4)
	}
	if grid[2][2] == "#" {
		setBit(&encoding, 3)
	}
	if grid[2][1] == "#" {
		setBit(&encoding, 2)
	}
	if grid[2][0] == "#" {
		setBit(&encoding, 1)
	}
	if grid[1][0] == "#" {
		setBit(&encoding, 0)
	}
	return
}

func makeSquareGrid(n int) [][]string {
	grid := make([][]string, n)
	for i := range grid {
		grid[i] = make([]string, n)
	}
	return grid
}

func inputToGrid(input string) [][]string {
	parts := strings.Split(input, "/")
	n := len(parts)

	grid := makeSquareGrid(n)
	for i := range n {
		for j := range n {
			grid[i][j] = string(parts[i][j])
		}
	}
	return grid
}

var outputMap2By2 map[int]([][]string)
var outputMap3By3 map[int]([][]string)

func readInputAndFillMaps() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get current file info")
	}

	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "input.txt")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		inputGrid := inputToGrid(fields[0])
		outputGrid := inputToGrid(fields[2])

		if len(inputGrid) == 2 {
			outputMap2By2[encode2(inputGrid)] = outputGrid
		} else {
			outputMap3By3[encode3(inputGrid)] = outputGrid
		}
	}
}

func breakInto(grid [][]string, size int) (pieces [][][]string) {
	n := len(grid)

	for i := range n / size {
		for j := range n / size {
			startRow := i * size
			startCol := j * size
			piece := makeSquareGrid(size)

			for row := range size {
				for col := range size {
					piece[row][col] = grid[row+startRow][col+startCol]
				}
			}
			pieces = append(pieces, piece)
		}
	}
	return
}

func copyGrid(grid [][]string) [][]string {
	newGrid := make([][]string, len(grid))
	for i := range grid {
		newGrid[i] = make([]string, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}

func flipVertically(grid [][]string) (flipped [][]string) {
	n := len(grid)
	flipped = copyGrid(grid)
	flipped[0], flipped[n-1] = flipped[n-1], flipped[0]
	return
}

func rotatePiece(grid [][]string) (rotated [][]string) {
	rotated = flipVertically(grid)
	n := len(rotated)
	for i := range n {
		for j := i + 1; j < n; j++ {
			rotated[i][j], rotated[j][i] = rotated[j][i], rotated[i][j]
		}
	}
	return
}

func getAllPieces(piece [][]string) (allPieces [][][]string) {
	for range 4 {
		allPieces = append(allPieces, piece)
		allPieces = append(allPieces, flipVertically(piece))
		piece = rotatePiece(piece)
	}
	return
}

func enhancePiece(piece [][]string, size int) (enhancedPiece [][]string) {
	allPieces := getAllPieces(piece)
	for _, newPiece := range allPieces {
		if size == 2 {
			enhancedPiece, ok := outputMap2By2[encode2(newPiece)]
			if ok {
				return enhancedPiece
			}
		} else {
			enhancedPiece, ok := outputMap3By3[encode3(newPiece)]
			if ok {
				return enhancedPiece
			}
		}
	}
	panic("Could not enhance piece")
}

func enhancePieces(pieces [][][]string, size int) (enhancedPieces [][][]string) {
	for _, piece := range pieces {
		enhancedPieces = append(enhancedPieces, enhancePiece(piece, size))
	}
	return
}

func mergePieces(pieces [][][]string) [][]string {
	totalPieces := len(pieces)
	sizeOfPiece := len(pieces[0])
	piecesInOneRow := int(math.Sqrt(float64(totalPieces)))
	sizeOfGrid := sizeOfPiece * piecesInOneRow

	grid := makeSquareGrid(sizeOfGrid)

	for idx, piece := range pieces {
		row := (idx / piecesInOneRow) * sizeOfPiece
		col := (idx % piecesInOneRow) * sizeOfPiece
		for i := range sizeOfPiece {
			for j := 0; j < sizeOfPiece; j++ {
				grid[row+i][col+j] = piece[i][j]
			}
		}
	}
	return grid
}

func enhance(grid [][]string) (finalGrid [][]string) {
	var size int

	n := len(grid)
	if n%2 == 0 {
		size = 2
	} else {
		size = 3
	}

	pieces := breakInto(grid, size)
	enhancedPieces := enhancePieces(pieces, size)
	finalGrid = mergePieces(enhancedPieces)
	return
}

func main() {
	outputMap2By2 = make(map[int]([][]string))
	outputMap3By3 = make(map[int]([][]string))

	readInputAndFillMaps()

	grid := [][]string{
		{".", "#", "."},
		{".", ".", "#"},
		{"#", "#", "#"},
	}

	for range 5 {
		grid = enhance(grid)
	}

	ans := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == "#" {
				ans++
			}
		}
	}

	fmt.Println("Answer: ", ans)
}
