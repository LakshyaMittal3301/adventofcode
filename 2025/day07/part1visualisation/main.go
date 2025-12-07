package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func readInput() []string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get current file info")
	}

	dir := filepath.Dir(filename)

	// path := filepath.Join(dir, "../sampleinput.txt")
	path := filepath.Join(dir, "../visualisationinput.txt")

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

var allSets []map[int]struct{}

func printTree(idx int, lines []string, currSet map[int]struct{}) {
	clearScreen()
	allSets[idx] = currSet
	for row := range lines {
		for col := range lines[0] {
			_, ok := allSets[row][col]
			if !ok {
				fmt.Print(string(lines[row][col]))
			} else {
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
	time.Sleep(175 * time.Millisecond)

}

func countSplits(lines []string) (count int) {
	start := 0
	for i := range lines[0] {
		if lines[0][i] != '.' {
			start = i
			break
		}
	}

	currSet := map[int]struct{}{}
	currSet[start] = struct{}{}

	printTree(0, lines, map[int]struct{}{})

	for i := 1; i < len(lines); i++ {
		newSet := map[int]struct{}{}

		for key := range currSet {
			if lines[i][key] == '^' {
				newSet[key-1] = struct{}{}
				newSet[key+1] = struct{}{}
				count++
			} else {
				newSet[key] = struct{}{}
			}
		}
		currSet = newSet
		printTree(i, lines, currSet)
	}
	return
}

func clearScreen() {
	fmt.Print("\033[H") // home only — NO clear
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
	input := readInput()
	allSets = make([]map[int]struct{}, len(input))
	ans := countSplits(input)
	fmt.Printf("Answer: %d\n", ans)
}
