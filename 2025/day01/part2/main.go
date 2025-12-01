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

func rotate(num *int, rotation string) (crossedTimes int) {
	d, err := strconv.ParseInt(rotation[1:], 10, 32)
	if err != nil {
		panic("Error parsing rotation")
	}
	dist := int(d)

	crossedTimes += dist / 100
	oldVal := *num
	switch rotation[0] {
	case 'L':
		*num = (((*num - dist) % 100) + 100) % 100
		if oldVal != 0 && (*num > oldVal || *num == 0) {
			crossedTimes++
		}
	case 'R':
		*num = (((*num + dist) % 100) + 100) % 100
		if *num < oldVal {
			crossedTimes++
		}
	}
	return
}

func main() {
	input := readInput()
	currNum := 50
	count := 0
	for _, rotation := range input {
		crossedTimes := rotate(&currNum, rotation)
		count += crossedTimes
	}

	fmt.Printf("Answer: %d\n", count)
}
