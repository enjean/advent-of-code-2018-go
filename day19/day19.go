package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var registers [6]int

type program struct {
	ip int
	instructions []instruction
}

type instruction struct {
	operation string
	a, b, c int
}

var operations = map[string]func(a, b, c int){
	"addr": func(a, b, c int) {
		registers[c] = registers[a] + registers[b]
	},
	"addi": func(a, b, c int) {
		registers[c] = registers[a] + b
	},
	"mulr": func(a, b, c int) {
		registers[c] = registers[a] * registers[b]
	},
	"muli": func(a, b, c int) {
		registers[c] = registers[a] * b
	},
	"banr": func(a, b, c int) {
		registers[c] = registers[a] & registers[b]
	},
	"bani": func(a, b, c int) {
		registers[c] = registers[a] & b
	},
	"borr": func(a, b, c int) {
		registers[c] = registers[a] | registers[b]
	},
	"bori": func(a, b, c int) {
		registers[c] = registers[a] | b
	},
	"setr": func(a, b, c int) {
		registers[c] = registers[a]
	},
	"seti": func(a, b, c int) {
		registers[c] = a
	},
	"gtir": func(a, b, c int) {
		registers[c] = boolToInt(a > registers[b])
	},
	"gtri": func(a, b, c int) {
		registers[c] = boolToInt(registers[a] > b)
	},
	"gtrr": func(a, b, c int) {
		registers[c] = boolToInt(registers[a] > registers[b])
	},
	"eqir": func(a, b, c int) {
		registers[c] = boolToInt(a == registers[b])
	},
	"eqri": func(a, b, c int) {
		registers[c] = boolToInt(registers[a] == b)
	},
	"eqrr": func(a, b, c int) {
		registers[c] = boolToInt(registers[a] == registers[b])
	},
}

func boolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func execute(program *program) {
	instructionPointer := 0
	for instructionPointer < len(program.instructions) {
		registers[program.ip] = instructionPointer
		instruction := program.instructions[instructionPointer]
		operations[instruction.operation](instruction.a, instruction.b, instruction.c)
		instructionPointer = registers[program.ip]
		instructionPointer++
	}
}

func main() {
	file, err := os.Open("day19/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var program program

	scanner.Scan()
	ipLine := scanner.Text()
	program.ip, _ = strconv.Atoi(ipLine[4:])

	for scanner.Scan() {
		instructionLine := scanner.Text()
		parts := strings.Split(instructionLine, " ")
		a, _ := strconv.Atoi(parts[1])
		b, _ := strconv.Atoi(parts[2])
		c, _ := strconv.Atoi(parts[3])
		program.instructions = append(program.instructions, instruction{parts[0], a, b, c})
	}

	execute(&program)
	fmt.Printf("Part 1: After execution, registers = %v", registers)
}



