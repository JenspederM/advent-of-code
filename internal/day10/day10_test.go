package day10_test

import (
	"testing"

	"github.com/jenspederm/advent-of-code/internal/day10"
)

func TestRun(t *testing.T) {
	testData := []string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	}

	t.Run("Part 1", func(t *testing.T) {
		sum := day10.Part1(testData)
		expected := 4
		if sum != expected {
			t.Errorf("Test1: Expected %d, got %d", expected, sum)
		}

		sum = day10.Part1([]string{
			"..F7.",
			".FJ|.",
			"SJ.L7",
			"|F--J",
			"LJ...",
		})
		expected = 8
		if sum != expected {
			t.Errorf("Test2: Expected %d, got %d", expected, sum)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		sum := day10.Part2(testData)
		expected := 0
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})

}
