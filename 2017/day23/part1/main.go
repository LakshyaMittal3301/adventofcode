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

	path := filepath.Join(dir, "../input1.txt")

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

type Type string

const (
	SetIns Type = "set"
	SubIns Type = "sub"
	MulIns Type = "mul"
	JnzIns Type = "jnz"
)

func ParseType(s string) (Type, error) {
	switch s {
	case string(SetIns):
		return SetIns, nil
	case string(SubIns):
		return SubIns, nil
	case string(MulIns):
		return MulIns, nil
	case string(JnzIns):
		return JnzIns, nil
	default:
		panic("Unexpected")
	}
}

type Instruction struct {
	InsType Type
	X       string
	Y       string
}

func getValue(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return registers[s]
	}
	return val
}

func executeInstruction(idx *int64, instructions []Instruction) (count int64) {
	instruction := instructions[*idx]
	switch instruction.InsType {
	case SetIns:
		registers[instruction.X] = getValue(instruction.Y)
		*idx += 1
	case SubIns:
		registers[instruction.X] -= getValue(instruction.Y)
		*idx += 1
	case MulIns:
		registers[instruction.X] *= getValue(instruction.Y)
		*idx += 1
		count = 1
	case JnzIns:
		val := getValue(instruction.X)
		if val != 0 {
			*idx += getValue(instruction.Y)
		} else {
			*idx += 1
		}
	default:
		panic("Unexpected")
	}
	return
}

func execute(instructions []Instruction) (count int64) {
	var idx int64 = 0
	n := int64(len(instructions))

	for idx != n {
		// fmt.Printf("Executing: %d, ins: %v\n", idx, instructions[idx])
		count += executeInstruction(&idx, instructions)
		// fmt.Printf("map: %v\n", registers)
	}
	return
}

var registers map[string]int64

func main() {
	input := readInput()
	instructions := make([]Instruction, 0)
	for idx := range input {
		parts := strings.Split(input[idx], " ")
		insType, _ := ParseType(parts[0])

		instructions = append(instructions, Instruction{
			InsType: insType,
			X:       parts[1],
			Y:       parts[2],
		})
	}

	registers = make(map[string]int64)

	count := execute(instructions)
	fmt.Printf("Answer: %v", count)

}
