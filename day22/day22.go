package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//part1 := riskLevel(9, 796, 6969)
	//fmt.Printf("Part 1 = %d\n", part1)
	part2 := shortestTimeTo(9, 796, 6969)
	fmt.Printf("Part 2 = %d\n", part2)
}

func erosionLevelTable(maxX, maxY, targetX, targetY, depth int) [][]int {
	table := make([][]int, maxX+1)
	for x := 0; x <= maxX; x++ {
		table[x] = make([]int, maxY+1)
		for y := 0; y <= maxY; y++ {
			var geologicIndex int
			if (x == 0 && y == 0) || (x == targetX && y == targetY) {
				geologicIndex = 0
			} else if y == 0 {
				geologicIndex = x * 16807
			} else if x == 0 {
				geologicIndex = y * 48271
			} else {
				geologicIndex = table[x-1][y] * table[x][y-1]
			}
			table[x][y] = (geologicIndex + depth) % 20183
		}
	}
	return table
}

type RegionType int

const (
	Rocky RegionType = iota
	Wet
	Narrow
)

func (r RegionType) String() string {
	switch r {
	case Rocky:
		return "Rocky"
	case Wet:
		return "Wet"
	case Narrow:
		return "Narrow"
	}
	return fmt.Sprintf("RegionType(%d)", r)
}

func (r RegionType) toolsForRegion() []Tool {
	switch r {
	case Rocky:
		return []Tool{ClimbingGear, Torch}
	case Wet:
		return []Tool{ClimbingGear, Neither}
	case Narrow:
		return []Tool{Torch, Neither}
	}
	panic("Unknown RegionType")
}

func riskLevelTable(maxX, maxY, targetX, targetY, depth int) [][]int {
	erosionLevelTable := erosionLevelTable(maxX, maxY, targetX, targetY, depth)
	riskLevelTable := make([][]int, len(erosionLevelTable))
	for x := 0; x < len(erosionLevelTable); x++ {
		riskLevelTable[x] = make([]int, len(erosionLevelTable[x]))
		for y := 0; y < len(erosionLevelTable[x]); y++ {
			riskLevelTable[x][y] = erosionLevelTable[x][y] % 3
		}
	}
	return riskLevelTable
}

func riskLevel(maxX, maxY, depth int) int {
	riskLevelTable := riskLevelTable(maxX, maxY, maxX, maxY, depth)
	totalRiskLevel := 0
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			totalRiskLevel += riskLevelTable[x][y]
		}
	}
	return totalRiskLevel
}

type Tool int

const (
	Neither Tool = iota
	ClimbingGear
	Torch
)

func (t Tool) String() string {
	switch t {
	case Neither:
		return "Neither"
	case ClimbingGear:
		return "ClimbingGear"
	case Torch:
		return "Torch"
	}
	return fmt.Sprintf("Tool(%d)", t)
}

type key struct {
	x, y int
	tool Tool
}

func (k key) String() string {
	return fmt.Sprintf("(%d, %d)-%v", k.x, k.y, k.tool)
}

type vertex struct {
	key      key
	prev     *vertex
	distance int
	index    int
}

func (v *vertex) String() string {
	return fmt.Sprintf("{%v:%d}", v.key, v.distance)
}

// Implements heap.Interface
type PriorityQueue []*vertex

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*vertex)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type TravelState struct {
	regionQueue *PriorityQueue
	regionMap   map[key]*vertex
	extracted   map[key]bool
}

func newTravelState() TravelState {
	var regionQueue PriorityQueue
	return TravelState{
		regionQueue: &regionQueue,
		regionMap:   make(map[key]*vertex),
		extracted:   make(map[key]bool),
	}
}

func (ts TravelState) add(key key, prev *vertex, distance int) {
	v := vertex{
		key:      key,
		prev:     prev,
		distance: distance,
	}
	heap.Push(ts.regionQueue, &v)
	ts.regionMap[key] = &v
}

func (ts TravelState) extractMin() *vertex {
	//fmt.Printf("Q=%v\n", ts.regionQueue)
	vertex := heap.Pop(ts.regionQueue).(*vertex)
	ts.extracted[vertex.key] = true
	return vertex
}

func (ts TravelState) unvisitedNeighbors(from key, riskLevelTable [][]int) []key {
	possibleCoords := []struct {
		x, y int
	}{
		{from.x - 1, from.y},
		{from.x + 1, from.y},
		{from.x, from.y - 1},
		{from.x, from.y + 1},
	}
	var unvisitedNeighbors []key
	for _, coord := range possibleCoords {
		if coord.x < 0 || coord.y < 0 {
			continue
		}
		toolsToUse := RegionType(riskLevelTable[coord.x][coord.y]).toolsForRegion()
		for _, tool := range toolsToUse {
			if tool == from.tool && !ts.extracted[key{coord.x, coord.y, tool}] {
				unvisitedNeighbors = append(unvisitedNeighbors, key{coord.x, coord.y, tool})
			}
		}
	}
	return unvisitedNeighbors
}

func (ts TravelState) updateDistance(key key, prev *vertex, distance int) {
	vertex, ok := ts.regionMap[key]
	if !ok {
		ts.add(key, prev, distance)
	} else if distance < vertex.distance {
		vertex.distance = distance
		vertex.prev = prev
		heap.Fix(ts.regionQueue, vertex.index)
	}
}

func shortestTimeTo(x, y, depth int) int {
	riskLevelTable := riskLevelTable(x+1000, y+1000, x, y, depth)

	target := key{x, y, Torch}

	travelState := newTravelState()
	travelState.add(key{0, 0, Torch}, nil, 0)

	var found *vertex
	for found == nil {
		vertex := travelState.extractMin()
		fmt.Printf("Considering %v at distance %d \n", vertex.key, vertex.distance)

		if vertex.key == target {
			found = vertex
			break
		}
		for _, t := range RegionType(riskLevelTable[vertex.key.x][vertex.key.y]).toolsForRegion() {
			if t != vertex.key.tool {
				travelState.updateDistance(key{vertex.key.x, vertex.key.y, t}, vertex, vertex.distance+7)
			}
		}

		neighbors := travelState.unvisitedNeighbors(vertex.key, riskLevelTable)

		for _, neighbor := range neighbors {
			travelState.updateDistance(neighbor, vertex, vertex.distance+1)
		}
	}
	distance := found.distance
	if found.key.tool != Torch {
		distance += 7
	}
	v := found
	for v != nil {
		fmt.Printf("->%v %v\n", v, RegionType(riskLevelTable[v.key.x][v.key.y]))
		v = v.prev
	}
	fmt.Printf("\n")
	return distance
}
