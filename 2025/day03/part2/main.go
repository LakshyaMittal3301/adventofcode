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

func findMaxJoltage(bank string, start int, digits int) int64 {
	if digits == 0 {
		return 0
	}
	maxIdx := start
	for idx := start; idx < len(bank)-digits+1; idx++ {
		if bank[maxIdx] < bank[idx] {
			maxIdx = idx
		}
	}

	var mul int64 = 1
	for range digits - 1 {
		mul *= 10
	}
	return mul*int64(bank[maxIdx]-'0') + findMaxJoltage(bank, maxIdx+1, digits-1)

}

func main() {
	input := readInput()

	var ans1 int64 = 0
	var ans2 int64 = 0

	for idx := range input {
		ans1 += findMaxJoltage(input[idx], 0, 2)
		ans2 += findMaxJoltage(input[idx], 0, 12)
	}
	fmt.Printf("Answer 1: %d\n", ans1)
	fmt.Printf("Answer 2: %d\n", ans2)
}
