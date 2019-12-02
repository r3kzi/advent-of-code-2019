package main

import (
	"reflect"
	"testing"
)

func TestToInteger(t *testing.T) {
	testCases := []struct {
		x []byte
		y []int
	}{
		{[]byte("1,0,0,0,99"), []int{1, 0, 0, 0, 99}},
		{[]byte("2,3,0,3,99"), []int{2, 3, 0, 3, 99}},
		{[]byte("2,4,4,5,99,0"), []int{2, 4, 4, 5, 99, 0}},
		{[]byte("1,1,1,4,99,5,6,0,99"), []int{1, 1, 1, 4, 99, 5, 6, 0, 99}},
	}

	for _, testCase := range testCases {
		result := toIntegers(testCase.x)
		if !reflect.DeepEqual(result, testCase.y) {
			t.Errorf("Failed - was %v - should be %v", result, testCase.y)
		}
	}
}

func TestRun(t *testing.T) {
	testCases := []struct {
		x []int
		y []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1}},
	}

	for _, testCase := range testCases {
		producedOutput, noun, verb := Run(testCase.x, testCase.x[1], testCase.x[2])
		if producedOutput != testCase.y[0] || noun != testCase.y[1] || verb != testCase.y[2] {
			t.Errorf("Failed - was %v,%v,%v - should be %v,%v,%v", producedOutput, noun, verb, testCase.y[0], testCase.y[1], testCase.y[2])
		}
	}
}
