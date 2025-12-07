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

func rec(row, col int, lines []string, dp [][]int) int {
	if row == len(lines) {
		return 1
	}
	if dp[row][col] != -1 {
		return dp[row][col]
	}
	count := 0

	if lines[row][col] == '^' {
		count += rec(row+1, col-1, lines, dp)
		count += rec(row+1, col+1, lines, dp)
	} else {
		count += rec(row+1, col, lines, dp)
	}
	dp[row][col] = count
	return count
}

func countSplits(lines []string) (count int) {
	start := 0
	for i := range lines[0] {
		if lines[0][i] != '.' {
			start = i
			break
		}
	}
	n := len(lines)
	m := len(lines[0])
	dp := make([][]int, 0)
	for range n {
		row := make([]int, 0)
		for range m {
			row = append(row, -1)
		}
		dp = append(dp, row)
	}
	count = rec(1, start, lines, dp)
	return
}

func main() {
	input := readInput()
	ans := countSplits(input)
	fmt.Printf("Answer: %d\n", ans)
}
