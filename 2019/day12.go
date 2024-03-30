package main

import (
	"regexp"
	"strconv"
)

type day12 struct {
	moons []Moon
}

var d12 day12

func (d *day12) init() {
	input := ReadFile("inputs/day12.txt")
	re := regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)

	d.moons = make([]Moon, 0, len(input))
	for _, moon := range input {
		matches := re.FindStringSubmatch(moon)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		z, _ := strconv.Atoi(matches[3])
		d.moons = append(d.moons, NewMoon(x, y, z))
	}
}

func (d *day12) part2() int {
	return -1
}

func (d *day12) part1() int {
	for i := 0; i < 1000; i++ {
		d.timeStep(d.moons)
	}

	energy := 0
	for _, moon := range d.moons {
		kineticEnergy, potentialEnergy := 0, 0
		for i := 0; i < 3; i++ {
			kineticEnergy += abs(moon.pos[i])
			potentialEnergy += abs(moon.vel[i])
		}
		energy += kineticEnergy * potentialEnergy
	}
	return energy
}

func (d day12) run() (int, int) {
	return d.part1(), d.part2()
}

func (d day12) timeStep(moons []Moon) {
	for i := range moons {
		for j := i + 1; j < len(moons); j++ {
			moons[i].applyGravity(&moons[j])
		}
	}

	//we just learnt that range functions in golang return a reference/value - not a pointer; so using indexes is advisable
	for i := range moons {
		moons[i].applyVelocity()
	}
}

// ------------------------------------------------------------

type Moon struct {
	pos [3]int
	vel [3]int
}

func NewMoon(x, y, z int) Moon {
	return Moon{[3]int{x, y, z}, [3]int{0, 0, 0}}
}

func (m *Moon) applyGravity(m1 *Moon) {
	for i := 0; i < 3; i++ {
		if m.pos[i] < m1.pos[i] {
			m.vel[i]++
			m1.vel[i]--
		} else if m.pos[i] > m1.pos[i] {
			m.vel[i]--
			m1.vel[i]++
		}
	}
}

func (m *Moon) applyVelocity() {
	for i := 0; i < 3; i++ {
		m.pos[i] += m.vel[i]
	}
}
