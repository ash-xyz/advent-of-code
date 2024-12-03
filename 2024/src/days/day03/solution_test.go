package day03

import "testing"

func TestPart1(t *testing.T) {
	testCases := []struct {
		input    []string
		expected any
	}{
		{
			input:    []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"},
			expected: 161,
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
			input:    []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"},
			expected: 48,
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
