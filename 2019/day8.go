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
	d.pixels = convertToInt(strings.Split(input[0], ""))
}

func (d *day8) part2() int {
	layerSize := 25 * 6
	decodedImage := make([]int, layerSize)
	//initialise array with all 2's
	for i := 0; i < layerSize; i++ {
		decodedImage[i] = 2
	}

	for i := 0; i < len(d.pixels); i += layerSize {
		for j := 0; j < layerSize; j++ {
			if decodedImage[j] == 2 {
				decodedImage[j] = d.pixels[i+j]
			}
		}
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			pix := decodedImage[i*25+j]
			if pix == 1 {
				print("#")
			} else {
				print(" ")
			}
		}
		println("")
	}

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
