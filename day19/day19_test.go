package main

import "testing"

func TestInstructions(t *testing.T) {
	var tests = []struct {
		instruction     string
		registersBefore [6]int
		args            [3]int
		registersAfter  [6]int
	}{
		{"addr",[6]int{2, 3, 0, 4},  [3]int{0, 1, 3}, [6]int{2, 3, 0, 5}},
		{"addi",[6]int{2, 1, 0, 4},  [3]int{0, 7, 3}, [6]int{2, 1, 0, 9}},
		{"mulr",[6]int{2, 3, 0, 4},  [3]int{0, 1, 3}, [6]int{2, 3, 0, 6}},
		{"muli",[6]int{2, 1, 0, 4},  [3]int{0, 7, 3}, [6]int{2, 1, 0, 14}},
		{"banr",[6]int{2, 3, 0, 4},  [3]int{0, 1, 3}, [6]int{2, 3, 0, 2}},
		{"bani",[6]int{9, 1, 0, 4},  [3]int{0, 7, 3}, [6]int{9, 1, 0, 1}},
		{"borr",[6]int{2, 3, 0, 4},  [3]int{0, 1, 3}, [6]int{2, 3, 0, 3}},
		{"bori",[6]int{9, 1, 0, 4},  [3]int{0, 7, 3}, [6]int{9, 1, 0, 15}},
		{"setr",[6]int{2, 3, 0, 4},  [3]int{0, 1, 3}, [6]int{2, 3, 0, 2}},
		{"seti",[6]int{9, 1, 0, 4},  [3]int{0, 7, 3}, [6]int{9, 1, 0, 0}},
		{"gtir",[6]int{9, 3, 0, 4},  [3]int{0, 1, 3}, [6]int{9, 3, 0, 0}},
		{"gtri",[6]int{2, 3, 0, 4},  [3]int{0, 1, 3}, [6]int{2, 3, 0, 1}},
		{"gtrr",[6]int{2, 3, 0, 4},  [3]int{0, 1, 3}, [6]int{2, 3, 0, 0}},
		{"gtrr",[6]int{3, 2, 0, 4},  [3]int{0, 1, 3}, [6]int{3, 2, 0, 1}},
		{"eqir",[6]int{9, 3, 0, 4},  [3]int{0, 1, 3}, [6]int{9, 3, 0, 0}},
		{"eqir",[6]int{9, 3, 0, 4},  [3]int{3, 1, 3}, [6]int{9, 3, 0, 1}},
		{"eqri",[6]int{9, 3, 0, 4},  [3]int{0, 1, 3}, [6]int{9, 3, 0, 0}},
		{"eqri",[6]int{9, 3, 0, 4},  [3]int{0, 9, 3}, [6]int{9, 3, 0, 1}},
		{"eqrr",[6]int{9, 3, 0, 4},  [3]int{0, 2, 3}, [6]int{9, 3, 0, 0}},
		{"eqrr",[6]int{0, 3, 0, 4},  [3]int{0, 2, 3}, [6]int{0, 3, 0, 1}},
	}

	for _, test := range tests {
		registers = test.registersBefore
		operations[test.instruction](test.args[0], test.args[1], test.args[2])
		if registers != test.registersAfter {
			t.Errorf("%q(%v, %v) = %v", test.instruction, test.registersBefore, test.args, test.registersAfter)
		}
	}
}

func TestExecute(t *testing.T) {
	program := program{ip: 0,
		instructions: []instruction{
			{"seti", 5, 0, 1},
			{"seti", 6, 0, 2},
			{"addi", 0, 1, 0},
			{"addr", 1, 2, 3},
			{"setr", 1, 0, 0},
			{"seti", 8, 0, 4},
			{"seti", 9, 0, 5},
		},
	}

	execute(&program)
	if registers != [6]int{6, 5, 6, 0, 0, 9} {
		t.Errorf("Final registers %v not equal to expected", registers)
	}
}