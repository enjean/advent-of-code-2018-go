package main

import "fmt"

func main() {
	part1 := riskLevel(9, 796, 6969)
	fmt.Printf("Part 1 = %d\n", part1)
}

func erosionLevelTable(maxX, maxY, depth int) [][]int {
	table := make([][]int, maxX + 1)
	for x := 0; x <= maxX; x++ {
		table[x] = make([]int, maxY + 1)
		for y:= 0; y <= maxY; y++ {
			var geologicIndex int
			if (x==0 && y==0) || (x==maxX && y==maxY) {
				geologicIndex = 0
			} else if y==0 {
				geologicIndex = x * 16807
			} else if x==0 {
				geologicIndex = y * 48271
			} else {
				geologicIndex = table[x-1][y] * table[x][y-1]
			}
			table[x][y] = (geologicIndex + depth) % 20183
		}
	}
	return table
}

func riskLevel(maxX, maxY, depth int) int {
	table := erosionLevelTable(maxX, maxY, depth)
	totalRiskLevel := 0
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			riskLevel := table[x][y] % 3
			totalRiskLevel += riskLevel
		}
	}
	return totalRiskLevel
}