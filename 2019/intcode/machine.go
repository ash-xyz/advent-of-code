package intcode

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

type Machine struct {
	RelativeBase int
}

func (m *Machine) getArgs(program []int, index int) (int, int, int, int, int) {

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
			arg1 += m.RelativeBase
		}
		return opcode, arg1, 0, 0, length
	}

	if mode1 == POSITION_MODE {
		arg1 = program[arg1]
	} else if mode1 == RELATIVE_MODE {
		arg1 = program[m.RelativeBase+arg1]
	}

	if opcode == OUTPUT {
		return opcode, arg1, 0, 0, length
	}

	arg2 := program[index+2]
	if mode2 == POSITION_MODE {
		arg2 = program[arg2]
	} else if mode2 == RELATIVE_MODE {
		arg2 = program[m.RelativeBase+arg2]
	}

	arg3 := program[index+3]
	if mode3 == RELATIVE_MODE {
		arg3 += m.RelativeBase
	}

	if opcode == JUMP_IF_TRUE || opcode == JUMP_IF_FALSE {
		length = 3
	} else if opcode != RELATIVE_BASE_OFFSET {
		length = 4
	}

	return opcode, arg1, arg2, arg3, length
}

func (m *Machine) Run(instructions []int, input <-chan int, output chan<- int) {
	program := make([]int, len(instructions)+30000)
	idx := 0
	copy(program, instructions)

	for {
		opcode, a1, a2, a3, instruction_length := m.getArgs(program, idx)
		switch opcode {
		case ADD:
			program[a3] = a1 + a2
		case MULTIPLY:
			program[a3] = a1 * a2
		case SAVE:
			program[a1] = <-input
		case OUTPUT:
			output <- a1
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
			m.RelativeBase += a1
		case TERMINATE:
			close(output)
			return
		default:
			panic("Bad Op Code")
		}

		idx += instruction_length
	}
}
