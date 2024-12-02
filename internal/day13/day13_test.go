package day13_test

import (
	"testing"

	"github.com/jenspederm/advent-of-code/internal/day13"
)

func TestRun(t *testing.T) {
	testData := []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
		"",
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}
	t.Run("Part 1", func(t *testing.T) {
		sum := day13.Part1(testData)
		expected := 405
		if sum != expected {
			t.Errorf("Expected %v, got %v", expected, sum)
		}

	})

	t.Run("Part 2", func(t *testing.T) {
		sum := day13.Part2(testData)
		expected := 0
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})

}
