package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func readInput() []string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get current file info")
	}

	dir := filepath.Dir(filename)

	// path := filepath.Join(dir, "../sampleinput.txt")
	path := filepath.Join(dir, "../visualisationinput.txt")

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

func randSign() int {
	if rand.Intn(2) == 0 {
		return -1
	}
	return 1
}

type Pair struct {
	i int
	j int
}

func printChar(i, j int, beams []Pair, lines []string) {
	if lines[i][j] == 'S' {
		fmt.Print(string(lines[i][j]))
		return
	}
	for idx := range beams {
		if beams[idx].i == i && beams[idx].j == j {
			fmt.Print("|")
			return
		}
	}
	fmt.Print(string(lines[i][j]))

}

func printTree(beams []Pair, lines []string) {
	clearScreen()
	for row := range lines {
		for col := range lines[0] {
			printChar(row, col, beams, lines)
		}
		fmt.Println()
	}
	time.Sleep(100 * time.Millisecond)

}

func iter(beams []Pair, lines []string) (newBeams []Pair) {
	printTree(beams, lines)
	for _, beam := range beams {
		if beam.i == len(lines)-1 {
			continue
		}
		if lines[beam.i+1][beam.j] == '^' {
			newBeams = append(newBeams, Pair{i: beam.i + 1, j: beam.j + randSign()})
		} else {
			newBeams = append(newBeams, Pair{i: beam.i + 1, j: beam.j})
		}
	}
	return
}

func countSplits(lines []string) (count int) {
	start := 0
	for i := range lines[0] {
		if lines[0][i] != '.' {
			start = i
			break
		}
	}
	beams := []Pair{}

	for idx := range 1000 {
		beams = iter(beams, lines)
		if idx%4 == 0 {
			beams = append(beams, Pair{i: 1, j: start})
		}
	}
	return
}

func clearScreen() {
	fmt.Print("\033[H") // home only — NO clear
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func main() {
	hideCursor()
	defer showCursor()
	input := readInput()
	ans := countSplits(input)
	fmt.Printf("Answer: %d\n", ans)
}
