package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2018-go/day19/device"
	"math"
)

func main() {
	testDevice := device.New(6)
	program := device.Parse("day19/input.txt")
	testDevice.Execute(program, math.MaxInt32)
	fmt.Printf("Part 1: After execution, registers = %v", testDevice.Registers)

	// TODO Need to work this out on paper
	//registers = [6]int{1}
	//
	//execute(&program)
	//fmt.Printf("Part 2: After execution, registers = %v", registers)
}



