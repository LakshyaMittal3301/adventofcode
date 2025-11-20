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

// 3 -> 4 -> 6 -> 9 -> 12 -> 18

// Divide a grid into 2x2 or 3x3 subgrids
// For each subgrid, enhance it and return either the 3x3 or 4x4 subgrid
// Merge the subgrid into one large grid
// Do this 5 times

/*
2x2

..
.. => .../.../..# => 0000
#... => #.#/..#/... => 1000
##
.. => #.#/..#/#.# => 1100
#.
#. => 1001
.#
.# => 0110
.#/#. => #../.../.## => 0110
##/#. => ###/#.#/..# => 1110
##/## => #.#/.../#.. => 1111

Rotation for 2x2: rotate the bits by 1
Flip horizontal: 3<->2, 1<->0
Flip vertical: 3<->0, 2<->1

----------------------------------------
3x3

##.
...
... => ##.#/..#./####/...# => 011000000
#.#/.../... => ##.#/##../#.#./.#..
###.../... => #..#/#..#/##../##.#
.#./#../... => #.##/##../.#.#/..##

#.#
...
#.. => ...#/##.#/#.#./#... => 010100010

#..
...
#.# => 010001010

#.#
...
..# => 010101000

765
084
123
Rotation for 3x3: rotate the last 8 bits by 2
Flip horizontally: 7<->5, 4<->0, 3<->1
Flip vertically: 7<->1, 6<->2, 5<->3

1100 -> 12
##/#.
1101 -> 13
*/

func setBit(num *int, bit int) {
	*num |= (1 << bit)
}

func isSet(num *int, bit int) bool {
	return ((*num) >> bit) & 1 == 1
}

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

func decode2(encoding int) (grid [][]string) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			grid[i][j] = "."
		}
	}

	if isSet(&encoding, 3) {
		grid[0][0] = "#"
	} 
	if isSet(&encoding, 2) {
		grid[0][1] = "#"
	} 
	if isSet(&encoding, 1) {
		grid[1][1] = "#"
	} 
	if isSet(&encoding, 0) {
		grid[1][0] = "#"
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

func decode3(encoding int) (grid [][]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid[i][j] = "."
		}
	}

	if isSet(&encoding, 8) {
		grid[1][1] = "#"
	}
	if isSet(&encoding, 7) {
		grid[0][0] = "#"
	}
	if isSet(&encoding, 6) {
		grid[0][1] = "#"
	}
	if isSet(&encoding, 5) {
		grid[0][2] = "#"
	}
	if isSet(&encoding, 4) {
		grid[1][2] = "#"
	}
	if isSet(&encoding, 3) {
		grid[2][2] = "#"
	}
	if isSet(&encoding, 2) {
		grid[2][1] = "#"
	}
	if isSet(&encoding, 1) {
		grid[2][0] = "#"
	}
	if isSet(&encoding, 0) {
		grid[1][0] = "#"
	}
	return
}

func inputToGrid(input string) [][]string {
	parts := strings.Split(input, "/")
	n := len(parts)

	grid := make([][]string, n)
    for i := range grid {
        grid[i] = make([]string, n)
    }
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = string(parts[i][j])
		}
	}
	return grid
}

func gridToInputFormat(grid [][]string) (input string) {
	for idx, row := range grid {
		for _, cell := range row {
			input += cell
		}
		if idx != len(grid) - 1 {
			input += "/"
		}
	}
	return
}

func printGrid(grid [][]string) {
    for _, row := range grid {
        fmt.Println(row)
    }
	fmt.Println()
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
		input, output := fields[0], fields[2]
		if len(input) == 5 {
			inputGrid := inputToGrid(input)
			outputGrid := inputToGrid(output)
			outputMap2By2[encode2(inputGrid)] = outputGrid
		} else {
			inputGrid := inputToGrid(input)
			outputGrid := inputToGrid(output)
			outputMap3By3[encode3(inputGrid)] = outputGrid
		}
	}
}

func breakInto2(grid [][]string) (pieces [][][]string) {
	n := len(grid)
	for i := 0; i < n / 2; i++ {
		for j := 0; j < n / 2; j++ {
			startRow := i * 2
			startCol := j * 2
			piece := [][]string{
				{"", ""},
				{"", ""},
			}
			for row := range 2 {
				for col := range 2 {
					piece[row][col] = grid[row + startRow][col + startCol]
				}
			}
			pieces = append(pieces, piece)
		}
	}
	return
}

func breakInto3(grid [][]string) (pieces [][][]string) {
	n := len(grid)
	for i := 0; i < n / 3; i++ {
		for j := 0; j < n / 3; j++ {
			startRow := i * 3
			startCol := j * 3
			piece := [][]string{
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
			}
			for row := 0; row < 3; row++ {
				for col := 0; col < 3; col++ {
					piece[row][col] = grid[row + startRow][col + startCol]
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

func flipHorizontally(grid [][]string) (flipped [][]string) {
	n := len(grid)
	flipped = copyGrid(grid)
	for i := 0; i < n; i++ {
		flipped[i][0], flipped[i][n - 1] = flipped[i][n - 1], flipped[i][0]
	}
	return
}

func flipVertically(grid [][]string) (flipped [][]string) {
	n := len(grid)
	flipped = copyGrid(grid)
	flipped[0], flipped[n - 1] = flipped[n - 1], flipped[0]
	return
}

func rotatePiece(grid [][]string) (rotated [][]string) {
	rotated = flipHorizontally(grid)
	n := len(rotated)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rotated[i][j], rotated[j][i] = rotated[j][i], rotated[i][j]
		}
	}
	return
}

func getAllPieces(piece [][]string) (allPieces [][][]string) {
	for i := 0; i < 4; i++ {
		allPieces = append(allPieces, piece)
		allPieces = append(allPieces, flipHorizontally(piece))
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
				// fmt.Println("New Piece")
				// printGrid(newPiece)
				// fmt.Println("Encoding", encode2(newPiece))
				// printGrid(enhancedPiece)
				// fmt.Println()
				// fmt.Printf("For Original Piece: \n")
				// printGrid(piece)
				// fmt.Printf("\nFound After rotation/flipping: \n")
				// printGrid(newPiece)
				// fmt.Printf("\nAfter enhancement: \n")
				// printGrid(enhancedPiece)

				// fmt.Printf("Enhancement Occurred: %s => %s\n\n", gridToInputFormat(newPiece), gridToInputFormat(enhancedPiece))
				return enhancedPiece
			}
		} else {
			enhancedPiece, ok := outputMap3By3[encode3(newPiece)]
			if ok {
				// fmt.Println("New Piece")
				// printGrid(newPiece)
				// fmt.Println("Encoding", encode3(newPiece))
				// printGrid(enhancedPiece)
				// fmt.Println()
				// fmt.Printf("For Original Piece: \n")
				// printGrid(piece)
				// fmt.Printf("\nFound After rotation/flipping: \n")
				// printGrid(newPiece)
				// fmt.Printf("\nAfter enhancement: \n")
				// printGrid(enhancedPiece)
				// fmt.Printf("Enhancement Occurred: %s => %s\n\n", gridToInputFormat(newPiece), gridToInputFormat(enhancedPiece))

				return enhancedPiece
			}
		}
	}
	fmt.Println("Could not enhance piece")
	
	return
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
	
	grid := make([][]string, sizeOfGrid)
    for i := range grid {
        grid[i] = make([]string, sizeOfGrid)
    }

	for idx, piece := range pieces {
		row := (idx / piecesInOneRow) * sizeOfPiece
		col := (idx % piecesInOneRow) * sizeOfPiece
		for i := 0; i < sizeOfPiece; i++ {
			for j := 0; j < sizeOfPiece; j++ {
				grid[row + i][col + j] = piece[i][j]
			}
		}
	}
	return grid
}

func enhance(grid [][]string) (finalGrid [][]string) {
	n := len(grid)
	if n % 2 == 0 {
		pieces := breakInto2(grid)

		// for idx, piece := range pieces {
		// 	fmt.Printf("Piece: %d\n", idx)
		// 	printGrid(piece)
		// }

		enhancedPieces := enhancePieces(pieces, 2)

		// for idx, piece := range enhancedPieces {
		// 	fmt.Printf("Enhanced Piece: %d\n", idx)
		// 	printGrid(piece)
		// }

		finalGrid = mergePieces(enhancedPieces)
	} else {
		pieces := breakInto3(grid)

		// for idx, piece := range pieces {
		// 	fmt.Printf("Piece: %d\n", idx)
		// 	printGrid(piece)
		// }

		enhancedPieces := enhancePieces(pieces, 3)

		// for idx, piece := range enhancedPieces {
		// 	fmt.Printf("Enhanced Piece: %d\n", idx)
		// 	printGrid(piece)
		// }

		finalGrid = mergePieces(enhancedPieces)
	}
	return
}

func main() {
	outputMap2By2 = make(map[int]([][]string))
	outputMap3By3 = make(map[int]([][]string))
	
	readInputAndFillMaps()

	grid := [][]string {
		{".", "#", "."},
		{".", ".", "#"},
		{"#", "#", "#"},
	} 

	for i := range 18 {
		fmt.Printf("Starting enhancement: %d\n", i)
		printGrid(grid)
		grid = enhance(grid)
		fmt.Println("Final Grid")
		printGrid(grid)
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