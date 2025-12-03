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

func findMaxJoltage(bank string) int {
	maxIdxFirst := 0
	n := len(bank)

	for idx := range n - 1 {
		if bank[maxIdxFirst] < bank[idx] {
			maxIdxFirst = idx
		}
	}

	maxIdxSecond := maxIdxFirst + 1

	for idx := maxIdxFirst + 1; idx < n; idx++ {
		if bank[maxIdxSecond] < bank[idx] {
			maxIdxSecond = idx
		}
	}
	digit1 := int(bank[maxIdxFirst] - '0')
	digit2 := int(bank[maxIdxSecond] - '0')
	return (digit1 * 10) + digit2

}

func main() {
	input := readInput()

	ans := 0

	for idx := range input {
		ans += findMaxJoltage(input[idx])
	}
	fmt.Printf("Answer: %d\n", ans)
}
