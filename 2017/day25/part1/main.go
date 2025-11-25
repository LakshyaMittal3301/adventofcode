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

type InstructionSet struct {
	WriteValue    string
	MoveDirection int
	NextState     string
}

type StateInstructionSet struct {
	OnValue0 InstructionSet
	OnValue1 InstructionSet
}

type Blueprint struct {
	BeginState           string
	ChecksumAfterSteps   int
	StateInstructionSets map[string]StateInstructionSet
}

func getLastValue(line string) string {
	parts := strings.Split(line, " ")
	if len(parts) == 0 {
		return ""
	}
	part := parts[len(parts)-1]
	if len(part) == 0 {
		return ""
	}
	return part[:len(part)-1]
}

func isOnlyWhitespace(line string) bool {
	return strings.TrimSpace(line) == ""
}

func parseBlueprint(lines []string) (blueprint Blueprint) {
	blueprint.StateInstructionSets = make(map[string]StateInstructionSet)
	blueprint.BeginState = getLastValue(lines[0])
	parts := strings.Split(lines[1], " ")
	steps, err := strconv.ParseInt(parts[len(parts)-2], 10, 0)
	if err != nil {
		panic("Error in parsing num of steps" + err.Error())
	}
	blueprint.ChecksumAfterSteps = int(steps)

	n := len(lines)
	idx := 2
	for idx < n {
		if isOnlyWhitespace(lines[idx]) {
			idx++
			continue
		}
		state := getLastValue(lines[idx])
		idx++
		blueprint.StateInstructionSets[state] = parseStateInstructionSet(&idx, lines)
	}
	return
}

func parseStateInstructionSet(idx *int, lines []string) (stateInstructionSet StateInstructionSet) {
	*idx++
	stateInstructionSet.OnValue0 = parseInstructionSet(idx, lines)
	*idx++
	stateInstructionSet.OnValue1 = parseInstructionSet(idx, lines)
	return
}

func parseInstructionSet(idx *int, lines []string) (instructionSet InstructionSet) {
	instructionSet.WriteValue = getLastValue(lines[*idx])
	*idx++

	dir := getLastValue(lines[*idx])
	*idx++

	if dir == "left" {
		instructionSet.MoveDirection = -1
	} else {
		instructionSet.MoveDirection = 1
	}

	instructionSet.NextState = getLastValue(lines[*idx])
	*idx++
	return
}

func readInput() []string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get current file info")
	}

	dir := filepath.Dir(filename)

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

func executeInstructionSet(tape map[int]string, position int, instructionSet InstructionSet) (int, string) {
	tape[position] = instructionSet.WriteValue
	position += instructionSet.MoveDirection
	return position, instructionSet.NextState
}

func execute(tape map[int]string, position int, state string, blueprint Blueprint) (int, string) {
	val, ok := tape[position]
	if !ok {
		val = "0"
	}
	if val == "0" {
		return executeInstructionSet(tape, position, blueprint.StateInstructionSets[state].OnValue0)
	} else {
		return executeInstructionSet(tape, position, blueprint.StateInstructionSets[state].OnValue1)
	}
}

func main() {
	lines := readInput()
	blueprint := parseBlueprint(lines)
	tape := make(map[int]string, 0)
	currPosition := 0
	currState := blueprint.BeginState
	// fmt.Printf("%v", blueprint)
	for range blueprint.ChecksumAfterSteps {
		currPosition, currState = execute(tape, currPosition, currState, blueprint)
	}

	count := 0
	for _, val := range tape {
		if val == "1" {
			count++
		}
	}
	fmt.Printf("Answer: %d\n", count)
}
