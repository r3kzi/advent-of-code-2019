package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	orbits := make(map[string]string)
	for _, or := range read("input.txt") {
		orbits[strings.Split(or, ")")[1]] = strings.Split(or, ")")[0]
	}

	total := 0
	for planet := range orbits {
		for {
			if parent, ok := orbits[planet]; ok {
				planet = parent
				total++
			} else {
				break
			}
		}
	}

	getPathWithSteps := func(planet string) map[string]int {
		planets := make(map[string]int)
		steps := 0
		for {
			if parent, ok := orbits[planet]; ok {
				planets[parent] = steps
				planet = parent
				steps++
			} else {
				break
			}
		}
		return planets
	}

	ch1 := make(chan map[string]int)
	ch2 := make(chan map[string]int)

	go func() {
		ch1 <- getPathWithSteps("YOU")
	}()

	go func() {
		ch2 <- getPathWithSteps("SAN")
	}()

	you := <- ch1
	san := <- ch2

	numberOrbitalTransfers := math.MaxFloat64
	for planet, s1 := range you {
		if s2, ok := san[planet]; ok {
			numberOrbitalTransfers = math.Min(numberOrbitalTransfers, float64(s1+s2))
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(total)

	fmt.Println("--- Part Two ---")
	fmt.Println(numberOrbitalTransfers)
}

func read(fileName string) []string {
	file, _ := ioutil.ReadFile(fileName)
	orbits := make([]string, 0)
	for _, s := range strings.Split(string(file), "\n") {
		orbits = append(orbits, s)
	}
	return orbits
}
