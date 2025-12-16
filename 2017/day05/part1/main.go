package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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

func getArr(lines []string) []int {
	arr := make([]int, 0)
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic("error parsing input")
		}
		arr = append(arr, num)
	}
	return arr
}

func solve(arr []int) (count int) {
	curr := 0
	for {
		if curr >= len(arr) {
			break
		}
		jump := arr[curr]
		arr[curr]++
		curr += jump
		count++
	}
	return
}

func main() {
	input := readInput()
	arr := getArr(input)
	ans := solve(arr)

	fmt.Printf("Answer: %d\n", ans)
}
