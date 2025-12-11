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

func makeGraph(lines []string) (graph map[string]([]string)) {
	graph = make(map[string]([]string))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		node := parts[0]
		graph[node[:len(node)-1]] = append(graph[node[:len(node)-1]], parts[1:]...)
	}
	return
}

func findNumOfPaths(node string, graph map[string]([]string), dp map[string]int64) (count int64) {
	if node == "out" {
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

func main() {
	input := readInput()
	graph := makeGraph(input)
	dp := map[string]int64{}
	for key := range graph {
		dp[key] = -1
	}
	ans := findNumOfPaths("you", graph, dp)
	fmt.Printf("Answer: %d\n", ans)
}
