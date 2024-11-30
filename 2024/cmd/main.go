package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"aoc2024/src/solver"
)

func main() {
	day := flag.Int("day", 0, "day to solve")
	benchmark := flag.Bool("bench", false, "run benchmarks")
	flag.Parse()

	if *day == 0 {
		log.Fatal("Please specify a day using -day flag")
	}

	s := solver.New(*day)

	if *benchmark {
		s.Benchmark()
		return
	}

	start := time.Now()
	part1, part2 := s.Solve()
	duration := time.Since(start)

	fmt.Printf("Day %d\n", *day)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
	fmt.Printf("Time: %v\n", duration)
}
