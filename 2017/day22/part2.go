package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func readInput() [][]string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get current file info")
	}

	dir := filepath.Dir(filename)

	path := filepath.Join(dir, "input2.txt")

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([][]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		line := []string{}
		for idx := range text {
			line = append(line, string(text[idx]))
		}
		input = append(input, line)

	}
	return input
}

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func burst(grid [][]string, x, y, dir *int) int {
	isInfected := 0
	switch grid[*x][*y] {
	case ".":
		*dir = (*dir + 3) % 4
		grid[*x][*y] = "W"
	case "W":
		isInfected++
		grid[*x][*y] = "#"
	case "#":
		*dir = (*dir + 1) % 4
		grid[*x][*y] = "F"
	case "F":
		*dir = (*dir + 2) % 4
		grid[*x][*y] = "."
	default:
		panic("Virus is winning!")
	}
	*x += dx[*dir]
	*y += dy[*dir]
	return isInfected
}

func makeGrid(n int) [][]string {
	grid := make([][]string, n)
	for i := 0; i < n; i++ {
		row := make([]string, n)
		for j := 0; j < n; j++ {
			row[j] = "."
		}
		grid[i] = row
	}
	return grid
}

func tripleGrid(grid [][]string, x, y *int) [][]string {
	n := len(grid)
	newSize := 3 * n
	newGrid := makeGrid(newSize)
	for i := range n {
		for j := range n {
			newi := i + n
			newj := j + n
			newGrid[newi][newj] = grid[i][j]
		}
	}

	*x += n
	*y += n
	return newGrid
}

var arrows = []string{
	"↑",
	"→",
	"↓",
	"←",
}

func printGrid(grid [][]string, currX, currY, dir int) {
	clearScreen()

	for i := range grid {
		for j := range grid[i] {
			if i == currX && j == currY {
				fmt.Print(arrows[dir])
			} else {
				fmt.Print(grid[i][j])
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func main() {
	hideCursor()
	defer showCursor()
	grid := readInput()

	n := len(grid) / 2
	currX, currY := n, n
	dir := 0

	iter := 10000000
	count := 0
	// for i := range iter {
	for range iter {
		if currX == len(grid)-1 || currY == len(grid)-1 || currX == 0 || currY == 0 {
			grid = tripleGrid(grid, &currX, &currY)
		}
		// printGrid(grid, currX, currY, dir)
		// fmt.Printf("\nstep: %d   infections: %d\n", i, count)
		// time.Sleep(100 * time.Millisecond)
		count += burst(grid, &currX, &currY, &dir)
	}
	fmt.Printf("Answer: %d", count)

}
