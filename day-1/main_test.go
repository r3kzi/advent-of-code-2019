package main

import (
	"testing"
)

func TestFor1969(t *testing.T) {
	mass := 1969
	total := CalculateFuel(mass)
	if total != 966 {
		t.Error("failed - should be 966")
	}
}

func TestFor14(t *testing.T) {
	mass := 14
	total := CalculateFuel(mass)
	if total != 2 {
		t.Error("failed - should be 2")
	}
}

func TestFor100756(t *testing.T) {
	mass := 100756
	total := CalculateFuel(mass)
	if total != 50346 {
		t.Error("failed - should be 50346")
	}
}
