package main

import (
	"fmt"
)

func main() {
	part1, part2 := 0, 0

	for i := 231832; i <= 767346; i++ {
		pw := ToArray(i)
		if HasDouble(pw) && NeverDecrease(pw) {
			part1++
			if checkLargerGroupOfMatches(pw) {
				part2++
			}
		}
	}

	fmt.Println("--- Part One ---")
	fmt.Println(part1)

	fmt.Println("--- Part Two ---")
	fmt.Println(part2)
}

func checkLargerGroupOfMatches(pw []int) bool {
	for i := 0; i < len(pw)-1; i++ {
		if pw[i] == pw[i+1] {
			if (i-1 < 0 || pw[i-1] != pw[i]) && (i+2 > len(pw)-1 || pw[i+2] != pw[i]) {
				return true
			}
		}
	}
	return false
}

func NeverDecrease(pw []int) bool {
	for i := 0; i < len(pw)-1; i++ {
		if pw[i] > pw[i+1] {
			return false
		}
	}
	return true
}

func HasDouble(pw []int) bool {
	for i := 0; i < len(pw)-1; i++ {
		if pw[i] == pw[i+1] {
			return true
		}
	}
	return false
}

func ToArray(pw int) []int {
	ints := make([]int, 0)
	ints = append(ints, (pw/100000)%10)
	ints = append(ints, (pw/10000)%10)
	ints = append(ints, (pw/1000)%10)
	ints = append(ints, (pw/100)%10)
	ints = append(ints, (pw/10)%10)
	ints = append(ints, (pw/1)%10)
	return ints
}
