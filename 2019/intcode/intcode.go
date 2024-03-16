package intcode

import "fmt"

const (
	ADD           = 1
	MULTIPLY      = 2
	SAVE          = 3
	OUTPUT        = 4
	JUMP_IF_TRUE  = 5
	JUMP_IF_FALSE = 6
	LESS_THAN     = 7
	EQUALS        = 8
	TERMINATE     = 99

	POSITION_MODE  = 0
	IMMEDIATE_MODE = 1
)

type Machine struct {
	LastOutput int
}

func (m Machine) getArgs(program []int, index int) (int, int, int, int, int) {

	instruction := program[index]
	if instruction == TERMINATE {
		return instruction, 0, 0, 0, 0
	}

	opcode := instruction % 100
	mode1 := (instruction / 100) % 10
	mode2 := (instruction / 1000) % 10
	length := 2 // Default instruction length

	if opcode == SAVE || opcode == OUTPUT {
		return opcode, program[index+1], 0, 0, length
	}

	arg1 := program[index+1]
	if mode1 == POSITION_MODE {
		arg1 = program[arg1]
	}

	arg2 := program[index+2]
	if mode2 == POSITION_MODE {
		arg2 = program[arg2]
	}

	arg3 := program[index+3]
	if opcode == JUMP_IF_TRUE || opcode == JUMP_IF_FALSE {
		length = 3
	} else {
		length = 4
	}

	return opcode, arg1, arg2, arg3, length
}

func (m Machine) Run(input []int, instructions []int) int {
	program := make([]int, len(instructions))
	var results []int
	copy(program, instructions)

	i := 0

	for {
		opcode, a1, a2, a3, instruction_length := m.getArgs(program, i)
		switch opcode {
		case ADD:
			program[a3] = a1 + a2
		case MULTIPLY:
			program[a3] = a1 * a2
		case JUMP_IF_TRUE:
			if a1 != 0 {
				i = a2
				continue
			}
		case JUMP_IF_FALSE:
			if a1 == 0 {
				i = a2
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
		case SAVE:
			if len(input) > 0 {
				code := input[len(input)-1]
				input = input[:1]
				program[a1] = code
			}
		case OUTPUT:
			results = append(results, program[a1])
			return program[a1]
		case TERMINATE:
			fmt.Println(results)
			return results[len(results)-1]
		default:
			panic("Bad Op Code")
		}

		i += instruction_length
	}
}
