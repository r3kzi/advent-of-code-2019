package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

type Layer struct {
	grid [6][25]string
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")

	width, height := 25, 6
	layers := make([]Layer, 0)

	index := 0
	for {
		if index == len(file) {
			break
		}
		layer := Layer{}
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				layer.grid[i][j] = string(file[index])
				index++
			}
		}
		layers = append(layers, layer)
	}

	maxZeroDigits, part1 := math.MaxInt64, 0
	for _, layer := range layers {
		zeroDigits, oneDigits, twoDigits := 0, 0, 0
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				switch layer.grid[i][j] {
				case "0":
					zeroDigits++
				case "1":
					oneDigits++
				case "2":
					twoDigits++
				}
			}
		}
		if zeroDigits < maxZeroDigits {
			maxZeroDigits = zeroDigits
			part1 = oneDigits * twoDigits
		}
	}
	fmt.Println("--- Part One ---")
	fmt.Println(part1)

	var image [6][25]string
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			image[i][j] = "2"
		}
	}
	for _, layer := range layers {
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if image[i][j] == "2" {
					image[i][j] = layer.grid[i][j]
				}
			}
		}
	}

	fmt.Println("--- Part Two ---")

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			switch image[i][j] {
			case "0":
				fmt.Print(" ")
			case "1":
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}
