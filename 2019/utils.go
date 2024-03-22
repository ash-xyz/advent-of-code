package main

import (
	"bufio"
	"os"
	"strconv"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type aoc interface {
	init() // Usually should be used to read in the file
	part1() int
	part2() int
	run() (int, int) // Returns answers for part 1 and part 2
}
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func abs[T Number](a T) T {
	if a > 0 {
		return a
	}
	return -a
}

func ReadFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func convertToInt(x []string) []int {
	array := make([]int, len(x))
	var err error

	for i, val := range x {
		array[i], err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}

	return array
}

func convertToInt64(x []string) []int64 {
	array := make([]int64, len(x))
	var err error

	for i, val := range x {
		array[i], err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	return array
}
