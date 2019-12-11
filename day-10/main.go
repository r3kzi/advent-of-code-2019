package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var asteroids []Coordinate

type Coordinate struct {
	x, y int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	asteroids = make([]Coordinate, 0)

	reader := bufio.NewReader(file)
	lineIndex := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		for index, char := range line {
			if char == '#' {
				asteroids = append(asteroids, Coordinate{
					x: lineIndex,
					y: index,
				})
			}
		}
		lineIndex++
	}

	maxCounter, bestLocation := 0, Coordinate{}

	for _, coordinateA := range asteroids {

		counter := 0

		keys := make(map[float64]Coordinate)

		for _, coordinateB := range asteroids {
			if coordinateB != coordinateA {
				angle := getAngle(coordinateA, coordinateB)
				if _, ok := keys[angle]; !ok {
					keys[angle] = coordinateB
					counter++
				}
			}
		}

		if counter > maxCounter {
			maxCounter = counter
			bestLocation = Coordinate{
				x: coordinateA.x,
				y: coordinateA.y,
			}
		}
	}
	fmt.Println(fmt.Sprintf("Found best location in %v with %v Asteriods in sight", bestLocation, maxCounter))
}

func getAngle(a Coordinate, b Coordinate) float64 {
	return math.Atan2(float64(a.y-b.y), float64(a.x-b.x))
}
