package main

import "math"

type Vector struct {
	x int
	y int
	z int
}

func (vector Vector) Add(other Vector) Vector {
	return Vector{
		x: vector.x + other.x,
		y: vector.y + other.y,
		z: vector.z + other.z,
	}
}

func (vector Vector) Subtract(other Vector) Vector {
	return Vector{
		x: vector.x - other.x,
		y: vector.y - other.y,
		z: vector.z - other.z,
	}
}

func (vector Vector) Normalize() Vector {
	return Vector{
		x: normalize(vector.x),
		y: normalize(vector.y),
		z: normalize(vector.z),
	}
}

func (vector Vector) Abs() Vector {
	return Vector{
		x: abs(vector.x),
		y: abs(vector.y),
		z: abs(vector.z),
	}
}

func (vector Vector) calculateEnergy() int {
	v := Vector{
		x: abs(vector.x),
		y: abs(vector.y),
		z: abs(vector.z),
	}
	return v.x + v.y + v.z
}

func normalize(x int) int {
	if x > 0 {
		return -1
	}
	if x < 0 {
		return 1
	}
	return 0
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
