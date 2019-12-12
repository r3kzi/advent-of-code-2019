package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Read(fileName string) []Moon {
	file, _ := os.Open(fileName)
	defer file.Close()

	moons := make([]Moon, 0)

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		regex := regexp.MustCompile("[-]*[0-9]+")
		digits := regex.FindAllString(string(line), -1)

		position := Vector{
			x: ToInt(digits[0]),
			y: ToInt(digits[1]),
			z: ToInt(digits[2]),
		}

		velocity := Vector{
			x: 0,
			y: 0,
			z: 0,
		}

		moon := Moon{
			pos:      position,
			velocity: velocity,
		}

		moons = append(moons, moon)
	}
	return moons
}

func ToInt(s string) int {
	atoi, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Shouldn't be here - error parsing string %s", s))
	}
	return atoi
}
