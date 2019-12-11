package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Read(fileName string) []int64 {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	ints := make([]int64, 0)
	for _, s := range strings.Split(string(file), ",") {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}
