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
	// path := filepath.Join(dir, "../sampleinput.txt")

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	text := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		text = append(text, parts...)
	}
	return text
}

func makeRanges(input []string) (ranges []Range) {
	for idx := range input {
		parts := strings.Split(input[idx], "-")
		num1, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic("Error in parsing")
		}
		num2, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic("Error in parsing")
		}
		if num1 > num2 {
			num1, num2 = num2, num1
		}
		ranges = append(ranges, Range{
			Start: num1,
			End:   num2,
		})
	}
	return
}

type Range struct {
	Start int64
	End   int64
}

func isNumInRanges(num int64, ranges []Range) bool {
	for idx := range ranges {
		if ranges[idx].Start <= num && num <= ranges[idx].End {
			return true
		}
	}
	return false
}

func calculate(num int, ranges []Range, visited map[int64]struct{}) (count int64) {
	numStr := strconv.Itoa(num)
	str := numStr
	for {
		str += numStr
		num1, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic("Error creating number")
		}
		if num1 > 10000000000 {
			return
		}
		_, ok := visited[num1]
		if !ok && isNumInRanges(num1, ranges) {
			count += num1
			visited[num1] = struct{}{}
		}

	}
}

func countInRanges(ranges []Range) (count int64) {
	visited := make(map[int64]struct{})
	for i := 1; i < 100000; i++ {
		count += calculate(i, ranges, visited)
	}
	return
}

func main() {
	input := readInput()
	ranges := makeRanges(input)
	ans := countInRanges(ranges)
	fmt.Printf("Answer: %d\n", ans)

}
