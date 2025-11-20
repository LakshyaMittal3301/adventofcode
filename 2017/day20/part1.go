package main

import (
	"bufio"
	"fmt"
	"math"
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

func absSum(arr [3]int) (sum int) {
	sum = int(math.Abs(float64(arr[0])) + math.Abs(float64(arr[1])) + math.Abs(float64(arr[2])))
	return
}

func main() {
	infoArr := readFile()
	var ans int = 1
	for idx := range infoArr {
		if absSum(infoArr[idx].Acceleration) != absSum(infoArr[ans].Acceleration) {
			if absSum(infoArr[idx].Acceleration) < absSum((infoArr[ans].Acceleration)) {
				ans = idx
			}
		} else if absSum(infoArr[idx].Velocity) != absSum(infoArr[ans].Velocity) {
			if absSum(infoArr[idx].Velocity) < absSum((infoArr[ans].Velocity)) {
				ans = idx
			}
		} else if absSum(infoArr[idx].Positions) != absSum(infoArr[ans].Positions) {
			if absSum(infoArr[idx].Positions) < absSum((infoArr[ans].Positions)) {
				ans = idx
			}
		}
	}
	// fmt.Printf("Answer found: %d\n", ans)
	// fmt.Printf("Info: %v\n", infoArr[ans])

	for idx := range infoArr {
		if absSum(infoArr[idx].Acceleration) == 1 {
			fmt.Printf("id: %d, info: %v\n", idx, infoArr[idx])
		}
	}
}
