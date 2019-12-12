package main

import (
	"fmt"
	"time"
)

type Moon struct {
	pos      Vector
	velocity Vector
}

func main() {
	start := time.Now()

	fmt.Println("----------------")

	moons := Read("input.txt")

	fmt.Println("--- Part One ---")
	calculateEnergy(moons)

	fmt.Println("--- Part Two ---")
	calculateSteps(moons)

	fmt.Println("----------------")
	elapsed := time.Since(start)
	fmt.Printf("Took %s", elapsed)
}

func calculateSteps(input []Moon) {
	moons := make([]Moon, len(input))
	copy(moons, input)

	mx, my, mz := make(map[string]bool), make(map[string]bool), make(map[string]bool)

	cx, cy, cz := 0, 0, 0

	steps := 0
	for {

		if cx == 0 {
			x := fmt.Sprintf("%v %v %v %v %v %v %v %v", moons[0].pos.x, moons[1].pos.x, moons[2].pos.x, moons[3].pos.x,
				moons[0].velocity.x, moons[1].velocity.x, moons[2].velocity.x, moons[3].velocity.x)

			if _, ok := mx[x]; ok {
				cx = steps
			} else {
				mx[x] = true
			}
		}

		if cy == 0 {
			y := fmt.Sprintf("%v %v %v %v %v %v %v %v", moons[0].pos.y, moons[1].pos.y, moons[2].pos.y, moons[3].pos.y,
				moons[0].velocity.y, moons[1].velocity.y, moons[2].velocity.y, moons[3].velocity.y)

			if _, ok := my[y]; ok {
				cy = steps
			} else {
				my[y] = true
			}
		}

		if cz == 0 {
			z := fmt.Sprintf("%v %v %v %v %v %v %v %v", moons[0].pos.z, moons[1].pos.z, moons[2].pos.z, moons[3].pos.z,
				moons[0].velocity.z, moons[1].velocity.z, moons[2].velocity.z, moons[3].velocity.z)

			if _, ok := mz[z]; ok {
				cz = steps
			} else {
				mz[z] = true
			}
		}

		simulateTimeSteps(moons, 1)

		if cx != 0 && cy != 0 && cz != 0 {
			break
		}
		steps++
	}
	fmt.Println(LCM(cx, cy, cz))
}

func calculateEnergy(input []Moon) {
	moons := make([]Moon, len(input))
	copy(moons, input)

	simulateTimeSteps(moons, 10)

	var totalEnergy int
	for _, moon := range moons {
		totalEnergy += moon.pos.calculateEnergy() * moon.velocity.calculateEnergy()
	}
	fmt.Println(totalEnergy)
}

func simulateTimeSteps(moons []Moon, i int) {
	for timeStep := 0; timeStep < i; timeStep++ {
		// Calculate velocity
		for index, m1 := range moons {
			for _, m2 := range moons {
				if m1.pos == m2.pos {
					continue
				}
				n := m1.pos.Subtract(m2.pos).Normalize()
				m1.velocity = m1.velocity.Add(n)
			}
			moons[index] = m1
		}

		// Apply velocity
		for index, moon := range moons {
			moons[index].pos = moon.pos.Add(moon.velocity)
		}
	}
}
