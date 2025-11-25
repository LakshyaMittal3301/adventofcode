package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
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

func getDivision(line []int) int {
	for j := range line {
		n := len(line)
		for k := j + 1; k < n; k++ {
			if line[k]%line[j] == 0 {
				return line[k] / line[j]
			}
		}
	}
	panic("Could not find evenly divisible numbers")
}

func main() {
	lines := readInput()
	spreadsheet := convertToSpreadsheet(lines)

	ans := 0
	for i := range spreadsheet {
		sort.Ints(spreadsheet[i])
		ans += getDivision(spreadsheet[i])
	}
	fmt.Printf("Answer: %d\n", ans)

}
