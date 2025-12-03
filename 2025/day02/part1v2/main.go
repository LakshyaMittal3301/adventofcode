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

func breakRangesToSameLength(ranges []Range) (newRanges []Range) {
	for _, rng := range ranges {
		startLength := int(math.Log10(float64(rng.Start))) + 1
		endLength := int(math.Log10(float64(rng.End))) + 1

		start := rng.Start
		for startLength != endLength {
			end := int64(math.Pow10(startLength) - 1)
			newRanges = append(newRanges, Range{
				Start: start,
				End:   end,
			})
			start = end + 1
			startLength++
		}

		newRanges = append(newRanges, Range{
			Start: start,
			End:   rng.End,
		})
	}
	return
}

func sumTillN(num int64) int64 {
	return (num * (num + 1)) / 2
}

func calculateInvalidInRange(rng Range) int64 {
	numLength := int(math.Log10(float64(rng.Start))) + 1
	if numLength%2 != 0 {
		return 0
	}
	power10 := int64(math.Pow10(numLength / 2))
	startFirstHalf := rng.Start / power10
	startSecondHalf := rng.Start % power10
	endFirstHalf := rng.End / power10
	endSecondHalf := rng.End % power10

	ans := ((sumTillN(endFirstHalf) - sumTillN(startFirstHalf-1)) * (power10 + 1))

	if startSecondHalf > startFirstHalf {
		ans -= startFirstHalf * (power10 + 1)
	}
	if endSecondHalf < endFirstHalf {
		ans -= endFirstHalf * (power10 + 1)
	}
	return ans

}

func main() {
	input := readInput()
	ranges := makeRanges(input)
	ranges = breakRangesToSameLength(ranges)
	var ans int64 = 0
	for idx := range ranges {
		ans += calculateInvalidInRange(ranges[idx])
	}

	fmt.Printf("Answer: %d\n", ans)

}
