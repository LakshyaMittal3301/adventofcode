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

type Point struct {
	Pos  int64
	Type int
}

func countFreshIngredients(ranges []Range) (count int64) {
	points := []Point{}

	for _, rng := range ranges {
		points = append(points, Point{
			Pos:  rng.Start,
			Type: 0,
		})
		points = append(points, Point{
			Pos:  rng.End,
			Type: 1,
		})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].Pos == points[j].Pos {
			return points[i].Type < points[j].Type
		}
		return points[i].Pos < points[j].Pos
	})

	currRanges := 0
	var start int64 = 0

	for _, point := range points {
		if point.Type == 0 {
			if currRanges == 0 {
				start = point.Pos
			}
			currRanges++
		} else {
			currRanges--
			if currRanges == 0 {
				count += point.Pos - start + 1
			}
		}

		if currRanges < 0 {
			panic("something went wrong, closed more ranges than opened")
		}
	}
	return
}

func main() {
	input := readInput()
	ranges, _ := getRangesAndIds(input)

	ans := countFreshIngredients(ranges)

	fmt.Printf("Answer: %d\n", ans)
}
