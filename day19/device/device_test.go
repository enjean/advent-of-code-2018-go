package device

import (
	"math"
	"testing"
)

func TestInstructions(t *testing.T) {
	var tests = []struct {
		instruction     string
		registersBefore []int
		args            [3]int
		registersAfter  []int
	}{
		{"addr",[]int{2, 3, 0, 4},  [3]int{0, 1, 3}, []int{2, 3, 0, 5}},
		{"addi",[]int{2, 1, 0, 4},  [3]int{0, 7, 3}, []int{2, 1, 0, 9}},
		{"mulr",[]int{2, 3, 0, 4},  [3]int{0, 1, 3}, []int{2, 3, 0, 6}},
		{"muli",[]int{2, 1, 0, 4},  [3]int{0, 7, 3}, []int{2, 1, 0, 14}},
		{"banr",[]int{2, 3, 0, 4},  [3]int{0, 1, 3}, []int{2, 3, 0, 2}},
		{"bani",[]int{9, 1, 0, 4},  [3]int{0, 7, 3}, []int{9, 1, 0, 1}},
		{"borr",[]int{2, 3, 0, 4},  [3]int{0, 1, 3}, []int{2, 3, 0, 3}},
		{"bori",[]int{9, 1, 0, 4},  [3]int{0, 7, 3}, []int{9, 1, 0, 15}},
		{"setr",[]int{2, 3, 0, 4},  [3]int{0, 1, 3}, []int{2, 3, 0, 2}},
		{"seti",[]int{9, 1, 0, 4},  [3]int{0, 7, 3}, []int{9, 1, 0, 0}},
		{"gtir",[]int{9, 3, 0, 4},  [3]int{0, 1, 3}, []int{9, 3, 0, 0}},
		{"gtri",[]int{2, 3, 0, 4},  [3]int{0, 1, 3}, []int{2, 3, 0, 1}},
		{"gtrr",[]int{2, 3, 0, 4},  [3]int{0, 1, 3}, []int{2, 3, 0, 0}},
		{"gtrr",[]int{3, 2, 0, 4},  [3]int{0, 1, 3}, []int{3, 2, 0, 1}},
		{"eqir",[]int{9, 3, 0, 4},  [3]int{0, 1, 3}, []int{9, 3, 0, 0}},
		{"eqir",[]int{9, 3, 0, 4},  [3]int{3, 1, 3}, []int{9, 3, 0, 1}},
		{"eqri",[]int{9, 3, 0, 4},  [3]int{0, 1, 3}, []int{9, 3, 0, 0}},
		{"eqri",[]int{9, 3, 0, 4},  [3]int{0, 9, 3}, []int{9, 3, 0, 1}},
		{"eqrr",[]int{9, 3, 0, 4},  [3]int{0, 2, 3}, []int{9, 3, 0, 0}},
		{"eqrr",[]int{0, 3, 0, 4},  [3]int{0, 2, 3}, []int{0, 3, 0, 1}},
	}
	for _, test := range tests {
		registers := make([]int, len(test.registersBefore))
		copy(registers, test.registersBefore)
		operations[test.instruction](registers, test.args[0], test.args[1], test.args[2])
		if !equal(registers, test.registersAfter) {
			t.Errorf("%q(%v, %v) = %v", test.instruction, test.registersBefore, test.args, test.registersAfter)
		}
	}
}

func TestExecute(t *testing.T) {
	program := Program{IP: 0,
		Instructions: []Instruction{
			{"seti", 5, 0, 1},
			{"seti", 6, 0, 2},
			{"addi", 0, 1, 0},
			{"addr", 1, 2, 3},
			{"setr", 1, 0, 0},
			{"seti", 8, 0, 4},
			{"seti", 9, 0, 5},
		},
	}

	testDevice := New(6)
	testDevice.Execute(&program, math.MaxInt32)
	if !equal(testDevice.Registers, []int{6, 5, 6, 0, 0, 9}) {
		t.Errorf("Final registers %v not equal to expected", testDevice.Registers)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}