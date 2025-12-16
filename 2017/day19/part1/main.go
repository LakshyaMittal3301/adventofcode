package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

func isAlpha(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

func isValid(i, j, n, m int) bool {
	return 0 <= i && i < n && 0 <= j && j < m
}

func findStart(grid []string) (int, int) {
	for j := range grid[0] {
		if grid[0][j] == '|' {
			return 0, j
		}
	}
	return 0, -1
}

var di []int
var dj []int

func moveInDir(i, j, dir int) (int, int) {
	i += di[dir]
	j += dj[dir]
	return i, j
}

func solve(grid []string) (ans string) {
	// find start
	n := len(grid)
	m := len(grid[0])
	i, j := findStart(grid)

	// dir == down
	dir := 2

	for {
		for isValid(i, j, n, m) && grid[i][j] != ' ' && grid[i][j] != '+' {
			if isAlpha(grid[i][j]) {
				ans += string(grid[i][j])
			}
			i, j = moveInDir(i, j, dir)
		}
		if isValid(i, j, n, m) && grid[i][j] == '+' {
			// change dir
			leftDir, rightDir := (dir+3)%4, (dir+1)%4
			il, jl := moveInDir(i, j, leftDir)
			if isValid(il, jl, n, m) && grid[il][jl] != ' ' {
				dir = leftDir
			} else {
				dir = rightDir
			}
			i, j = moveInDir(i, j, dir)
		} else {
			break
		}
	}
	return
}

func main() {
	di = []int{-1, 0, 1, 0}
	dj = []int{0, 1, 0, -1}

	input := readInput()
	ans := solve(input)
	fmt.Printf("Answer: %s\n", ans)
}
