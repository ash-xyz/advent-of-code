package main

import (
	"fmt"
	"strings"
)

type day11 struct {
	instructions []int
}

var d11 day11

func (d *day11) init() {
	input := ReadFile("inputs/day11.txt")
	array := strings.Split(input[0], ",")
	d.instructions = convertToInt(array)
}

func (d day11) runRobot(instructions []int, startingColour int) map[Position]int {
	panels := make(map[Position]int)
	machine := Machine11{0}
	r := newRobot()
	panels[r.location] = startingColour
	input, output, done := make(chan int), make(chan int), make(chan bool)

	go func() {
		machine.Run(instructions, input, output)
		done <- true
	}()

	for {
		select {
		case input <- panels[r.location]:
			paint, ok := <-output
			if !ok {
				break
			}

			panels[r.location] = paint

			dir, ok2 := <-output
			if !ok2 {
				break
			}
			if dir == 0 {
				r.Left()
			} else if dir == 1 {
				r.Right()
			}

			r.Forward()
		case <-done:
			return panels
		}
	}
}

func (d *day11) part2() int {
	panels := d.runRobot(d.instructions, 1)

	minX, minY, maxX, maxY := 0, 0, 0, 0

	for k := range panels {
		if k.X < minX {
			minX = k.X
		}
		if k.X > maxX {
			maxX = k.X
		}
		if k.Y < minY {
			minY = k.Y
		}
		if k.Y > maxY {
			maxY = k.Y
		}
	}

	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			c := panels[Position{i, j}]
			switch c {
			case 1:
				fmt.Print("#")
			case 0:
				fmt.Print("-")
			}
		}
		fmt.Println()
	}

	return len(panels)
}

func (d *day11) part1() int {
	return len(d.runRobot(d.instructions, 0))
}

func (d day11) run() (int, int) {
	return d.part1(), d.part2()
}

// -----
type Position struct {
	X int
	Y int
}

type Robot struct {
	location  Position
	Direction Position
}

func newRobot() *Robot {
	r := Robot{Position{0, 0}, Position{0, -1}}

	return &r
}

func (r *Robot) Left() {
	r.Direction = Position{r.Direction.Y, -r.Direction.X}
}

func (r *Robot) Right() {
	r.Direction = Position{-r.Direction.Y, r.Direction.X}
}

func (r *Robot) Forward() {
	r.location.X += r.Direction.X
	r.location.Y += r.Direction.Y
}

// ----
type Machine11 struct {
	relative_base int
}

func (m *Machine11) getArgs(program []int, index int) (int, int, int, int, int) {

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

func (m *Machine11) Run(instructions []int, input <-chan int, output chan<- int) {
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
			m.relative_base += a1
		case TERMINATE:
			close(output)
			return
		default:
			panic("Bad Op Code")
		}

		idx += instruction_length
	}
}
