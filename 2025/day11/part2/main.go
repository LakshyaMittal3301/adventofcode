package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func readInput() []string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get current file info")
	}

	dir := filepath.Dir(filename)

	// path := filepath.Join(dir, "../sampleinput2.txt")
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

func makeGraph(lines []string) (graph map[string]([]string)) {
	graph = make(map[string]([]string))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		node := parts[0]
		graph[node[:len(node)-1]] = append(graph[node[:len(node)-1]], parts[1:]...)
	}
	return
}

var target string

func findNumOfPaths(node string, graph map[string]([]string), dp map[string]int64) (count int64) {
	if node == target {
		return 1
	}
	if dp[node] != -1 {
		return dp[node]
	}
	for _, neigh := range graph[node] {
		count += findNumOfPaths(neigh, graph, dp)
	}
	dp[node] = count
	return
}

func findTotalPaths(nodes []string, graph map[string][]string) (count int64) {
	count = 1
	dp := map[string]int64{}
	for i := 0; i < len(nodes)-1; i++ {
		for key := range graph {
			dp[key] = -1
		}
		target = nodes[i+1]
		count *= findNumOfPaths(nodes[i], graph, dp)
	}
	return
}

func main() {
	input := readInput()
	graph := makeGraph(input)
	part1 := findTotalPaths([]string{"you", "out"}, graph)
	part2 := findTotalPaths([]string{"svr", "fft", "dac", "out"}, graph)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}
