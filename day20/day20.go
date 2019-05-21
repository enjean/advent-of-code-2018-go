package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type position struct {
	x int
	y int
}

func main() {
	file, err := os.Open("day20/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	pathRegex := scanner.Text()

	furthestRoom := FurthestRoom(pathRegex)
	fmt.Printf("Day 20 Part 1: Furthest Room = %d", furthestRoom)
}

func FurthestRoom(pathRegex string) int {
	roomStack := []*position{}
	currentPosition := &position{0, 0}
	distances := make(map[position]int)
	maxDistance := 0
	for _, runeVal := range pathRegex {
		previousDistance := distances[*currentPosition]
		switch runeVal {
		case '^', '$':
			continue
		case 'N' :
			currentPosition = &position{currentPosition.x, currentPosition.y-1}
		case 'S' :
			currentPosition = &position{currentPosition.x, currentPosition.y+1}
		case 'W' :
			currentPosition = &position{currentPosition.x-1, currentPosition.y}
		case 'E' :
			currentPosition = &position{currentPosition.x+1, currentPosition.y}
		case '(':
			roomStack = append(roomStack, currentPosition)
		case '|':
			currentPosition = roomStack[len(roomStack) - 1]
		case ')':
			n := len(roomStack) - 1
			currentPosition = roomStack[n]
			roomStack = roomStack[:n]
		}

		currentDistance, roomSeen := distances[*currentPosition]
		if !roomSeen {
			currentDistance = previousDistance + 1
			distances[*currentPosition] = currentDistance
		}
		//fmt.Printf("%c In room %+v distance %d\n", runeVal, currentPosition, currentDistance)
		if currentDistance > maxDistance {
			maxDistance = currentDistance
		}
	}
	return maxDistance
}