package day7_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/jenspederm/advent-of-code/internal/day7"
)

func TestRun(t *testing.T) {
	testData := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}

	paringData := [][]string{
		{"234AQ", "0"},
		{"234AA", "1"},
		{"233AA", "2"},
		{"23AAA", "3"},
		{"33AAA", "4"},
		{"3AAAA", "5"},
		{"AAAAA", "6"},
	}

	t.Run("Test Paring", func(t *testing.T) {
		cardValueMap := day7.ConstructCardValueMap()
		for _, data := range paringData {
			cardString := strings.Split(data[0], "")
			tmp, err := strconv.Atoi(data[1])
			if err != nil {
				t.Errorf("Error converting %s to int", data[1])
			}
			cards := day7.ParseCards(cardString, cardValueMap)
			expectedRank := day7.Rank(tmp)
			rank := day7.GetRank(cards)
			if rank != expectedRank {
				t.Errorf("Expected %d, got %d", expectedRank, rank)
			}
		}
	})

	t.Run("Part 1", func(t *testing.T) {
		sum := day7.Part1(testData)
		expected := 6440
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})
	t.Run("Part 2", func(t *testing.T) {
		sum := day7.Part2(testData)
		expected := 0
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})
}
