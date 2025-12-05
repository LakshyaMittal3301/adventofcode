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

type Range struct {
	Start int64
	End   int64
}

func getRangesAndIds(input []string) (ranges []Range, ids []int64) {
	idx := 0
	for idx < len(input) {
		if input[idx] == "" {
			break
		}
		parts := strings.Split(input[idx], "-")
		num1, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic("error parsing num: " + parts[0])
		}

		num2, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic("error parsing num: " + parts[1])
		}

		ranges = append(ranges, Range{
			Start: num1,
			End:   num2,
		})

		idx++
	}

	idx++

	for idx < len(input) {
		num, err := strconv.ParseInt(input[idx], 10, 64)
		if err != nil {
			panic("error parsing num: " + input[idx])
		}
		ids = append(ids, num)
		idx++
	}
	return
}

func idInRanges(id int64, ranges []Range) bool {
	for _, rng := range ranges {
		if rng.Start <= id && id <= rng.End {
			return true
		}
	}
	return false
}

func countFreshIngredients(ranges []Range, ids []int64) (count int) {
	for _, id := range ids {
		if idInRanges(id, ranges) {
			count++
		}
	}
	return
}

func main() {
	input := readInput()
	ranges, ids := getRangesAndIds(input)

	ans := countFreshIngredients(ranges, ids)

	fmt.Printf("Answer: %d\n", ans)
}
