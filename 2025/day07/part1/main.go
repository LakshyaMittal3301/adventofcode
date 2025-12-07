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
	}
	// count += len(currSet)
	return
}

func main() {
	input := readInput()
	ans := countSplits(input)
	fmt.Printf("Answer: %d\n", ans)
}
