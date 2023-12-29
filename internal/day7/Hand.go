package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	id    int
	value string
}

type Cards []Card

type Hand struct {
	cards Cards
	bid   int
	Rank  Rank
	Ranks []Rank
}

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

func (c Cards) String() string {
	s := []string{}
	for _, card := range c {
		s = append(s, card.value)
	}
	return strings.Join(s, "")
}

func NewHand(cardString string, bidString string, cardValueMap map[string]int, withJokers ...bool) Hand {
	h := Hand{}
	cards := h.parseCards(strings.Split(cardString, ""), cardValueMap)
	bid, err := strconv.Atoi(bidString)
	if err != nil {
		panic(err)
	}
	h.cards = cards
	h.Rank = h.GetRank(cards, withJokers...)
	h.bid = bid
	return h
}

func (h Hand) parseCards(cardString []string, cardValueMap map[string]int) Cards {
	cards := Cards{}
	for _, char := range cardString {
		cards = append(cards, Card{id: cardValueMap[char], value: char})
	}
	return cards
}

func (h Hand) String() string {
	return fmt.Sprintf("%s %d %d", h.cards.String(), h.bid, h.Rank)
}

func (h *Hand) GetRank(cards Cards, withJokers ...bool) Rank {
	nJokers := 0
	tally := map[Card]int{}
	for _, card := range cards {
		if len(withJokers) > 0 && withJokers[0] && card.value == "J" {
			nJokers++
			continue
		}
		tally[card]++
	}

	counts := []int{}
	for _, v := range tally {
		counts = append(counts, v)
	}

	sort.Ints(counts)

	highest := 0 + nJokers
	if len(counts) > 0 {
		highest = counts[len(counts)-1] + nJokers
	}
	secondHighest := 0 + nJokers
	if len(counts) > 1 {
		secondHighest = counts[len(counts)-2]
	}

	switch highest {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if secondHighest == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 2:
		if secondHighest == 2 {
			return TwoPairs
		}
		return OnePair
	default:
		return HighCard
	}
}
