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

type ParticleInfo struct {
	Positions    [3]int
	Velocity     [3]int
	Acceleration [3]int
}

func readFile() (infoArr []ParticleInfo) {
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
		infoArr = append(infoArr, createParticleInfo(line))
	}
	return
}

func createParticleInfo(line string) (info ParticleInfo) {
	parts := strings.Split(line, ", ")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	info.Positions = read3Integers(parts[0])
	info.Velocity = read3Integers(parts[1])
	info.Acceleration = read3Integers(parts[2])
	return
}

func read3Integers(field string) [3]int {
	n := len(field)
	parts := strings.Split(field[3:n-1], ",")
	n1, _ := strconv.Atoi(parts[0])
	n2, _ := strconv.Atoi(parts[1])
	n3, _ := strconv.Atoi(parts[2])
	return [3]int{n1, n2, n3}
}

type Triple struct {
	X, Y, Z int
}

func tick(infoArr []ParticleInfo) {
	for idx := range infoArr {
		for i := range 3 {
			infoArr[idx].Velocity[i] += infoArr[idx].Acceleration[i]
			infoArr[idx].Positions[i] += infoArr[idx].Velocity[i]
		}
	}
}

func makeTriple(arr [3]int) (triple Triple) {
	triple.X = arr[0]
	triple.Y = arr[1]
	triple.Z = arr[2]
	return
}

func destroy(infoArr []ParticleInfo) (newArr []ParticleInfo) {
	mp := make(map[Triple]int)
	for idx := range infoArr {
		triple := makeTriple(infoArr[idx].Positions)

		if _, ok := mp[triple]; ok {
			mp[triple] = -1
		} else {
			mp[triple] = idx
		}
	}

	for _, idx := range mp {
		if idx != -1 {
			newArr = append(newArr, infoArr[idx])
		}
	}
	return
}

func main() {
	infoArr := readFile()
	n := 200000
	for range n {
		tick(infoArr)
		infoArr = destroy(infoArr)
	}
	fmt.Printf("Particles Left: %d", len(infoArr))
}
