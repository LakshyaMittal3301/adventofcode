package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

type Pair struct {
	i int
	j int
}

func isValid(i, j, n, m int) bool {
	return 0 <= i && i < n && 0 <= j && j < m
}

func countLiftableRolls(grid []string) (count int) {
	di := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	dj := []int{0, 1, 1, 1, 0, -1, -1, -1}
	n := len(grid)
	m := len(grid[0])

	// Degree: Number of neighbours
	degree := make([][]int, 0)
	visited := make([][]bool, 0)
	for range n {
		degree = append(degree, make([]int, m))
		visited = append(visited, make([]bool, m))
	}

	// Calculate degrees
	for i := range n {
		for j := range m {
			for x := range 8 {
				ix := i + di[x]
				jx := j + dj[x]
				if isValid(ix, jx, n, m) && grid[ix][jx] == '@' {
					degree[i][j] += 1
				}
			}
		}
	}

	q := make([]Pair, 0)

	// Push all cells with degree < 4 to the queue
	for i := range n {
		for j := range m {
			if grid[i][j] == '@' && degree[i][j] < 4 {
				q = append(q, Pair{i, j})
				visited[i][j] = true
			}
		}
	}

	for len(q) != 0 {
		p := q[len(q)-1]
		q = q[:len(q)-1]
		count++

		// Iterate over neighbors
		for x := range 8 {
			ix := p.i + di[x]
			jx := p.j + dj[x]

			if isValid(ix, jx, n, m) && grid[ix][jx] == '@' {

				degree[ix][jx]--

				// If degree just became < 4, ready to be removed -> add to queue
				if degree[ix][jx] < 4 && !visited[ix][jx] {
					q = append(q, Pair{ix, jx})
					visited[ix][jx] = true
				}
			}
		}
	}
	return
}

func main() {
	input := readInput()
	ans := countLiftableRolls(input)

	fmt.Printf("Answer: %d\n", ans)
}
