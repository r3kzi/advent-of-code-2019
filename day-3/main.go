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
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	posA := getPosForWire(strings.Split(string(file), "\n")[0])
	posB := getPosForWire(strings.Split(string(file), "\n")[1])

	intersections := getIntersections(posA, posB)

	closestManhattan, closestSteps := math.MaxFloat64, math.MaxFloat64
	for _, intersection := range intersections {
		closestManhattan = math.Min(closestManhattan, calculateManhattan(intersection))
		closestSteps = math.Min(closestSteps, float64(posA[intersection]+posB[intersection]))
	}
	fmt.Println(closestManhattan)
	fmt.Println(closestSteps)
}

func calculateManhattan(pos Position) float64 {
	return math.Abs(float64(pos.x)) + math.Abs(float64(pos.y))
}

func getIntersections(posA map[Position]int, posB map[Position]int) []Position {
	intersections := make([]Position, 0)
	for keyA, _ := range posA {
		for keyB, _ := range posB {
			if keyA == keyB {
				intersections = append(intersections, keyA)
			}
		}
	}

	return intersections
}

func getPosForWire(wire string) map[Position]int {
	positions := make(map[Position]int)

	position := Position{
		x: 0,
		y: 0,
	}

	var steps = 0
	for _, segment := range strings.Split(wire, ",") {
		direction := string(segment[0])
		for i := 0; i < toInt(segment[1:]); i++ {
			switch direction {
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
	return positions
}

func toInt(s string) int {
	atoi, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return atoi
}