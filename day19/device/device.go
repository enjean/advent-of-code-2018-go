package device

import "fmt"

type Device struct {
	Registers []int
}

func New(numRegisters int) *Device {
	return &Device{make([]int, numRegisters)}
}

func (d Device) Execute(program *Program, maxInstructions int) int {
	instructionPointer := 0
	instructionsExecuted := 0
	for instructionsExecuted < maxInstructions && instructionPointer < len(program.Instructions) {
		d.Registers[program.IP] = instructionPointer
		instruction := program.Instructions[instructionPointer]
		fmt.Printf("ip=%d %v %v %d %d %d ", instructionPointer, d.Registers, instruction.operation, instruction.a, instruction.b, instruction.c)
		operations[instruction.operation](d.Registers, instruction.a, instruction.b, instruction.c)
		fmt.Printf("%v\n", d.Registers)
		instructionPointer = d.Registers[program.IP]
		instructionPointer++
		instructionsExecuted++
	}
	return instructionsExecuted
}

var operations = map[string]func(registers []int, a, b, c int){
	"addr": func(registers []int, a, b, c int) {
		registers[c] = registers[a] + registers[b]
	},
	"addi": func(registers []int, a, b, c int) {
		registers[c] = registers[a] + b
	},
	"mulr": func(registers []int, a, b, c int) {
		registers[c] = registers[a] * registers[b]
	},
	"muli": func(registers []int, a, b, c int) {
		registers[c] = registers[a] * b
	},
	"banr": func(registers []int, a, b, c int) {
		registers[c] = registers[a] & registers[b]
	},
	"bani": func(registers []int, a, b, c int) {
		registers[c] = registers[a] & b
	},
	"borr": func(registers []int, a, b, c int) {
		registers[c] = registers[a] | registers[b]
	},
	"bori": func(registers []int, a, b, c int) {
		registers[c] = registers[a] | b
	},
	"setr": func(registers []int, a, b, c int) {
		registers[c] = registers[a]
	},
	"seti": func(registers []int, a, b, c int) {
		registers[c] = a
	},
	"gtir": func(registers []int, a, b, c int) {
		registers[c] = boolToInt(a > registers[b])
	},
	"gtri": func(registers []int, a, b, c int) {
		registers[c] = boolToInt(registers[a] > b)
	},
	"gtrr": func(registers []int, a, b, c int) {
		registers[c] = boolToInt(registers[a] > registers[b])
	},
	"eqir": func(registers []int, a, b, c int) {
		registers[c] = boolToInt(a == registers[b])
	},
	"eqri": func(registers []int, a, b, c int) {
		registers[c] = boolToInt(registers[a] == b)
	},
	"eqrr": func(registers []int, a, b, c int) {
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

