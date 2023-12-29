package day7_test

import (
	"strconv"
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

	t.Run("Test Paring", func(t *testing.T) {
		paringData := [][]string{
			{"234AQ 0", "0"},
			{"234AA 0", "1"},
			{"233AA 0", "2"},
			{"23AAA 0", "3"},
			{"33AAA 0", "4"},
			{"2AAAA 0", "5"},
			{"3AAAA 0", "5"},
			{"AAAAA 0", "6"},
		}

		for i := range paringData {
			data := paringData[i]
			camelCards := day7.NewCamelCards([]string{data[0]})
			tmp, err := strconv.Atoi(data[1])
			if err != nil {
				t.Errorf("Error converting %s to int", data[1])
			}
			hands := camelCards.Hands
			rank := hands[0].Rank
			expectedRank := day7.Rank(tmp)

			if rank != expectedRank {
				t.Errorf("Expected %d, got %d", expectedRank, rank)
			}
		}
	})

	type Pair struct {
		hand string
		rank day7.Rank
	}

	t.Run("Test Paring with Jokers", func(t *testing.T) {
		paringDataWithJokers := []Pair{
			{"33AJJ 0", day7.FourOfAKind},
			{"2JJAA 0", day7.FourOfAKind},
			{"JJJJJ 0", day7.FiveOfAKind},
			{"33AAJ 0", day7.FullHouse},
			{"33A2J 1", day7.ThreeOfAKind},
			{"31AJJ 1", day7.ThreeOfAKind},
			{"3142J 1", day7.OnePair},
		}

		for i := range paringDataWithJokers {
			data := paringDataWithJokers[i]
			camelCards := day7.NewCamelCards([]string{data.hand}, true)

			hands := camelCards.Hands
			rank := hands[0].Rank
			if rank != data.rank {
				ranks := hands[0].Ranks
				t.Errorf("Failed on %s. Expected %d, got %d, %v", data.hand, data.rank, rank, ranks)
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
		expected := 5905
		if sum != expected {
			t.Errorf("Expected %d, got %d", expected, sum)
		}
	})
}
