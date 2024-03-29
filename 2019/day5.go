package main

import (
	"fmt"
	"strings"
)

type day5 struct {
	input []int
}

const (
	ADD                  = 1
	MULTIPLY             = 2
	SAVE                 = 3
	OUTPUT               = 4
	JUMP_IF_TRUE         = 5
	JUMP_IF_FALSE        = 6
	LESS_THAN            = 7
	EQUALS               = 8
	RELATIVE_BASE_OFFSET = 9
	TERMINATE            = 99

	POSITION_MODE  = 0
	IMMEDIATE_MODE = 1
	RELATIVE_MODE  = 2
)

var d5 day5

func (d day5) getArgs(program []int, index int) (int, int, int, int) {

	instruction := program[index]
	if instruction == TERMINATE {
		return instruction, 0, 0, 0
	}

	opcode := instruction % 100

	if opcode == SAVE || opcode == OUTPUT {
		return opcode, program[index+1], 0, 0
	}

	if opcode == ADD || opcode == MULTIPLY || opcode == LESS_THAN || opcode == EQUALS || opcode == JUMP_IF_FALSE || opcode == JUMP_IF_TRUE {

		var arg1, arg2, arg3 int

		arg3 = program[index+3]

		if (instruction%1000)/100 == IMMEDIATE_MODE {
			arg1 = program[index+1]
		} else {
			arg1 = program[program[index+1]]
		}

		if (instruction%10000)/1000 == IMMEDIATE_MODE {
			arg2 = program[index+2]
		} else {
			arg2 = program[program[index+2]]
		}

		return opcode, arg1, arg2, arg3
	}

	return opcode, 0, 0, 0
}

func (d day5) runProgram(input_code int) int {
	program := make([]int, len(d.input))
	var results []int
	copy(program, d.input)

	i := 0

	for {
		opcode, a1, a2, a3 := d.getArgs(program, i)
		switch opcode {
		case ADD:
			program[a3] = a1 + a2
		case MULTIPLY:
			program[a3] = a1 * a2
		case JUMP_IF_TRUE:
			if a1 != 0 {
				i = a2
			} else {
				i += 3
			}
			continue
		case JUMP_IF_FALSE:
			if a1 == 0 {
				i = a2
			} else {
				i += 3
			}
			continue
		case LESS_THAN:
			if a1 < a2 {
				program[a3] = 1
			} else {
				program[a3] = 0
			}
		case EQUALS:
			if a1 == a2 {
				program[a3] = 1
			} else {
				program[a3] = 0
			}
		case SAVE:
			program[a1] = input_code
		case OUTPUT:
			results = append(results, program[a1])
		case TERMINATE:
			fmt.Println(results)
			return results[len(results)-1]
		default:
			panic("Bad Op Code")
		}

		if opcode == ADD || opcode == MULTIPLY || opcode == LESS_THAN || opcode == EQUALS {
			i += 4
		} else {
			i += 2
		}
	}
}

func (d day5) init() {
	input := ReadFile("inputs/day5.txt")
	array := strings.Split(input[0], ",")
	d5.input = convertToInt(array)
}

func (d day5) part2() int {
	return d.runProgram(5)
}

func (d day5) part1() int {
	return -1 // check previous commit for part1 code
}

func (d day5) run() (int, int) {
	return d.part1(), d.part2()
}
