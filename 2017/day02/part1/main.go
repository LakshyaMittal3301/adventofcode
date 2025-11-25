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

func convertToSpreadsheet(lines []string) [][]int {
	n := len(lines)
	spreadsheet := make([][]int, n)

	for i := range n {
		parts := strings.Split(lines[i], "\t")
		spreadsheet[i] = make([]int, len(parts))
		for j := range parts {
			num, err := strconv.ParseInt(parts[j], 10, 0)
			if err != nil {
				panic("Error parsing input: " + err.Error())
			}
			spreadsheet[i][j] = int(num)
		}
	}
	return spreadsheet
}

func main() {
	lines := readInput()
	spreadsheet := convertToSpreadsheet(lines)
	ans := 0
	for i := range spreadsheet {
		maxNum := -10000000
		minNum := 10000000
		for j := range spreadsheet[i] {
			maxNum = max(maxNum, spreadsheet[i][j])
			minNum = min(minNum, spreadsheet[i][j])
		}
		ans += maxNum - minNum
	}
	fmt.Printf("Answer: %d\n", ans)
}
