package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Point struct{
	w, x, y, z int
}

func (p Point) distance(other Point) int {
	return abs(p.w-other.w) + abs(p.x-other.x) + abs(p.y-other.y) + abs(p.z-other.z)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type constellation struct {
	points []int
}

func CountConstellations(points []Point) int {
	adjacencyList := make(map[int][]int)
	for i := 0; i < len(points) - 1; i++ {
		for j := i + 1; j < len(points); j++ {
			distance := points[i].distance(points[j])
			//fmt.Printf("Distance %d->%d=%d\n", i, j, distance)
			if distance <= 3 {
				adjacencyList[i] = append(adjacencyList[i], j)
				adjacencyList[j] = append(adjacencyList[j], i)
			}
		}
	}

	visited := make(map[int]bool)
	numConstellations := 0
	for i :=0; i < len(points); i++ {
		if visited[i] {
			continue
		}
		//fmt.Printf("Starting at %d\n", i)
		numConstellations++
		queue := []int{i}
		visited[i] = true
		for len(queue) > 0 {
			n := len(queue) - 1
			w := queue[n]
			queue = queue[:n]

			//fmt.Printf("    Includes %d\n", w)
			for _, neighbor := range adjacencyList[w] {
				if visited[neighbor] {
					continue
				}
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return numConstellations
}

func main() {
	file, err := os.Open("day25/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pointExp := regexp.MustCompile(`(-?\d+),(-?\d+),(-?\d+),(-?\d+)`)
	var points []Point
	for scanner.Scan() {
		line := scanner.Text()
		pointMatch := pointExp.FindStringSubmatch(line)
		w, _ := strconv.Atoi(pointMatch[1])
		x, _ := strconv.Atoi(pointMatch[2])
		y, _ := strconv.Atoi(pointMatch[3])
		z, _ := strconv.Atoi(pointMatch[4])
		points = append(points, Point{w,x,y,z})
	}

	numConstellations := CountConstellations(points)
	fmt.Printf("Part 1: Number of constellations = %d\n", numConstellations)
}
