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

func rotate(num int, rotation string) int {
	dist, err := strconv.ParseInt(rotation[1:], 10, 32)
	if err != nil {
		panic("Error parsing rotation")
	}

	if rotation[0] == 'L' {
		num -= int(dist)
	} else {
		num += int(dist)
	}
	num = (((num) % 100) + 100) % 100
	return num
}

func main() {
	input := readInput()
	currNum := 50
	count := 0
	for _, rotation := range input {
		currNum = rotate(currNum, rotation)
		if currNum == 0 {
			count++
		}
	}

	fmt.Printf("Answer: %d\n", count)
}
