package day11_test

import (
	"testing"

	"github.com/jenspederm/advent-of-code/internal/day11"
)

func TestRun(t *testing.T) {
	testData := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	t.Run("Part 1", func(t *testing.T) {
		sum := day11.Part1(testData)
		expected := 374
		if sum != expected {
			t.Errorf("Expected %v, got %v", expected, sum)
		}

	})

	t.Run("Part 2", func(t *testing.T) {
		sum := day11.Part2(testData, 10)
		expected := 1030
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
		sum = day11.Part2(testData, 100)
		expected = 8410
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})

}
