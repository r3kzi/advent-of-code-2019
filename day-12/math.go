package main

// see https://stackoverflow.com/questions/3154454/what-is-the-most-efficient-way-to-calculate-the-least-common-multiple-of-two-int
// see https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
