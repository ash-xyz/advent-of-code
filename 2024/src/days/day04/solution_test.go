package day04

import "testing"

func TestPart1(t *testing.T) {
	testCases := []struct {
		input    []string
		expected any
	}{
		{
			input: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			expected: 18,
		},
	}

	s := &Solution{}
	for _, tc := range testCases {
		if got := s.Part1(tc.input); got != tc.expected {
			t.Errorf("Part1() = %v, want %v", got, tc.expected)
		}
	}
}
func TestPart2(t *testing.T) {
	testCases := []struct {
		input    []string
		expected any
	}{
		{
			input: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			expected: 9,
		},
	}

	s := &Solution{}
	for _, tc := range testCases {
		if got := s.Part2(tc.input); got != tc.expected {
			t.Errorf("Part2() = %v, want %v", got, tc.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	input := []string{"test input"}
	s := &Solution{}
	for i := 0; i < b.N; i++ {
		s.Part1(input)
	}
}
