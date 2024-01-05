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
		sum := day10.Part2([]string{
			"...........",
			".S-------7.",
			".|F-----7|.",
			".||.....||.",
			".||.....||.",
			".|L-7.F-J|.",
			".|..|.|..|.",
			".L--J.L--J.",
			"...........",
		})
		expected := 4
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}

		sum = day10.Part2([]string{
			".F----7F7F7F7F-7....",
			".|F--7||||||||FJ....",
			".||.FJ||||||||L7....",
			"FJL7L7LJLJ||LJ.L-7..",
			"L--J.L7...LJS7F-7L7.",
			"....F-J..F7FJ|L7L7L7",
			"....L7.F7||L7|.L7L7|",
			".....|FJLJ|FJ|F7|.LJ",
			"....FJL-7.||.||||...",
			"....L---J.LJ.LJLJ...",
		})
		expected = 8
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}

		sum = day10.Part2([]string{
			"FF7FSF7F7F7F7F7F---7",
			"L|LJ||||||||||||F--J",
			"FL-7LJLJ||||||LJL-77",
			"F--JF--7||LJLJ7F7FJ-",
			"L---JF-JLJ.||-FJLJJ7",
			"|F|F-JF---7F7-L7L|7|",
			"|FFJF7L7F-JF7|JL---7",
			"7-L-JL7||F7|L7F-7F7|",
			"L.L7LFJ|||||FJL7||LJ",
			"L7JLJL-JLJLJL--JLJ.L",
		})
		expected = 10
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})

}
