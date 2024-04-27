package main

import (
	"aoc2019/intcode"
	"strings"
)

type day17 struct {
	instructions []int
}

var d17 day17

func (d *day17) init() {
	input := ReadFile("inputs/day17.txt")
	array := strings.Split(input[0], ",")
	d.instructions = convertToInt(array)
}

func (d *day17) part2() int {
	//I give up; Maybe I'll come back to this some other year :p
	return -1
}
func (d *day17) part1() int {
	machine := intcode.Machine{RelativeBase: 0}
	input, output := make(chan int, 1), make(chan int)

	go machine.Run(d.instructions, input, output)

	scaffoldMap := make([][]byte, 0)
	row := make([]byte, 0)

	for {
		char, ok := <-output
		if !ok {
			break
		}

		if char == 10 {
			if len(row) > 0 {
				scaffoldMap = append(scaffoldMap, row)
				row = make([]byte, 0)
			}
		} else {
			row = append(row, byte(char))
		}
	}

	sum := 0
	for i := 1; i < len(scaffoldMap)-1; i++ {
		for j := 1; j < len(scaffoldMap[i])-1; j++ {
			if scaffoldMap[i][j] == '#' &&
				scaffoldMap[i-1][j] == '#' &&
				scaffoldMap[i+1][j] == '#' &&
				scaffoldMap[i][j-1] == '#' &&
				scaffoldMap[i][j+1] == '#' {
				sum += i * j
			}
		}
	}

	return sum
}

func (d day17) run() (int, int) {
	return d.part1(), d.part2()
}
