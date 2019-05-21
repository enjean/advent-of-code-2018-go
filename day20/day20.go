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

type distanceMap map[position]int

func main() {
	file, err := os.Open("day20/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	pathRegex := scanner.Text()

	distancesToRooms := distancesToRooms(pathRegex)

	fmt.Printf("Day 20 Part 1: Furthest Room = %d\n", distancesToRooms.max())
	fmt.Printf("Day 20 Part 2: At Least 1000 Away = %d\n", distancesToRooms.atLeastNAway(1000))
}

func (dm *distanceMap) max() int {
	maxDistance := 0
	for _, distanceToRoom := range *dm {
		if distanceToRoom > maxDistance {
			maxDistance = distanceToRoom
		}
	}
	return maxDistance
}

func (dm* distanceMap) atLeastNAway(n int) int {
	count := 0
	for _, distanceToRoom := range *dm {
		if distanceToRoom >= n {
			count++
		}
	}
	return count
}

func distancesToRooms(pathRegex string) *distanceMap{
	roomStack := []*position{}
	currentPosition := &position{0, 0}
	distancesToRooms := make(distanceMap)
	maxDistance := 0
	for _, runeVal := range pathRegex {
		previousDistance := distancesToRooms[*currentPosition]
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

		currentDistance, roomSeen := distancesToRooms[*currentPosition]
		if !roomSeen {
			currentDistance = previousDistance + 1
			distancesToRooms[*currentPosition] = currentDistance
		}
		//fmt.Printf("%c In room %+v distance %d\n", runeVal, currentPosition, currentDistance)
		if currentDistance > maxDistance {
			maxDistance = currentDistance
		}
	}
	return &distancesToRooms
}