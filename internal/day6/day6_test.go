package day6_test

import (
	"testing"

	"github.com/jenspederm/advent-of-code/internal/day6"
)

func TestRun(t *testing.T) {
	testData := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	t.Run("Part 1", func(t *testing.T) {
		sum := day6.Part1(testData)
		expected := 288
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})
	t.Run("Part 2", func(t *testing.T) {
		sum := day6.Part2(testData)
		expected := 71503
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})
}
