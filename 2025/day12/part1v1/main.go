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

	path := filepath.Join(dir, "../sampleinput.txt")
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

func countActiveInPieces(pieces []Piece) (count []int) {
	count = make([]int, len(pieces))
	for idx, piece := range pieces {
		for i := range 3 {
			for j := range 3 {
				if piece.shape[i][j] {
					count[idx]++
				}
			}
		}
	}
	return
}

func mul(activeArr, qty []int) (count int) {
	for i := range activeArr {
		count += activeArr[i] * qty[i]
	}
	return
}

func main() {
	lines := readInput()
	pieces, queries := parseInput(lines)
	activeArr := countActiveInPieces(pieces)
	ans := 0
	for _, q := range queries {
		required := mul(activeArr, q.quantities)
		if required <= q.nArea*q.mArea {
			ans += 1
		}
	}
	fmt.Printf("Answer: %d", ans)
}
