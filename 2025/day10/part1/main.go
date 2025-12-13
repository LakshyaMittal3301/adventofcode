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

func makeArrays(strs []string) (arrs [][]int) {
	for _, str := range strs {
		str = str[1 : len(str)-1]
		parts := strings.Split(str, ",")
		arr := []int{}
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				panic("Error parsing number: " + err.Error())
			}
			arr = append(arr, num)
		}
		arrs = append(arrs, arr)
	}
	return
}

func makeJoltage(str string) (joltage []int) {
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic("Error parsing number: " + err.Error())
		}
		joltage = append(joltage, num)
	}
	return
}

type Machine struct {
	finalJoltage []int
	arrs         [][]int
}

func makeStateAndMasks(lines []string) (machines []Machine) {
	for idx := range lines {
		parts := strings.Split(lines[idx], " ")
		finalJoltage := makeJoltage(parts[len(parts)-1])
		arrs := makeArrays(parts[1 : len(parts)-1])
		machines = append(machines, Machine{finalJoltage, arrs})
	}
	return
}

func bfs(finalJoltage []int, arrs [][]int) int {
	q := [][]int{}
	visited := map[int]struct{}{}

	q = append(q, 0)
	visited[0] = struct{}{}

	layer := 0
	for len(q) != 0 {
		size := len(q)
		layer++
		for range size {
			node := q[0]
			q = q[1:]
			if node == finalState {
				return layer - 1
			}
			for _, mask := range masks {
				neigh := node ^ mask
				_, ok := visited[neigh]
				if !ok {
					q = append(q, neigh)
					visited[neigh] = struct{}{}
				}
			}
		}
	}
	panic("Cannot reach state")
}

func main() {
	input := readInput()
	machines := makeStateAndMasks(input)
	ans := 0
	for idx := range machines {
		ans += bfs(machines[idx].finalJoltage, machines[idx].arrs)
	}
	fmt.Printf("Answer: %d\n", ans)
}
