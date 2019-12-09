package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	program := read("input.txt")

	output := make(chan int64)

	go Run(program, int64(1), output)
	go Run(program, int64(2), output)

	for i := 0; i < 2; i++ {
		fmt.Printf("--- Part %d ---\n", i+1)
		fmt.Println(<-output)
	}
	close(output)
}

func Run(program []int64, input int64, output chan int64) {
	integers := make([]int64, len(program))
	copy(integers, program)

	var index, relativeBase int64

	index = 0
	relativeBase = 0

	appendIfNecessary := func(address int64) *int64 {
		for int64(len(integers)) <= address {
			integers = append(integers, 0)
		}
		return &integers[address]
	}

	getParameterForMode := func(offset int64, mode int64) *int64 {
		parameter := integers[index+offset]
		switch mode {
		case 0:
			return appendIfNecessary(parameter)
		case 1:
			return &parameter
		case 2:
			return appendIfNecessary(parameter + relativeBase)
		default:
			panic(fmt.Sprintf("Shouldn't be here - invalid mode %d", mode))
		}
	}

	for {
		instructions := ToArray(integers[index])

		switch instructions[0] { // check opcode
		case 99: // output
			return
		case 1: // add
			*getParameterForMode(3, instructions[3]) = *getParameterForMode(1, instructions[1]) + *getParameterForMode(2, instructions[2])
			index += 4
		case 2: // multiply
			*getParameterForMode(3, instructions[3]) = *getParameterForMode(1, instructions[1]) * *getParameterForMode(2, instructions[2])
			index += 4
		case 3: // save input
			*getParameterForMode(1, instructions[1]) = input
			index += 2
		case 4: // output
			output <- *getParameterForMode(1, instructions[1])
			index += 2
		case 5: // jump-if-true
			if *getParameterForMode(1, instructions[1]) != 0 {
				index = *getParameterForMode(2, instructions[2])
			} else {
				index += 3
			}
		case 6: // jump-if-false
			if *getParameterForMode(1, instructions[1]) == 0 {
				index = *getParameterForMode(2, instructions[2])
			} else {
				index += 3
			}
		case 7: // less than
			if *getParameterForMode(1, instructions[1]) < *getParameterForMode(2, instructions[2]) {
				*getParameterForMode(3, instructions[3]) = 1
			} else {
				*getParameterForMode(3, instructions[3]) = 0
			}
			index += 4
		case 8: // equals
			if *getParameterForMode(1, instructions[1]) == *getParameterForMode(2, instructions[2]) {
				*getParameterForMode(3, instructions[3]) = 1
			} else {
				*getParameterForMode(3, instructions[3]) = 0
			}
			index += 4
		case 9: // adjusts the relative base
			relativeBase += *getParameterForMode(1, instructions[1])
			index += 2
		default:
			panic(fmt.Sprintf("Shouldn't be here - invalid opcode %d", instructions[0]))
		}
	}
}

func read(fileName string) []int64 {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	ints := make([]int64, 0)
	for _, s := range strings.Split(string(file), ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints = append(ints, int64(i))
	}
	return ints
}

func ToArray(index int64) []int64 {
	ints := make([]int64, 0)
	ints = append(ints, index%100)      // opcode
	ints = append(ints, index/100%10)   // mode 1
	ints = append(ints, index/1000%10)  // mode 2
	ints = append(ints, index/10000%10) // mode 3
	return ints
}
