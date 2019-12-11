package main

import "fmt"

func Run(program []int64, input chan int64, output chan int64, stop chan bool) {
	integers := make([]int64, len(program))

	copy(integers, program)

	var index, relativeBase int64 = 0, 0

	for {
		instructions := parseInstructionSet(integers[index])

		get := func(offset int64) (parameter int64) {
			switch instructions[offset] {
			case 0:
				parameter = integers[index+offset]
			case 1:
				parameter = index + offset
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

		switch instructions[0] { // check opcode
		case 1: // add
			first, second, third := get(1), get(2), get(3)
			integers[third] = integers[first] + integers[second]
			index += 4
		case 2: // multiply
			first, second, third := get(1), get(2), get(3)
			integers[third] = integers[first] * integers[second]
			index += 4
		case 3: // save input
			first := get(1)
			integers[first] = <-input
			index += 2
		case 4: // output
			first := get(1)
			output <- integers[first]
			index += 2
		case 5: // jump-if-true
			first, second := get(1), get(2)
			if integers[first] != 0 {
				index = integers[second]
			} else {
				index += 3
			}
		case 6: // jump-if-false
			first, second := get(1), get(2)
			if integers[first] == 0 {
				index = integers[second]
			} else {
				index += 3
			}
		case 7: // less than
			first, second, third := get(1), get(2), get(3)
			if integers[first] < integers[second] {
				integers[third] = 1
			} else {
				integers[third] = 0
			}
			index += 4
		case 8: // equals
			first, second, third := get(1), get(2), get(3)
			if integers[first] == integers[second] {
				integers[third] = 1
			} else {
				integers[third] = 0
			}
			index += 4
		case 9: // adjusts the relative base
			first := get(1)
			relativeBase += integers[first]
			index += 2
		case 99: // output
			stop <- true
			return
		default:
			panic(fmt.Sprintf("Shouldn't be here - invalid opcode %d", instructions[0]))
		}
	}
}

func parseInstructionSet(index int64) (ints []int64) {
	ints = append(ints, index%100, index/100%10, index/1000%10, index/10000%10)
	return ints
}
