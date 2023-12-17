package day2_test

import (
	"testing"

	"github.com/jenspederm/advent-of-code/internal/day2"
)

func TestRun(t *testing.T) {
	testData := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	t.Run("Part 1", func(t *testing.T) {
		sum := day2.Part1(testData)
		expected := 8
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		sum := day2.Part2(testData)
		expected := 2286
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})

}
