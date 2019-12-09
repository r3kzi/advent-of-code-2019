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

	var index, relativeBase int64 = 0,0

	for {
		instructions := parseInstructionSet(integers[index])

		get := func(offset int64) (parameter int64) {
			switch instructions[offset] {
			case 0:
				parameter = integers[index+offset]
			case 1:
				parameter = index+offset
			case 2:
				parameter = integers[index+offset] + relativeBase
			default:
				panic(fmt.Sprintf("Shouldn't be here - invalid mode %d", instructions[3-offset]))
			}
			for int64(len(integers)) <= parameter {
				integers = append(integers, 0)
			}
			return
		}

		first, second, third := get(1), get(2), get(3)

		switch instructions[0] { // check opcode
		case 99: // output
			return
		case 1: // add
			integers[third] = integers[first] + integers[second]
			index += 4
		case 2: // multiply
			integers[third] = integers[first] * integers[second]
			index += 4
		case 3: // save input
			integers[first] = input
			index += 2
		case 4: // output
			output <- integers[first]
			index += 2
		case 5: // jump-if-true
			if integers[first] != 0 {
				index = integers[second]
			} else {
				index += 3
			}
		case 6: // jump-if-false
			if integers[first] == 0 {
				index = integers[second]
			} else {
				index += 3
			}
		case 7: // less than
			if integers[first] < integers[second] {
				integers[third] = 1
			} else {
				integers[third] = 0
			}
			index += 4
		case 8: // equals
			if integers[first] == integers[second] {
				integers[third] = 1
			} else {
				integers[third] = 0
			}
			index += 4
		case 9: // adjusts the relative base
			relativeBase += integers[first]
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
		i, err := strconv.ParseInt(s,10, 64)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func parseInstructionSet(index int64) (ints []int64) {
	ints = append(ints, index%100, index/100%10, index/1000%10, index/10000%10)
	return ints
}
