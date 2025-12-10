package main

import (
	"bufio"
	"fmt"
	"math"
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

type Pair struct {
	x int64
	y int64
}

func makePairs(lines []string) (pairs []Pair) {
	for _, line := range lines {
		parts := strings.Split(line, ",")
		num1, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic("error parsing int: " + err.Error())
		}
		num2, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic("error parsing int: " + err.Error())
		}
		pairs = append(pairs, Pair{
			x: num1,
			y: num2,
		})
	}
	return
}
func getArea(pair1 Pair, pair2 Pair) (area int64) {
	return int64(math.Abs(float64((pair1.x - pair2.x + 1) * (pair1.y - pair2.y + 1))))
}

func getLargestRectangleArea(pairs []Pair) (area int64) {
	for i := range pairs {
		for j := i + 1; j < len(pairs); j++ {
			area = max(area, getArea(pairs[i], pairs[j]))
		}
	}
	return
}

func main() {
	input := readInput()
	pairs := makePairs(input)
	ans := getLargestRectangleArea(pairs)
	fmt.Printf("Answer: %d\n", ans)
}
