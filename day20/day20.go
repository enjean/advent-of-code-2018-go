package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Path struct {
	route string
	branches []*Path
}

type position struct {
	x int
	y int
}

//func (rs *RoomStack) New() *RoomStack {
//	rs.rooms = []*Room{}
//	return rs
//}

//func (ps *positionStack) Push(position *position) {
//	ps.positions = append(ps.positions, position)
//}
//
//func (ps *positionStack) Pop() *position {
//	n := len(ps.positions) - 1
//	position := ps.positions[n]
//	ps.positions = ps.positions[:n]
//	return position
//}
//
//func (ps *positionStack) Peek() *position {
//	return ps.positions[len(ps.positions) - 1]
//}

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

func ParsePath(input string) *Path {
	var pathBuilder strings.Builder
	for _, runeVal := range input {
		switch runeVal {
		case '^', '$':
			continue

		default:
			pathBuilder.WriteRune(runeVal)
		}

	}
	return &Path{pathBuilder.String(), nil}
}

func (p *Path) print() string {
	var sb strings.Builder
	sb.WriteString("{")
	sb.WriteString(p.route)
	sb.WriteString(" [")
	for _, branch := range p.branches {
		sb.WriteString(branch.print())
	}
	sb.WriteString("]}")
	return sb.String()
}