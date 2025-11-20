package main

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
)

func readFile() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get current file info")
	}

	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "input.txt")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		return line
	}
	panic("No text found")
}

func captcha(num string) (ans int64) {
	n := len(num)
	for idx := range n {
		if idx == 0 {
			continue
		}
		if num[idx] == num[idx-1] {
			ans += int64(num[idx] - '0')
		}
	}
	if num[0] == num[n-1] {
		ans += int64(num[0] - '0')
	}
	return ans
}

func main() {
	num := readFile()
	println(captcha(num))
}
