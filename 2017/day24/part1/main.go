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

type Component struct {
	Port1 int
	Port2 int
}

func createComponents(lines []string) (components []Component) {
	for idx := range lines {
		parts := strings.Split(lines[idx], "/")
		portA, err := strconv.ParseInt(parts[0], 10, 0)
		if err != nil {
			panic("Error parsing input: " + err.Error())
		}
		portB, err := strconv.ParseInt(parts[1], 10, 0)
		if err != nil {
			panic("Error parsing input: " + err.Error())
		}

		components = append(components, Component{
			Port1: int(portA),
			Port2: int(portB),
		})
	}
	return
}

var maxAns int
var taken []bool
var components []Component

func rec(lastStr int, currStr int) {
	maxAns = max(maxAns, currStr)
	for idx, component := range components {
		if !taken[idx] && (component.Port1 == lastStr || component.Port2 == lastStr) {
			taken[idx] = true
			if component.Port1 == lastStr {
				rec(component.Port2, currStr+component.Port1+component.Port2)
			} else {
				rec(component.Port1, currStr+component.Port1+component.Port2)
			}
			taken[idx] = false
		}
	}
}
func main() {
	lines := readInput()
	components = createComponents(lines)
	n := len(components)
	taken = make([]bool, n)

	maxAns = 0
	rec(0, 0)
	fmt.Printf("Answer: %d\n", maxAns)

}
