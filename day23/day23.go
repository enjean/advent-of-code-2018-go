package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Position struct {
	x, y, z int64
}

func (p Position) distanceTo(other Position) int64 {
	return abs(p.x - other.x) + abs(p.y - other.y) + abs(p.z - other.z)
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

type Nanobot struct {
	position Position
	radius int64
}

var lineExp = regexp.MustCompile(`pos=<(-?\d+),(-?\d+),(-?\d+)>, r=(\d+)`)
func Parse(input string) Nanobot {
	matches := lineExp.FindStringSubmatch(input)
	x, _ := strconv.ParseInt(matches[1], 10, 64)
	y, _ := strconv.ParseInt(matches[2], 10, 64)
	z, _ := strconv.ParseInt(matches[3], 10, 64)
	r, _ := strconv.ParseInt(matches[4], 10, 64)
	return Nanobot{Position{x, y, z}, r}
}

func inRangeOfStrongest(bots []Nanobot) int64 {
	var strongest Nanobot
	var maxRadius int64
	for _, bot := range bots {
		if bot.radius > maxRadius {
			strongest = bot
			maxRadius = bot.radius
		}
	}
	var inRange int64
	for _, bot := range bots {
		if strongest.position.distanceTo(bot.position) <= strongest.radius {
			inRange++
		}
	}
	return inRange
}

func main() {
	file, err := os.Open("day23/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bots []Nanobot
	for scanner.Scan() {
		line := scanner.Text()
		bot := Parse(line)
		bots = append(bots, bot)
	}

	part1 := inRangeOfStrongest(bots)
	fmt.Printf("Part 1: %d\n", part1)
}
