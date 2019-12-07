package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	permutations "github.com/dbyio/heappermutations"
)

func main() {
	program := read("input.txt")

	part1 := 0
	for _, phaseSetting := range permutations.Ints([]int{0, 1, 2, 3, 4}) {
		part1 = max(part1, testPhaseSetting(phaseSetting, program))
	}

	part2 := 0
	for _, phaseSetting := range permutations.Ints([]int{5, 6, 7, 8, 9}) {
		part2 = max(part2, testPhaseSetting(phaseSetting, program))
	}

	fmt.Println("--- Part One ---")
	fmt.Println(part1)

	fmt.Println("--- Part One ---")
	fmt.Println(part2)
}

func testPhaseSetting(phaseSetting []int, program []int) int {
	stop := make(chan bool)

	xa := make(chan int)
	ab := make(chan int)
	bc := make(chan int)
	cd := make(chan int)
	dx := make(chan int)

	go Run(program, xa, ab, stop)
	go Run(program, ab, bc, stop)
	go Run(program, bc, cd, stop)
	go Run(program, cd, dx, stop)
	go Run(program, dx, xa, stop)

	xa <- phaseSetting[0]
	ab <- phaseSetting[1]
	bc <- phaseSetting[2]
	cd <- phaseSetting[3]
	dx <- phaseSetting[4]

	xa <- 0

	for i := 0; i < 5; i++ {
		<-stop
	}

	return <- xa
}


func Run(program []int, input chan int, output chan int, stop chan bool) {
	integers := make([]int, len(program))
	copy(integers, program)

	integers = append(integers, 0,0 )

	index := 0
	for {
		instructions := toArray(integers[index])

		opCode := instructions[0]
		modes := instructions[1:]

		if opCode == 99 {
			stop <- true
		}

		first := integers[index+1]
		second := integers[index+2]
		// Parameters that an instruction writes to will never be in immediate mode.
		if opCode != 4 && opCode != 3 {
			if modes[0] == 0 {
				first = integers[first]
			}
			if modes[1] == 0 {
				second = integers[second]
			}
		}

		switch opCode {
		case 1: // add
			integers[integers[index+3]] = first + second
			index += 4
		case 2: // multiply
			integers[integers[index+3]] = first * second
			index += 4
		case 3: // save input
			integers[integers[index+1]] = <- input
			index += 2
		case 4: // output
			output <- integers[first]
			index += 2
		case 5: // jump-if-true
			index += 3
			if first != 0 {
				index = second
			}
		case 6: // jump-if-false
			index += 3
			if first == 0 {
				index = second
			}
		case 7: // less than
			if first < second {
				integers[integers[index+3]] = 1
			} else {
				integers[integers[index+3]] = 0
			}
			index += 4
		case 8: // equals
			if first == second {
				integers[integers[index+3]] = 1
			} else {
				integers[integers[index+3]] = 0
			}
			index += 4
		}
	}
}

func read(fileName string) []int {
	file, _ := ioutil.ReadFile(fileName)
	ints := make([]int, 0)
	for _, s := range strings.Split(string(file), ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func toArray(index int) []int {
	ints := make([]int, 0)
	ints = append(ints, index%100)        // opcode
	ints = append(ints, (index/100)%10)   // mode 1
	ints = append(ints, (index/1000)%10)  // mode 2
	ints = append(ints, (index/10000)%10) // mode 3
	return ints
}

func max(a int, b int) int {
	if a > b {
		return  a
	}
	return b
}
