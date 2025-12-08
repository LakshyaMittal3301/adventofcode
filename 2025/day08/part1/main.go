package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func readInput() []string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get current file info")
	}

	dir := filepath.Dir(filename)

	// path := filepath.Join(dir, "../sampleinput.xtxt")
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

type Triplet struct {
	id int
	x  int
	y  int
	z  int
}

type Edge struct {
	t1   Triplet
	t2   Triplet
	dist int
}

func makeTriplets(lines []string) (triplets []Triplet) {
	for idx := range lines {
		parts := strings.Split(lines[idx], ",")
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("error parsing integer" + err.Error())
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("error parsing integer" + err.Error())
		}
		num3, err := strconv.Atoi(parts[2])
		if err != nil {
			panic("error parsing integer" + err.Error())
		}
		triplets = append(triplets, Triplet{
			id: idx,
			x:  num1,
			y:  num2,
			z:  num3,
		})
	}
	return
}

func makeEdge(t1 Triplet, t2 Triplet) Edge {
	dist := ((t1.x - t2.x) * (t1.x - t2.x)) + ((t1.y - t2.y) * (t1.y - t2.y)) + ((t1.z - t2.z) * (t1.z - t2.z))
	return Edge{
		t1:   t1,
		t2:   t2,
		dist: dist,
	}
}

func makeEdgesAndSort(triplets []Triplet) (edges []Edge) {
	n := len(triplets)
	for i := range n {
		for j := i + 1; j < n; j++ {
			edges = append(edges, makeEdge(triplets[i], triplets[j]))
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})
	return
}

type DSU struct {
	parent []int
	rank   []int
}

func (d *DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) merge(x, y int) bool {
	x = d.find(x)
	y = d.find(y)
	if x == y {
		return false
	}
	if d.rank[x] < d.rank[y] {
		x, y = y, x
	}

	d.parent[y] = x
	d.rank[x] += d.rank[y]
	return true
}

func NewDSU(n int) *DSU {
	dsu := DSU{
		parent: make([]int, n),
		rank:   make([]int, n),
	}

	for i := range n {
		dsu.parent[i] = i
		dsu.rank[i] = 1
	}
	return &dsu
}

func main() {
	input := readInput()
	triplets := makeTriplets(input)
	edges := makeEdgesAndSort(triplets)

	dsu := NewDSU(len(triplets))

	for i := range 1000 {
		dsu.merge(edges[i].t1.id, edges[i].t2.id)
	}

	compSizesMap := map[int]int{}

	for _, t := range triplets {
		par := dsu.find(t.id)
		compSizesMap[par] = dsu.rank[par]
	}

	compSizes := []int{}
	for _, val := range compSizesMap {
		compSizes = append(compSizes, val)
	}
	sort.Slice(compSizes, func(i, j int) bool {
		return compSizes[i] > compSizes[j]
	})

	ans := compSizes[0] * compSizes[1] * compSizes[2]

	fmt.Printf("Answer: %d\n", ans)
}
