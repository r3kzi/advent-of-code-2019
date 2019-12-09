package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")

	positionsA := forWire(strings.Split(string(file), "\n")[0])
	positionsB := forWire(strings.Split(string(file), "\n")[1])

	closestManhattan, closestSteps := math.MaxInt64, math.MaxInt64

	for _, intersection := range intersectionsOf(positionsA, positionsB) {
		closestManhattan = min(closestManhattan, abs(intersection.x, intersection.y))
		closestSteps = min(closestSteps, positionsA[intersection]+positionsB[intersection])
	}

	fmt.Println("--- Part One ---")
	fmt.Println(closestManhattan)

	fmt.Println("--- Part Two ---")
	fmt.Println(closestSteps)
}

func intersectionsOf(positionsA map[Position]int, positionsB map[Position]int) (intersections []Position) {
	intersections = make([]Position, 0)
	for position, _ := range positionsA {
		if _, ok := positionsB[position]; ok {
			intersections = append(intersections, position)
		}
	}
	return
}

func forWire(wire string) (positions map[Position]int) {
	positions = make(map[Position]int)
	position := Position{x: 0, y: 0}

	var steps = 0
	for _, segment := range strings.Split(wire, ",") {
		for i := 0; i < toInt(segment[1:]); i++ {
			switch string(segment[0]) {
			case "R":
				position.x += 1
			case "L":
				position.x -= 1
			case "U":
				position.y += 1
			case "D":
				position.y -= 1
			}
			steps++
			positions[position] = steps
		}
	}
	return
}

func toInt(s string) int {
	atoi, _ := strconv.Atoi(s)
	return atoi
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int, b int) int {
	return int(math.Abs(float64(a)) + math.Abs(float64(b)))
}
