package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2018-go/day19/device"
)

func main() {
	program := device.Parse("day21/input.txt")

	fmt.Println(program.ToGo())

	testDevice := device.New(6)
	testDevice.Registers[0] = 15823996
	executions := testDevice.Execute(program, 100000)
	fmt.Println(executions)
}
