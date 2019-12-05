package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("--- Part One ---")
	fmt.Println(Run(read("input.txt"), 1))

	fmt.Println("--- Part Two ---")
	fmt.Println(Run(read("input.txt"), 5))

}

func Run(integers []int, input int) int {

	output := make([]int, 0)

	index := 0
	for {
		instructions := ToArray(integers[index])

		opCode := instructions[0]
		modes := instructions[1:]

		if opCode == 99 {
			return output[len(output)-1]
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
			integers[integers[index+1]] = input
			index += 2
		case 4: // output
			output = append(output, integers[first])
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

func ToArray(index int) []int {
	ints := make([]int, 0)
	ints = append(ints, index%100)        // opcode
	ints = append(ints, (index/100)%10)   // mode 1
	ints = append(ints, (index/1000)%10)  // mode 2
	ints = append(ints, (index/10000)%10) // mode 3
	return ints
}
