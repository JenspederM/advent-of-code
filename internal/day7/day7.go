package day7

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

type Rank int

const (
	HighCard Rank = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Card struct {
	id    int
	value string
}

type Cards []Card

func (c Cards) String() string {
	s := []string{}
	for _, card := range c {
		s = append(s, card.value)
	}
	return strings.Join(s, "")
}

type Hand struct {
	cards Cards
	bid   int
	rank  Rank
}

func GetRank(cards Cards) Rank {
	tally := map[Card]int{}
	for _, card := range cards {
		tally[card]++
	}

	switch len(tally) {
	case 1:
		return FiveOfAKind
	case 2:
		nPairs := 0
		for _, v := range tally {
			if v == 4 {
				return FourOfAKind
			} else if v == 2 {
				nPairs++
			}
		}
		if nPairs == 2 {
			return TwoPairs
		} else {
			return FullHouse
		}
	case 3:
		for _, v := range tally {
			if v == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPairs
	case 4:
		return OnePair
	default:
		return HighCard
	}
}

func (h Hand) String() string {
	return strings.Join([]string{h.cards.String(), strconv.Itoa(h.bid), strconv.Itoa(int(h.rank))}, " ")
}

func ConstructCardValueMap() map[string]int {
	cardValueMap := map[string]int{}
	for i, c := range "23456789TJQKA" {
		cardValueMap[string(c)] = i
	}
	return cardValueMap
}

func ParseCards(cardString []string, cardValueMap map[string]int) Cards {
	cards := Cards{}
	for _, c := range cardString {
		cards = append(cards, Card{id: cardValueMap[c], value: c})
	}
	return cards
}

func ParseLines(lines []string, cardValueMap map[string]int) []Hand {
	hands := []Hand{}
	for _, line := range lines {
		s := strings.Split(line, " ")
		cardString := strings.Split(s[0], "")
		bid, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		cards := ParseCards(cardString, cardValueMap)
		rank := GetRank(cards)
		hands = append(hands, Hand{cards: cards, bid: bid, rank: rank})
	}
	return hands
}

func Part1(lines []string) int {
	cardValueMap := ConstructCardValueMap()
	hands := ParseLines(lines, cardValueMap)

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].rank == hands[j].rank {
			for k := 0; k < len(hands[i].cards); k++ {
				if hands[i].cards[k].id != hands[j].cards[k].id {
					return hands[i].cards[k].id < hands[j].cards[k].id
				}
			}
		}
		return hands[i].rank < hands[j].rank
	})

	sum := 0

	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}

	return sum
}

func Part2(lines []string) int {
	return 0
}

func Run() {
	lines := utils.LoadText("./data/day7.txt")

	println()
	println("##############################")
	println("#            Day 7           #")
	println("##############################")
	println()
	println("Part 1")
	println(Part1(lines))

	println("Part 2")
	println(Part2(lines))
}
