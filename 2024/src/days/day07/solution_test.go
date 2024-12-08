package day07

import "testing"

func TestPart1(t *testing.T) {
	testCases := []struct {
		input    []string
		expected any
	}{
		{
			input: []string{
				"190: 10 19",
				"3267: 81 40 27",
				"83: 17 5",
				"156: 15 6",
				"7290: 6 8 6 15",
				"161011: 16 10 13",
				"192: 17 8 14",
				"21037: 9 7 18 13",
				"292: 11 6 16 20",
			},
			expected: 3749,
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
				"190: 10 19",
				"3267: 81 40 27",
				"83: 17 5",
				"156: 15 6",
				"7290: 6 8 6 15",
				"161011: 16 10 13",
				"192: 17 8 14",
				"21037: 9 7 18 13",
				"292: 11 6 16 20",
			},
			expected: 11387,
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
