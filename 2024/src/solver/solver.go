package solver

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"aoc2024/src/days/day01"
	"aoc2024/src/days/day02"
)

type Day interface {
	Part1(input []string) any
	Part2(input []string) any
}

type Solver struct {
	day      int
	solution Day
}

func New(day int) *Solver {
	var solution Day
	switch day {
	case 1:
		solution = &day01.Solution{}
	case 2:
		solution = &day02.Solution{}
	// Add other days as needed
	default:
		panic(fmt.Sprintf("Day %d not implemented", day))
	}

	return &Solver{
		day:      day,
		solution: solution,
	}
}

func (s *Solver) Solve() (any, any) {
	input := s.loadInput()
	return s.solution.Part1(input), s.solution.Part2(input)
}

func (s *Solver) Benchmark() {
	input := s.loadInput()

	result := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.solution.Part1(input)
		}
	})
	fmt.Printf("Part 1: %v\n", result)

	result = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s.solution.Part2(input)
		}
	})
	fmt.Printf("Part 2: %v\n", result)
}

func (s *Solver) loadInput() []string {
	inputPath := filepath.Join("inputs", fmt.Sprintf("day%02d.txt", s.day))
	content, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}
