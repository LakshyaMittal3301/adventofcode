package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/sets/treeset"
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
	First  int
	Second int
}

func makePairsAndSort(lines []string) (pairs []Pair) {
	for _, line := range lines {
		parts := strings.Split(line, ",")
		num1, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic("error parsing int: " + err.Error())
		}
		num2, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic("error parsing int: " + err.Error())
		}
		pairs = append(pairs, Pair{
			First:  int(num1),
			Second: int(num2),
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].First == pairs[j].First {
			return pairs[i].Second < pairs[j].Second
		}
		return pairs[i].First < pairs[j].First
	})
	return
}

type Level struct {
	points *treeset.Set
}

func (l *Level) AddPair(p Pair) {
	if l.points.Contains(p.First) {
		l.points.Remove(p.First)
	} else {
		l.points.Add(p.First)
	}
	if l.points.Contains(p.Second) {
		l.points.Remove(p.Second)
	} else {
		l.points.Add(p.Second)
	}
}

func (l *Level) IsInRange(p Pair) bool {
	it := l.points.Iterator()
	for it.Next() {
		firstVal := it.Value().(int)
		if !it.Next() {
			panic("odd number of elements in level")
		}
		secondVal := it.Value().(int)
		if firstVal <= p.First && p.Second <= secondVal {
			return true
		}
	}
	return false
}

func NewLevel() *Level {
	return &Level{
		points: treeset.NewWithIntComparator(),
	}
}

func (l *Level) Clone() *Level {
	newSet := treeset.NewWithIntComparator()

	it := l.points.Iterator()
	for it.Next() {
		newSet.Add(it.Value())
	}

	return &Level{
		points: newSet,
	}
}

func makeLevelMap(pairs []Pair) map[int]*Level {
	levels := make(map[int]*Level)
	n := len(pairs)
	idx := 0
	level := NewLevel()
	for idx < n {
		startX := pairs[idx].First
		sameLevelPairs := []Pair{}
		for idx < n && pairs[idx].First == startX {
			sameLevelPairs = append(sameLevelPairs, pairs[idx])
			idx++
		}
		if len(sameLevelPairs)%2 != 0 {
			panic("odd lengthed level: HOW??")
		}
		for j := 0; j < len(sameLevelPairs); j += 2 {
			level.AddPair(Pair{sameLevelPairs[j].Second, sameLevelPairs[j+1].Second})
		}
		levels[startX] = level.Clone()
	}
	return levels
}

func formValidRectangle(p1, p2 Pair, levelMap map[int]*Level) bool {
	startX := min(p1.First, p2.First)
	endX := max(p1.First, p2.First)

	startY := min(p1.Second, p2.Second)
	endY := max(p1.Second, p2.Second)

	for levelIdx := range levelMap {
		if levelIdx < startX || levelIdx >= endX {
			continue
		}
		if !levelMap[levelIdx].IsInRange(Pair{startY, endY}) {
			return false
		}
	}
	return true
}

func countMaxArea(pairs []Pair, levelMap map[int]*Level) (ans int64) {
	for i := range len(pairs) {
		for j := i + 1; j < len(pairs); j++ {
			if formValidRectangle(pairs[i], pairs[j], levelMap) {
				ans = max(ans, getArea(pairs[i], pairs[j]))
			}
		}
	}
	return ans
}

func getArea(p1 Pair, p2 Pair) (area int64) {
	return int64(math.Abs(float64(p1.First-p2.First))+1) * int64(math.Abs(float64(p1.Second-p2.Second))+1)
}

func main() {
	input := readInput()
	pairs := makePairsAndSort(input)
	levelMap := makeLevelMap(pairs)
	ans := countMaxArea(pairs, levelMap)
	fmt.Printf("Answer: %v\n", ans)
}
