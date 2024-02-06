package main

import (
	"bufio"
	"os"
	"strconv"
)

type aoc interface {
	part1() int
	part2() int
	run() (int, int)
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
