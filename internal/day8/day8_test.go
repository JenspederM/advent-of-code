package day8_test

import (
	"testing"

	"github.com/jenspederm/advent-of-code/internal/day8"
)

func TestRun(t *testing.T) {
	test1 := []string{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}
	test2 := []string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	t.Run("Part 1", func(t *testing.T) {
		sum := day8.Part1(test1)
		expected := 2
		if sum != expected {
			t.Errorf("Test 1: Expected %d, got %d", expected, sum)
		}
		sum2 := day8.Part1(test2)
		expected2 := 6
		if sum2 != expected2 {
			t.Errorf("Test 2: Expected %d, got %d", expected2, sum2)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		sum := day8.Part2(test1)
		expected := 0
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})

}
