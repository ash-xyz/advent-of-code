package main

import (
	"fmt"
	"strings"
)

type day9 struct {
	instructions []int
}

var d9 day9

func (d *day9) init() {
	input := ReadFile("inputs/day9.txt")
	array := strings.Split(input[0], ",")
	d.instructions = convertToInt(array)
}

func (d *day9) part2() int {
	machine := Machine9{0}
	return machine.Run(d.instructions, 2)
}

func (d *day9) part1() int {
	machine := Machine9{0}
	return machine.Run(d.instructions, 1)
}

func (d day9) run() (int, int) {
	return d.part1(), d.part2()
}

// -------------------------------------------------------------------------------------------------------------------------

type Machine9 struct {
	relative_base int
}

func (m *Machine9) getArgs(program []int, index int) (int, int, int, int, int) {

	instruction := program[index]
	if instruction == TERMINATE {
		return instruction, 0, 0, 0, 0
	}

	opcode := instruction % 100
	mode1 := (instruction / 100) % 10
	mode2 := (instruction / 1000) % 10
	mode3 := (instruction / 10000) % 10
	length := 2 // Default instruction length

	arg1 := program[index+1]
	if opcode == SAVE {
		if mode1 == RELATIVE_MODE {
			arg1 += m.relative_base
		}
		return opcode, arg1, 0, 0, length
	}

	if mode1 == POSITION_MODE {
		arg1 = program[arg1]
	} else if mode1 == RELATIVE_MODE {
		arg1 = program[m.relative_base+arg1]
	}

	if opcode == OUTPUT {
		return opcode, arg1, 0, 0, length
	}

	arg2 := program[index+2]
	if mode2 == POSITION_MODE {
		arg2 = program[arg2]
	} else if mode2 == RELATIVE_MODE {
		arg2 = program[m.relative_base+arg2]
	}

	arg3 := program[index+3]
	if mode3 == RELATIVE_MODE {
		arg3 += m.relative_base
	}

	if opcode == JUMP_IF_TRUE || opcode == JUMP_IF_FALSE {
		length = 3
	} else if opcode != RELATIVE_BASE_OFFSET {
		length = 4
	}

	return opcode, arg1, arg2, arg3, length
}

func (m *Machine9) Run(instructions []int, input int) int {
	program := make([]int, len(instructions)+30000)
	idx := 0
	copy(program, instructions)
	last_result := -1

	for {
		opcode, a1, a2, a3, instruction_length := m.getArgs(program, idx)
		switch opcode {
		case ADD:
			program[a3] = a1 + a2
		case MULTIPLY:
			program[a3] = a1 * a2
		case SAVE:
			program[a1] = input
		case OUTPUT:
			fmt.Println("Output: ", a1)
			last_result = a1
		case JUMP_IF_TRUE:
			if a1 != 0 {
				idx = a2
				continue
			}
		case JUMP_IF_FALSE:
			if a1 == 0 {
				idx = a2
				continue
			}
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
		case RELATIVE_BASE_OFFSET:
			m.relative_base += a1
		case TERMINATE:
			fmt.Println("DONE")
			return last_result
		default:
			panic("Bad Op Code")
		}

		idx += instruction_length
	}
}
