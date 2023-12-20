package day3_test

import (
	"testing"

	"github.com/jenspederm/advent-of-code/internal/day3"
)

func TestRun(t *testing.T) {
	testData := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	single := []string{
		"467..9..",
		"....*...",
	}

	last := []string{
		"467..99.",
		"....*..9",
	}

	t.Run("single", func(t *testing.T) {
		sum := day3.Part1(single)
		expected := 9
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})
	t.Run("last", func(t *testing.T) {
		sum := day3.Part1(last)
		expected := 108
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})
	t.Run("Part 1", func(t *testing.T) {
		sum := day3.Part1(testData)
		expected := 4361
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})
	t.Run("Part 2", func(t *testing.T) {
		sum := day3.Part2(testData)
		expected := 467835
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})
}
