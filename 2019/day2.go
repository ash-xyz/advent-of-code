package main

import (
	"fmt"
	"strings"
)

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

func p2(input []int) {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			input[1], input[2] = i, j
			if runProgram(input) == 19690720 {
				fmt.Println("Part 2: ", 100*i+j)
				return
			}
		}
	}
}

func p1(input []int) {
	input[1], input[2] = 12, 2

	for i := 0; i < len(input); i += 4 {
		opcode := input[i]

		if opcode == 99 {
			break
		} else if opcode == 1 {
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
		} else if opcode == 2 {
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		}
	}

	println("Part 1: ", input[0])
}

func day2() {
	input := ReadFile("inputs/day2.txt")
	array := strings.Split(input[0], ",")
	intArray := convertToInt(array)

	p2(intArray)
}
