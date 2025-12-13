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

func makeJoltages(str string) (joltages []int) {
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic("error parsing joltage: " + err.Error())
		}
		joltages = append(joltages, num)
	}
	return
}

func getButtonsAndJoltages(line string) (arrs [][]int, joltages []int) {
	parts := strings.Split(line, " ")
	return makeArrays(parts[1 : len(parts)-1]), makeJoltages(parts[len(parts)-1])
}

// Find parity of all possible subsets of input buttons
// Find the subset with the same parity as the required joltage, and that has individual contributions >= individual joltages, take a minimum over those
// Subtract individual contributions from jotlages.
// Divide joltages / 2 and solve.

// Parity -> contribution array, countOfButtons

const inf = 1000000000

type ButtonSubset struct {
	contribution   []int
	countOfButtons int
}

func NewButtonSubset(n int) *ButtonSubset {
	return &ButtonSubset{
		contribution:   make([]int, n),
		countOfButtons: 0,
	}
}

func updateParityAndButtonSubset(button []int, parity *int, buttonSubset *ButtonSubset) {
	for _, i := range button {
		*parity ^= (1 << i)
		buttonSubset.contribution[i] += 1
	}
	buttonSubset.countOfButtons += 1
}

func createParityMap(machines int, buttons [][]int) map[int][]ButtonSubset {
	mp := make(map[int]([]ButtonSubset))
	n := len(buttons)
	for i := 0; i < (1 << n); i++ {
		buttonSubset := NewButtonSubset(machines)
		parity := 0
		for j := range n {
			if (i>>j)&1 == 1 {
				updateParityAndButtonSubset(buttons[j], &parity, buttonSubset)
			}
		}
		mp[parity] = append(mp[parity], *buttonSubset)
	}
	return mp
}

func findParity(joltages []int) (parity int) {
	for i := range joltages {
		if joltages[i]%2 != 0 {
			parity |= (1 << i)
		}
	}
	return
}

func contributionDoesNotExceed(contribution, joltages []int) bool {
	for i := range joltages {
		if contribution[i] > joltages[i] {
			return false
		}
	}
	return true
}

func subtractAndDivideBy2(contribution, joltages []int) (newJoltages []int) {
	newJoltages = make([]int, len(joltages))
	for i := range joltages {
		newJoltages[i] = joltages[i] - contribution[i]
		if newJoltages[i]%2 != 0 {
			panic("odd joltage after subtracting")
		}
		newJoltages[i] /= 2
	}
	return newJoltages
}

func all0(joltages []int) bool {
	for _, j := range joltages {
		if j != 0 {
			return false
		}
	}
	return true
}

func rec(mp map[int][]ButtonSubset, joltages []int) int {
	if all0(joltages) {
		return 0
	}
	parity := findParity(joltages)

	ans := inf
	for _, subset := range mp[parity] {
		if contributionDoesNotExceed(subset.contribution, joltages) {
			newJoltages := subtractAndDivideBy2(subset.contribution, joltages)
			ans = min(ans, subset.countOfButtons+(2*rec(mp, newJoltages)))
		}
	}
	return ans
}

func solve(buttons [][]int, joltages []int) int {
	mp := createParityMap(len(joltages), buttons)
	return rec(mp, joltages)
}

func main() {
	input := readInput()
	ans := 0
	for idx := range input {
		buttons, joltages := getButtonsAndJoltages(input[idx])
		ans += solve(buttons, joltages)
	}
	fmt.Printf("Answer: %v\n", ans)
}
