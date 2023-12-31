package day9_test

import (
	"testing"

	"github.com/jenspederm/advent-of-code/internal/day9"
)

func TestRun(t *testing.T) {
	testData := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}

	t.Run("Part 1", func(t *testing.T) {
		sum := day9.Part1(testData)
		expected := 114
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		sum := day9.Part2(testData)
		expected := 0
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})

}
