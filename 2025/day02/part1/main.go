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

func countInRanges(ranges []Range) (count int64) {
	for i := 1; i < 100000; i++ {
		str := strconv.Itoa(i)
		num, err := strconv.ParseInt(str+str, 10, 64)
		if err != nil {
			panic("Error creating number")
		}
		// fmt.Printf("str: %s, str + str: %s, num: %d\n", str, str+str, num)
		if isNumInRanges(num, ranges) {
			count += num
		}
	}
	return
}

func main() {
	input := readInput()
	ranges := makeRanges(input)
	ans := countInRanges(ranges)
	fmt.Printf("Answer: %d\n", ans)

}
