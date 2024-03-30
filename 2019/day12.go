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
	leMoons := make([]Moon, 0, len(d.moons))
	leMoons = append(leMoons, d.moons...)

	xCycleLength := d.findCycleLengthAlongAxis(leMoons, d.moons, 0)
	yCycleLength := d.findCycleLengthAlongAxis(leMoons, d.moons, 1)
	zCycleLength := d.findCycleLengthAlongAxis(leMoons, d.moons, 2)

	return LCM(xCycleLength, yCycleLength, zCycleLength)
}

func (d *day12) part1() int {
	leMoons := make([]Moon, 0, len(d.moons))
	leMoons = append(leMoons, d.moons...)
	for i := 0; i < 1000; i++ {
		d.timeStep(leMoons)
	}

	energy := 0
	for _, moon := range leMoons {
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

	for i := range moons {
		moons[i].applyVelocity()
	}
}

func (d day12) timeStepAlongAxis(moons []Moon, axis int) {
	for i := range moons {
		for j := i + 1; j < len(moons); j++ {
			moons[i].applyGravityAlongAxis(&moons[j], axis)
		}
	}

	for i := range moons {
		moons[i].applyVelocityAlongAxis(axis)
	}
}

func (d day12) findCycleLengthAlongAxis(moons, original []Moon, axis int) int {
	cycleLength := 1
	for ; ; cycleLength++ {
		d.timeStepAlongAxis(moons, axis)

		cycleIsMet := true
		for i := range d.moons {
			if original[i].pos[axis] != moons[i].pos[axis] || original[i].vel[axis] != moons[i].vel[axis] {
				cycleIsMet = false
				break
			}
		}
		if cycleIsMet {
			break
		}
	}

	return cycleLength
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

func (m *Moon) applyGravityAlongAxis(m1 *Moon, axis int) {
	if m.pos[axis] < m1.pos[axis] {
		m.vel[axis]++
		m1.vel[axis]--
	} else if m.pos[axis] > m1.pos[axis] {
		m.vel[axis]--
		m1.vel[axis]++
	}

}

func (m *Moon) applyVelocityAlongAxis(axis int) {
	m.pos[axis] += m.vel[axis]
}
