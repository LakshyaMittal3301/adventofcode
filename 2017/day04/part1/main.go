package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
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

func sortString(str string) string {
	b := []byte(str)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	return string(b)
}

func isValidPassphrase(line string) bool {
	parts := strings.Split(line, " ")
	mp := map[string]struct{}{}
	for _, word := range parts {
		word = sortString(word)
		_, ok := mp[word]
		if ok {
			return false
		}
		mp[word] = struct{}{}
	}
	return true
}

func main() {
	input := readInput()
	ans := 0
	for _, line := range input {
		if isValidPassphrase(line) {
			ans += 1
		}
	}

	fmt.Printf("Answer: %d\n", ans)
}
