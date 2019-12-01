package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day-1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	reader := bufio.NewReader(file)

	var sum = 0
	for {
		mass, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		byteToInt, _ := strconv.Atoi(string(mass))

		fuel := CalculateFuel(byteToInt, 0)

		sum += fuel
	}
	fmt.Println(sum)
}

func CalculateFuel(mass int, total int) int {
	fuel := (mass / 3) - 2
	if fuel <= 0 {
		// negative fuel = zero fuel
		return total
	}
	total += fuel
	return CalculateFuel(fuel, total)
}