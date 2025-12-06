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

func getNumsAndOps(lines []string) (nums [][]int64, ops []string) {
	for idx, line := range lines {
		if idx == len(lines)-1 {
			break
		}
		row := []int64{}
		parts := strings.Fields(line)

		for _, part := range parts {
			num, err := strconv.ParseInt(part, 10, 64)
			if err != nil {
				panic("error parsing number: " + err.Error())
			}
			row = append(row, num)
		}
		nums = append(nums, row)
	}

	ops = strings.Fields(lines[len(lines)-1])
	return
}

func doMath(nums [][]int64, ops []string) (ans int64) {
	for j := range nums[0] {
		var colAns int64
		if ops[j] == "*" {
			colAns = 1
		} else {
			colAns = 0
		}

		for i := range nums {
			if ops[j] == "*" {
				colAns *= nums[i][j]
			} else {
				colAns += nums[i][j]
			}
		}
		ans += colAns
	}
	return
}

func main() {
	input := readInput()
	nums, ops := getNumsAndOps(input)
	// fmt.Printf("nums: %v, lenRow: %d, lenCol: %d\n", nums, len(nums), len(nums[0]))
	// fmt.Printf("ops: %v, len: %d\n", ops, len(ops))
	ans := doMath(nums, ops)
	fmt.Printf("Answer: %d\n", ans)
}
