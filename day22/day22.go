package main

import (
	"container/heap"
	"fmt"
)

const Infinity = int(^uint(0) >> 1)

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

func toolsForRegion(regionType RegionType) []Tool {
	switch regionType {
	case Rocky:
		return []Tool{ClimbingGear, Torch}
	case Wet:
		return []Tool{ClimbingGear, Neither}
	case Narrow:
		return []Tool{Torch, Neither}
	}
	panic("Unknown RegionType")
}

type regionWithTool struct {
	x, y int
	tool Tool
}

func (k regionWithTool) String() string {
	return fmt.Sprintf("(%d, %d)-%v", k.x, k.y, k.tool)
}

type vertex struct {
	regionWithTool regionWithTool
	prev           *vertex
	distance       int
	index          int
}

func (v *vertex) String() string {
	return fmt.Sprintf("{%v:%d}", v.regionWithTool, v.distance)
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
	regionMap   map[regionWithTool]*vertex
	extracted   map[regionWithTool]bool
}

func newTravelState() TravelState {
	var regionQueue PriorityQueue
	return TravelState{
		regionQueue: &regionQueue,
		regionMap:   make(map[regionWithTool]*vertex),
		extracted:   make(map[regionWithTool]bool),
	}
}

func (ts TravelState) add(key regionWithTool, prev *vertex, distance int) {
	v := vertex{
		regionWithTool: key,
		prev:           prev,
		distance:       distance,
	}
	heap.Push(ts.regionQueue, &v)
	ts.regionMap[key] = &v
}

func (ts TravelState) extractMin() *vertex {
	//fmt.Printf("Q=%v\n", ts.regionQueue)
	vertex := heap.Pop(ts.regionQueue).(*vertex)
	ts.extracted[vertex.regionWithTool] = true
	return vertex
}

func (ts TravelState) unvisitedNeighbors(x, y int, riskLevelTable [][]int) []regionWithTool {
	possibleCoords := []struct {
		x, y int
	}{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}
	var unvisitedNeighbors []regionWithTool
	for _, coord := range possibleCoords {
		if coord.x < 0 || coord.y < 0 {
			continue
		}
		toolsToUse := toolsForRegion(RegionType(riskLevelTable[coord.x][coord.y]))
		for _, tool := range toolsToUse {
			if !ts.extracted[regionWithTool{coord.x, coord.y, tool}] {
				unvisitedNeighbors = append(unvisitedNeighbors, regionWithTool{coord.x, coord.y, tool})
			}
		}
	}
	return unvisitedNeighbors
}

func (ts TravelState) updateDistance(key regionWithTool, prev *vertex, distance int) {
	vertex, ok := ts.regionMap[key]
	if !ok {
		ts.add(key, prev, distance)
	} else if distance < vertex.distance {
		vertex.distance = distance
		vertex.prev = prev
		heap.Fix(ts.regionQueue, vertex.index)
	}
}

func canSwitchTo(regionType RegionType, tool Tool) bool {
	switch regionType {
	case Rocky:
		return tool == ClimbingGear || tool == Torch
	case Wet:
		return tool == ClimbingGear || tool == Neither
	case Narrow:
		return tool == Neither || tool == Torch
	}
	return false
}

func travelTime(currentRegion RegionType, source, target Tool) (int, bool) {
	if source == target {
		return 1, true
	} else if canSwitchTo(currentRegion, target){
		return 8, true
	} else {
		return Infinity, false
	}
}

func shortestTimeTo(x, y, depth int) int {
	riskLevelTable := riskLevelTable(x+1000, y+1000, x, y, depth)

	travelState := newTravelState()
	travelState.add(regionWithTool{0, 0, Torch}, nil, 0)

	//start := vertex{
	//	regionWithTool: regionWithTool{x: 0, y: 0, tool:Torch},
	//	prev: nil,
	//	distance: 0,
	//}
	//heap.Push(&regionQueue, &start)
	//regionMap[]

	var found *vertex
	for found == nil {
		vertex := travelState.extractMin()
		fmt.Printf("Considering %v at distance %d \n", vertex.regionWithTool, vertex.distance)
		if vertex.regionWithTool.x > x * 5 || vertex.regionWithTool.y > y * 5 {
			continue
		}
		if vertex.regionWithTool.x == x && vertex.regionWithTool.y == y {
			found = vertex
			break
		}
		neighbors := travelState.unvisitedNeighbors(vertex.regionWithTool.x, vertex.regionWithTool.y, riskLevelTable)

		for _, neighbor := range neighbors {
			distance, canTravel := travelTime(RegionType(riskLevelTable[vertex.regionWithTool.x][vertex.regionWithTool.y]),
				vertex.regionWithTool.tool, neighbor.tool)
			if canTravel {
				travelState.updateDistance(neighbor, vertex, vertex.distance + distance)
			}
		}
	}
	distance := found.distance
	if found.regionWithTool.tool != Torch {
		distance += 7
	}
	v := found
	for v != nil {
		fmt.Printf("->%v %v\n", v, RegionType(riskLevelTable[v.regionWithTool.x][v.regionWithTool.y]))
		v = v.prev
	}
	fmt.Printf("\n")
	return distance
}
