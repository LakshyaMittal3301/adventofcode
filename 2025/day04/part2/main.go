package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func readInput() []string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get current file info")
	}

	dir := filepath.Dir(filename)

	// path := filepath.Join(dir, "../sampleinput.txt")
	path := filepath.Join(dir, "../input.txt")

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

func isValid(i, j, n, m int) bool {
	return 0 <= i && i < n && 0 <= j && j < m
}

func countLiftableRolls(grid []string) (count int, newGrid []string) {
	di := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	dj := []int{0, 1, 1, 1, 0, -1, -1, -1}
	n := len(grid)
	m := len(grid[0])

	for i := range n {
		var row strings.Builder
		for j := range m {
			if grid[i][j] != '@' {
				row.WriteByte('.')
				continue
			}
			neigh := 0
			for x := range 8 {
				ix := i + di[x]
				jx := j + dj[x]
				if isValid(ix, jx, n, m) && grid[ix][jx] == '@' {
					neigh++
				}
			}
			if neigh < 4 {
				row.WriteByte('.')
				count++
			} else {
				row.WriteByte('@')
			}
		}
		newGrid = append(newGrid, row.String())
	}
	return
}

func main() {
	input := readInput()
	ans := 0

	for {
		var count int
		count, input = countLiftableRolls(input)
		ans += count
		if count == 0 {
			break
		}
	}

	fmt.Printf("Answer: %d\n", ans)
}
