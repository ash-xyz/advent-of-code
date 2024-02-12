package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"
)

type day3 struct {
	input [2][]string
}

var d3 day3

func (d *day3) init() {
	input := ReadFile("inputs/day3.txt")
	d.input[0] = strings.Split(input[0], ",")
	d.input[1] = strings.Split(input[1], ",")
}

type state struct {
	x     int
	y     int
	steps int
}

func (d *day3) part2() int {
	seen := make([]map[image.Point]int, len(d.input))
	fmt.Println(seen[0][image.Point{0, 0}])

	for i, wire := range d.input {
		curState := state{0, 0, 0}
		seen[i] = map[image.Point]int{}
		for _, move := range wire {
			for j, _ := strconv.Atoi(move[1:]); j > 0; j-- {
				d := map[byte]image.Point{'U': {0, -1}, 'D': {0, 1}, 'L': {-1, 0}, 'R': {1, 0}}[move[0]]
				curState.x, curState.y, curState.steps = curState.x+d.X, curState.y+d.Y, curState.steps+1
				if (seen[i][image.Point{curState.x, curState.y}] == 0) {
					seen[i][image.Point{curState.x, curState.y}] = curState.steps
				} else {
					seen[i][image.Point{curState.x, curState.y}] = min(seen[i][image.Point{curState.x, curState.y}], curState.steps)
				}
			}
		}
	}

	mini := MaxInt
	for p := range seen[1] {
		if _, ok := seen[0][p]; ok {
			steps := seen[0][p] + seen[1][p]
			mini = min(steps, mini)
		}
	}

	return mini
}

func (d *day3) part1() int {
	seen := make([]map[image.Point]struct{}, len(d.input))

	for i, wire := range d.input {
		curState := state{0, 0, 0}
		seen[i] = map[image.Point]struct{}{}
		for _, move := range wire {
			for j, _ := strconv.Atoi(move[1:]); j > 0; j-- {
				d := map[byte]image.Point{'U': {0, -1}, 'D': {0, 1}, 'L': {-1, 0}, 'R': {1, 0}}[move[0]]
				curState.x, curState.y = curState.x+d.X, curState.y+d.Y
				seen[i][image.Point{curState.x, curState.y}] = struct{}{}
			}
		}
	}

	mini := MaxInt
	for p := range seen[1] {
		if _, ok := seen[0][p]; ok {
			dist := abs(p.X) + abs(p.Y)
			mini = min(dist, mini)
		}
	}

	return mini
}

func (d day3) run() (int, int) {
	return d.part1(), d.part2()
}
