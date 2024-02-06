package main

import (
	"strings"
)

type day2 struct {
	input []int
}

var d2 day2

func runProgram(input []int) int {
	program := make([]int, len(input))
	copy(program, input)
	i := 0
	for {
		opcode := program[i]
		noun, verb, output := program[i+1], program[i+2], program[i+3]
		switch opcode {
		case 1:
			program[output] = program[noun] + program[verb]
		case 2:
			program[output] = program[noun] * program[verb]
		case 99:
			return program[0]
		default:
			panic("Bad Op Code")
		}
		i += 4
	}
}

func (d day2) init() {
	input := ReadFile("inputs/day2.txt")
	array := strings.Split(input[0], ",")
	d2.input = convertToInt(array)
}

func (d day2) part2() int {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			d.input[1], d.input[2] = i, j
			if runProgram(d.input) == 19690720 {
				return 100*i + j
			}
		}
	}
	return -1
}

func (d day2) part1() int {
	d.input[1], d.input[2] = 12, 2
	return runProgram(d.input)
}

func (d day2) run() (int, int) {
	return d.part1(), d.part2()
}
