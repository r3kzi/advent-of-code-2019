package main

import (
	"fmt"
)

const (
	WHITE = 1
	BLACK = 0
)

var up = V{X: 0, Y: 1}
var down = V{X: 0, Y: -1}
var right = V{X: 1, Y: 0}
var left = V{X: -1, Y: 0}

type V struct {
	X int8
	Y int8
}

func (a *V) Move(b V) {
	a.X += b.X
	a.Y += b.Y
}

func main() {
	program := Read("input.txt")

	fmt.Println("--- Part One ---")
	fmt.Println(len(paint(program, BLACK)))

	grid := paint(program, WHITE)

	var minX, maxX, minY, maxY int8
	for v, _ := range grid {
		minX = Min(v.X, minX)
		minY = Min(v.Y, minY)
		maxX = Max(v.X, maxX)
		maxY = Max(v.Y, maxY)
	}

	fmt.Println("--- Part Two ---")
	for y := maxY; y >= minY; y-- {
		for x := maxX; x >= minX; x-- {
			switch grid[V{X: x, Y: y}] {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}

func paint(program []int64, initial int64) map[V]int64 {
	// input channel has to buffered since we dont have <- input
	input, output, done := make(chan int64, 1), make(chan int64), make(chan bool)
	go Run(program, input, output, done)

	grid := make(map[V]int64)

	position := V{X: 0, Y: 0}
	direction := up

	grid[position] = initial

	for {
		input <- grid[position]

		select {
		case color := <-output:
			// first output indicates color to paint
			switch color {
			case 0:
				grid[position] = BLACK
			case 1:
				grid[position] = WHITE
			default:
				panic(fmt.Sprintf("Shouldn't be here - invalid color %d", color))
			}
			// second output indicates which turn robot will take next
			if turn := <-output; turn == 1 {
				// turn left
				switch direction {
				case up:
					direction = left
				case down:
					direction = right
				case right:
					direction = up
				case left:
					direction = down
				}
			} else {
				// turn right
				switch direction {
				case up:
					direction = right
				case down:
					direction = left
				case right:
					direction = down
				case left:
					direction = up
				}
			}
			position.Move(direction)
		case <-done:
			return grid
		}
	}
}
