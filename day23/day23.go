package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Position struct {
	x, y, z int64
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.x, p.y, p.z)
}

func (p Position) distanceTo(other Position) int64 {
	return abs(p.x-other.x) + abs(p.y-other.y) + abs(p.z-other.z)
}

func (p Position) distanceToOrigin() int64 {
	return p.distanceTo(Position{0, 0, 0})
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

type Nanobot struct {
	position Position
	radius   int64
}

func (bot Nanobot) inRange() *[]Position {
	var positions []Position
	for x := int64(0); x <= bot.radius; x++ {
		for y := int64(0); y <= bot.radius-x; y++ {
			for z := int64(0); z <= bot.radius-x-y; z++ {
				positions = append(positions, Position{bot.position.x + x, bot.position.y + y, bot.position.z + z})
				positions = append(positions, Position{bot.position.x - x, bot.position.y - y, bot.position.z - z})
			}
		}
	}
	return &positions
}

func (bot Nanobot) minDistanceToOrigin() int64 {
	return abs(bot.position.x) + abs(bot.position.y) + abs(bot.position.z) - bot.radius
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

type searchSquare struct {
	bottomLeftCorner Position
	sideLength       int64
}

func (ss searchSquare) String() string {
	return fmt.Sprintf("Square{%v}:%d", ss.bottomLeftCorner, ss.sideLength)
}

type searchSquareItem struct {
	searchSquare searchSquare
	numInRange   int
	index        int
}

func (ssi searchSquareItem) String() string {
	return fmt.Sprintf("Item{%v}:%d", ssi.searchSquare, ssi.numInRange)
}

type SearchQueue []*searchSquareItem

func (pq SearchQueue) Len() int { return len(pq) }

func (pq SearchQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].numInRange != pq[j].numInRange {
		return pq[i].numInRange > pq[j].numInRange
	}
	iDistance := pq[i].searchSquare.bottomLeftCorner.distanceToOrigin()
	jDistance := pq[j].searchSquare.bottomLeftCorner.distanceToOrigin()
	if iDistance != jDistance {
		return iDistance < jDistance
	}
	return pq[i].searchSquare.sideLength < pq[j].searchSquare.sideLength
}

func (pq SearchQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *SearchQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*searchSquareItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *SearchQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func rangeDist(val, low, high int64) int64 {
	if val < low {
		return low - val
	}
	if val > high {
		return val - high
	}
	return 0
}

func (ss searchSquare) inRange(bots []Nanobot) int {
	boxX0 := ss.bottomLeftCorner.x
	boxX1 := ss.bottomLeftCorner.x + ss.sideLength - 1
	boxY0 := ss.bottomLeftCorner.y
	boxY1 := ss.bottomLeftCorner.y + ss.sideLength - 1
	boxZ0 := ss.bottomLeftCorner.z
	boxZ1 := ss.bottomLeftCorner.z + ss.sideLength - 1
	inRange := 0
	for _, bot := range bots {
		distance := rangeDist(bot.position.x, boxX0, boxX1) + rangeDist(bot.position.y, boxY0, boxY1) +
			rangeDist(bot.position.z, boxZ0, boxZ1)
		if distance <= bot.radius {
			inRange++
		}
	}
	return inRange
}

func closestOverlapSearch(bots []Nanobot) int64 {
	maxX := bots[0].position.x + bots[0].radius
	minX := bots[0].position.x - bots[0].radius
	maxY := bots[0].position.y + bots[0].radius
	minY := bots[0].position.y - bots[0].radius
	maxZ := bots[0].position.z + bots[0].radius
	minZ := bots[0].position.z - bots[0].radius

	for _, bot := range bots[1:] {
		highX := bot.position.x + bot.radius
		if highX > maxX {
			maxX = highX
		}
		lowX := bot.position.x - bot.radius
		if lowX < minX {
			minX = lowX
		}
		highY := bot.position.y + bot.radius
		if highY > maxY {
			maxY = highY
		}
		lowY := bot.position.y - bot.radius
		if lowY < minY {
			minY = lowY
		}
		highZ := bot.position.z + bot.radius
		if highZ > maxZ {
			maxZ = highZ
		}
		lowZ := bot.position.z - bot.radius
		if lowZ < minZ {
			minZ = lowZ
		}
	}

	initialLength := int64(1)
	for minX+initialLength < maxX || minY+initialLength < maxY || minZ+initialLength < maxZ {
		initialLength *= 2
	}
	initialItem := searchSquareItem{
		searchSquare: searchSquare{
			bottomLeftCorner: Position{minX, minY, minZ},
			sideLength:       initialLength,
		},
		numInRange: len(bots),
	}

	var searchQueue SearchQueue
	heap.Push(&searchQueue, &initialItem)

	for {
		currentItem := heap.Pop(&searchQueue).(*searchSquareItem)
		currentSquare := currentItem.searchSquare

		if currentSquare.sideLength == 1 {
			return currentSquare.bottomLeftCorner.distanceToOrigin()
		}
		newSideLength := currentSquare.sideLength / 2
		newCoords := []Position{
			{currentSquare.bottomLeftCorner.x, currentSquare.bottomLeftCorner.y, currentSquare.bottomLeftCorner.z},
			{currentSquare.bottomLeftCorner.x, currentSquare.bottomLeftCorner.y, currentSquare.bottomLeftCorner.z + newSideLength},
			{currentSquare.bottomLeftCorner.x, currentSquare.bottomLeftCorner.y + newSideLength, currentSquare.bottomLeftCorner.z},
			{currentSquare.bottomLeftCorner.x, currentSquare.bottomLeftCorner.y + newSideLength, currentSquare.bottomLeftCorner.z + newSideLength},
			{currentSquare.bottomLeftCorner.x + newSideLength, currentSquare.bottomLeftCorner.y, currentSquare.bottomLeftCorner.z},
			{currentSquare.bottomLeftCorner.x + newSideLength, currentSquare.bottomLeftCorner.y, currentSquare.bottomLeftCorner.z + newSideLength},
			{currentSquare.bottomLeftCorner.x + newSideLength, currentSquare.bottomLeftCorner.y + newSideLength, currentSquare.bottomLeftCorner.z},
			{currentSquare.bottomLeftCorner.x + newSideLength, currentSquare.bottomLeftCorner.y + newSideLength, currentSquare.bottomLeftCorner.z + newSideLength},
		}
		for _, newCoord := range newCoords {
			newSquare := searchSquare{
				bottomLeftCorner: Position{newCoord.x, newCoord.y, newCoord.z},
				sideLength:       newSideLength,
			}
			heap.Push(&searchQueue, &searchSquareItem{
				searchSquare: newSquare,
				numInRange:   newSquare.inRange(bots),
			})
		}
	}
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

	//part1 := inRangeOfStrongest(bots)
	//fmt.Printf("Part 1: %d\n", part1)

	closest := closestOverlapSearch(bots)
	fmt.Printf("Part 2: %v\n", closest)
}
