package main

import (
	"math"
	"strings"
)

type day7 struct {
	instructions []int
}

var d7 day7

func (d *day7) init() {
	input := ReadFile("inputs/day7.txt")
	array := strings.Split(input[0], ",")
	d.instructions = convertToInt(array)
}

func (d *day7) nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i == -1 {
		d.reverse(nums, 0, n-1)
		return
	}
	j := i + 1
	for j < n && nums[j] > nums[i] {
		j++
	}
	nums[i], nums[j-1] = nums[j-1], nums[i]
	d.reverse(nums, i+1, n-1)
}

func (d *day7) reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func (d *day7) part2() int {
	phaseSetting := []int{5, 6, 7, 8, 9}
	thrust := math.MinInt

	for i := 0; i < 120; i++ {
		amplifiers := [5]*Machine{{0, true}, {0, true}, {0, true}, {0, true}, {0, true}}
		output := 0
		done := false

		for !done {
			for i, amplifier := range amplifiers {
				last_output := amplifier.Run(d.instructions, output, phaseSetting[i])
				if last_output == -1 {
					done = true
					break
				}
				output = last_output
			}
		}

		thrust = max(thrust, output)
		d.nextPermutation(phaseSetting)
	}

	return thrust
}

func (d *day7) part1() int {
	phaseSetting := []int{0, 1, 2, 3, 4}
	thrust := math.MinInt
	for i := 0; i < 120; i++ {
		output := 0
		d.nextPermutation(phaseSetting)
		for _, phase := range phaseSetting {
			machine := &Machine{0, true}
			output = machine.Run(d.instructions, output, phase)
		}

		thrust = max(thrust, output)
	}

	return thrust
}

func (d day7) run() (int, int) {
	return d.part1(), d.part2()
}

// -------------------------------------------------------------------------------------------------------------------------

type Machine struct {
	idx             int // beginning at 0
	hasNotSeenPhase bool
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

func (m *Machine) Run(instructions []int, input int, phaseNum int) int {
	program := make([]int, len(instructions))
	copy(program, instructions)

	for {
		opcode, a1, a2, a3, instruction_length := m.getArgs(program, m.idx)
		switch opcode {
		case ADD:
			program[a3] = a1 + a2
		case MULTIPLY:
			program[a3] = a1 * a2
		case SAVE:
			if m.hasNotSeenPhase {
				program[a1] = phaseNum
				m.hasNotSeenPhase = false
			} else {
				program[a1] = input
			}
		case OUTPUT:
			m.idx += instruction_length
			return program[a1]
		case JUMP_IF_TRUE:
			if a1 != 0 {
				m.idx = a2
				continue
			}
		case JUMP_IF_FALSE:
			if a1 == 0 {
				m.idx = a2
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
		case TERMINATE:
			return -1
		default:
			panic("Bad Op Code")
		}

		m.idx += instruction_length
	}
}
