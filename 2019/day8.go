package main

import (
	"math"
	"strings"
)

type day8 struct {
	pixels []int
}

var d8 day8

func (d *day8) init() {
	input := ReadFile("inputs/day8.txt")
	array := strings.Split(input[0], "")
	d.pixels = convertToInt(array)
}

func (d *day8) part2() int {
	return -1
}

func (d *day8) part1() int {
	layerSize := 25 * 6

	minLayerZeros := math.MaxInt
	minLayerIdx := -1
	for i := 0; i < len(d.pixels); i += layerSize {

		zeroCount := 0
		for j := 0; j < layerSize; j++ {
			if d.pixels[i+j] == 0 {
				zeroCount++
			}
		}

		if zeroCount < minLayerZeros {
			minLayerZeros = zeroCount
			minLayerIdx = i
		}
	}

	oneCount, twoCount := 0, 0
	for i := minLayerIdx; i < minLayerIdx+layerSize; i++ {
		if d.pixels[i] == 1 {
			oneCount++
		}
		if d.pixels[i] == 2 {
			twoCount++
		}
	}

	return oneCount * twoCount
}

func (d day8) run() (int, int) {
	return d.part1(), d.part2()
}
