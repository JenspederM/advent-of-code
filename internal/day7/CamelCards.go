package day7

import (
	"fmt"
	"sort"
	"strings"
)

type CamelCards struct {
	cardValueMap map[string]int
	Hands        []Hand
}

func NewCamelCards(lines []string, withJokers ...bool) CamelCards {
	c := CamelCards{}

	values := ""

	if len(withJokers) > 0 {
		values = "J23456789TQKA"
	} else {
		values = "23456789TJQKA"
	}

	cardValueMap := map[string]int{}
	for i, c := range values {
		cardValueMap[string(c)] = i
	}

	hands := []Hand{}
	for _, line := range lines {
		s := strings.Split(line, " ")
		hands = append(hands, NewHand(s[0], s[1], cardValueMap, withJokers...))
	}

	c.cardValueMap = cardValueMap
	c.Hands = hands
	return c
}

func (c CamelCards) Sum(verbose ...bool) int {
	c.SortHands()
	sum := 0
	for i, hand := range c.Hands {
		pos := i + 1
		sum += hand.bid * pos
		if len(verbose) > 0 && verbose[0] {
			fmt.Printf("%s: (%3d * %4d) == %d\n", hand.cards.String(), hand.bid, pos, hand.bid*pos)
		}
	}
	return sum
}

func (c *CamelCards) AddHand(hand Hand) {
	c.Hands = append(c.Hands, hand)
}

func (c *CamelCards) SortHands() {
	sort.Slice(c.Hands, func(i, j int) bool {
		return c.Hands[i].Rank < c.Hands[j].Rank
	})
	sort.Slice(c.Hands, func(i, j int) bool {
		if c.Hands[i].Rank == c.Hands[j].Rank {
			for k := 0; k < len(c.Hands[i].cards); k++ {
				if c.Hands[i].cards[k].id != c.Hands[j].cards[k].id {
					return c.Hands[i].cards[k].id < c.Hands[j].cards[k].id
				}
			}
		}
		return c.Hands[i].Rank < c.Hands[j].Rank
	})
}
