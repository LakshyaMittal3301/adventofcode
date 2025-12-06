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

	path := filepath.Join(dir, "../sampleinput.txt")
	// path := filepath.Join(dir, "../input.txt")

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

func readNum(idx int, nums []string) (num int64) {
	digits := []byte{}
	for i := range len(nums) {
		if nums[i][idx] != ' ' {
			digits = append(digits, nums[i][idx])
		}
	}

	// fmt.Printf("NumRead: %s\n", string(digits))
	num, err := strconv.ParseInt(string(digits), 10, 64)
	if err != nil {
		panic("error parsing num from bytes")
	}
	return
}

func calc(start, end int, op byte, nums []string) (ans int64) {
	ans = readNum(start, nums)
	for i := start + 1; i < end; i++ {
		num := readNum(i, nums)
		if op == '*' {
			ans *= num
		} else {
			ans += num
		}
	}
	return
}

func doMath(lines []string) (ans int64) {
	ops := lines[len(lines)-1]
	numGrid := lines[:len(lines)-1]

	idx := 0
	for idx < len(ops) {
		start := idx
		currOp := ops[idx]
		idx++
		for idx < len(ops) && ops[idx] == ' ' {
			idx++
		}
		// fmt.Printf("Start: %d, End: %d, currOp: %v\n", start, idx, string(currOp))
		if idx == len(ops) {
			ans += calc(start, idx, currOp, numGrid)
		} else {
			ans += calc(start, idx-1, currOp, numGrid)
		}
	}
	return
}

func main() {
	input := readInput()
	ans := doMath(input)
	fmt.Printf("Answer: %d\n", ans)
}
