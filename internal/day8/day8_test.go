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

	part2_test := []string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
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
		sum := day8.Part2(part2_test)
		expected := 6
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})

}
