package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	program := toIntegers(file)

	/* Part 1 */
	producedOutput, _, _ := Run(program, 12, 2)
	fmt.Printf("answer for part 1 is %d\n", producedOutput)

	/* Part 2 */
	noun, verb := FindNounVerb(19690720, program)
	fmt.Printf("answer for part 2 is %d\n", 100*noun+verb)
}

func FindNounVerb(output int, program []int) (int, int) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			producedOutput, noun, verb := Run(program, noun, verb)
			if producedOutput == output {
				return noun, verb
			}
		}
	}
	return 0, 0
}

func Run(program []int, noun int, verb int) (int, int, int) {
	integers := make([]int, len(program))
	copy(integers, program)

	integers[1] = noun
	integers[2] = verb

	for index, _ := range integers {
		// check opcode
		if index%4 == 0 {
			switch integers[index] {
			case 99:
				break
			case 1:
				output := integers[index+3]
				integers[output] = integers[integers[index+1]] + integers[integers[index+2]]
			case 2:
				output := integers[index+3]
				integers[output] = integers[integers[index+1]] * integers[integers[index+2]]
			}
		} else {
			continue
		}
	}
	// return output, noun, verb
	return integers[0], integers[1], integers[2]
}

func toIntegers(file []byte) []int {
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
