package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func readInput() []string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get current file info")
	}

	dir := filepath.Dir(filename)

	path := filepath.Join(dir, "../sampleinput2.txt")
	// path := filepath.Join(dir, "../input.txt")

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	text := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)

	}
	return text
}

type Piece struct {
	id    int
	shape [3][3]bool
}

func findCellsForShape(shape [3][3]bool, cellRow, cellCol, mArea int) (cells []int) {
	for i := range 3 {
		for j := range 3 {
			if shape[i][j] {
				ix := cellRow + i
				jx := cellCol + j
				cells = append(cells, (ix*mArea)+jx)
			}
		}
	}
	return
}

func addRowsForPiece(shape [3][3]bool, cellRow, cellCol, mArea, qty, offset, totalCells int) (newRows [][]int) {
	cells := findCellsForShape(shape, cellRow, cellCol, mArea)
	for i := range qty {
		row := []int{}
		row = append(row, cells...)
		row = append(row, totalCells+offset+i)
		// fmt.Printf("Adding row: %v\n", row)
		newRows = append(newRows, row)
	}
	return
}

func createRows(pieces []Piece, idQty []int, nArea, mArea int) (rows [][]int, totalColumns int) {
	totalCells := nArea * mArea
	uniquePieces := len(idQty)
	totalPieces := 0
	idOffset := make([]int, uniquePieces)
	for id := range idQty {
		idOffset[id] = totalPieces
		totalPieces += idQty[id]
	}

	totalColumns = totalCells + totalPieces

	rows = make([][]int, 0)

	// For empty cells
	for i := range totalCells {
		row := []int{i}
		rows = append(rows, row)
	}

	for _, piece := range pieces {
		id, shape := piece.id, piece.shape
		for i := range nArea - 2 {
			for j := range mArea - 2 {
				rows = append(rows, addRowsForPiece(shape, i, j, mArea, idQty[id], idOffset[id], totalCells)...)
			}
		}
	}
	return
}

type Query struct {
	nArea, mArea int
	quantities   []int
}

func readPiece(lines []string, idx *int) (piece Piece) {
	line := lines[*idx]
	*idx++
	idStr := line[:len(line)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic("error parsing id: " + err.Error())
	}

	piece.id = id
	piece.shape = [3][3]bool{}
	for i := range 3 {
		for j := range 3 {
			if lines[*idx][j] == '#' {
				piece.shape[i][j] = true
			}
		}
		*idx++
	}
	return

}
func readQuery(line string) (query Query) {
	parts := strings.Split(line, ":")
	dimensionStr, qtyStr := parts[0], parts[1]
	parts = strings.Split(dimensionStr, "x")
	nArea, err := strconv.Atoi(parts[0])
	if err != nil {
		panic("error parsing dimension: " + err.Error())
	}
	mArea, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("error parsing dimension: " + err.Error())
	}
	query.nArea = nArea
	query.mArea = mArea
	qtyStr = strings.TrimSpace(qtyStr)
	parts = strings.Split(qtyStr, " ")
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic("error parsing qty: " + err.Error())
		}
		query.quantities = append(query.quantities, num)
	}
	return
}

func parseInput(lines []string) (pieces []Piece, queries []Query) {
	pieces = make([]Piece, 0)
	queries = make([]Query, 0)
	idx := 0
	for range 6 {
		pieces = append(pieces, readPiece(lines, &idx))
		idx++
	}
	for idx < len(lines) {
		queries = append(queries, readQuery(lines[idx]))
		idx++
	}
	return
}

func flip(shape [3][3]bool) (newShape [3][3]bool) {
	for i := range 3 {
		for j := range 3 {
			newShape[i][2-j] = shape[i][j]
		}
	}

	return
}

func rotate(shape [3][3]bool) (newShape [3][3]bool) {
	for i := range 3 {
		for j := range 3 {
			newShape[i][j] = shape[j][i]
		}
	}
	return
}

func inPieces(piece Piece, pieces []Piece) bool {
	for _, currPiece := range pieces {
		if isShapeSame(currPiece.shape, piece.shape) {
			return true
		}
	}
	return false
}

func isShapeSame(s1, s2 [3][3]bool) bool {
	for i := range 3 {
		for j := range 3 {
			if s1[i][j] != s2[i][j] {
				return false
			}
		}
	}
	return true
}

func rotateAndFlipPiece(piece Piece) (pieces []Piece) {
	id := piece.id
	shape := piece.shape
	var newPiece Piece
	for range 4 {
		newPiece = Piece{
			id:    id,
			shape: shape,
		}
		if !inPieces(newPiece, pieces) {
			pieces = append(pieces, newPiece)
		}

		shape = flip(shape)
		newPiece = Piece{
			id:    id,
			shape: shape,
		}
		if !inPieces(newPiece, pieces) {
			pieces = append(pieces, newPiece)
		}

		shape = rotate(shape)
	}
	// fmt.Printf("Pieces for Id: %d\n", id)
	// for _, newPiece := range pieces {
	// 	printShape(newPiece.shape)
	// 	fmt.Println()
	// }
	return
}

func printShape(shape [3][3]bool) {
	for i := range 3 {
		for j := range 3 {
			if shape[i][j] {
				fmt.Printf("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func rotateAndFlipPieces(pieces []Piece) (newPieces []Piece) {
	for _, piece := range pieces {
		newPieces = append(newPieces, rotateAndFlipPiece(piece)...)
	}
	return
}

func main() {
	lines := readInput()
	pieces, queries := parseInput(lines)
	pieces = rotateAndFlipPieces(pieces)
	for _, q := range queries {
		rows, totalColumns := createRows(pieces, q.quantities, q.nArea, q.mArea)
		// fmt.Printf("Total Columns: %d\n", totalColumns)
		// for i := range rows {
		// 	fmt.Printf("Row: %v\n", rows[i])
		// }
		dlx := NewDLX(rows, totalColumns)
		dlx.Search()
		if dlx.HasSolution {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	}
	// rows := [][]int{
	// 	/* A */ {2, 4, 5}, // C E F
	// 	/* B */ {0, 3, 6}, // A D G
	// 	/* C */ {1, 2, 5}, // B C F
	// 	/* D */ {0, 3}, // A D
	// 	/* E */ {1, 6}, // B G
	// 	/* F */ {3, 4, 6}, // D E G
	// }
	// dlx := NewDLX(rows, 7)
	// dlx.Search()

	// if !dlx.HasSolution {
	// 	fmt.Printf("No Solution Found\n")
	// } else {
	// 	fmt.Printf("Solution rows:\n")
	// 	for i := range dlx.Solution {
	// 		fmt.Printf("%d\n", dlx.Solution[i]+1)
	// 	}
	// }
}
