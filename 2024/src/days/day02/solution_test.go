package day02

import "testing"

func TestPart1(t *testing.T) {
	testCases := []struct {
		input    []string
		expected any
	}{
		{
			input: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			expected: 2,
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
			input: []string{"3   4",
				"4   3",
				"2   5",
				"1   3",
				"3   9",
				"3   3"},
			expected: 31,
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
