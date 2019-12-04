package main

import (
	"testing"
)

func Test_ToArray(t *testing.T) {
	result := ToArray(122456)
	if len(result) != 6 {
		t.Errorf("Failed")
	}
}

func Test_HasDouble(t *testing.T) {
	result := HasDouble([]int{1, 2, 3, 4, 5, 5})
	if !result {
		t.Errorf("Failed")
	}
}

func Test_NeverDecrease(t *testing.T) {
	result := NeverDecrease([]int{2, 2, 3, 4, 5, 4})
	if result {
		t.Errorf("Failed")
	}
}
