package day12_test

import (
	"testing"

	"github.com/jenspederm/advent-of-code/internal/day12"
)

func TestRun(t *testing.T) {
	testData := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}
	t.Run("Part 1", func(t *testing.T) {
		sum := day12.Part1(testData)
		expected := 21
		if sum != expected {
			t.Errorf("Expected %v, got %v", expected, sum)
		}

	})

	t.Run("Part 2", func(t *testing.T) {
		sum := day12.Part2(testData)
		expected := 0
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})

}
