package day4

import (
	"strconv"
	"strings"

	"github.com/cloudflare/cfssl/log"
	"github.com/jenspederm/advent-of-code/internal/utils"
)

type Card struct {
	id      int
	name    string
	winning []int
	own     []int
}

func NewCard(id int, name string, winning []int, own []int) Card {
	log.Debugf("Creating card with id %d from %s", id, name)
	return Card{id: id, name: name, winning: winning, own: own}
}

func (c Card) Copy() Card {
	return Card{id: c.id, name: c.name, winning: c.winning, own: c.own}
}

func CardFromLine(line string) Card {
	cardData := strings.Split(line, ":")
	if len(cardData) != 2 {
		panic("Invalid cardData")
	}
	cardId, _ := strconv.Atoi(strings.Split(cardData[0], " ")[1])
	card := strings.Split(cardData[1], " | ")
	if len(card) != 2 {
		panic("Invalid card")
	}
	winning := utils.NumbersFromLine(card[0])
	own := utils.NumbersFromLine(card[1])
	return NewCard(cardId, cardData[0], winning, own)
}

func Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		card := CardFromLine(line)
		cardSum := 0

		for _, num := range card.own {
			for _, winningNum := range card.winning {
				if num == winningNum {
					if cardSum == 0 {
						cardSum += 1
					} else {
						cardSum *= 2
					}
				}
			}
		}

		sum += cardSum
	}
	return sum
}

func Part2(lines []string) int {
	sum := 0

	dict := map[int][]Card{}
	cards := []Card{}

	for _, line := range lines {
		card := CardFromLine(line)
		dict[card.id] = []Card{card}
		cards = append(cards, card)
	}
	i := 0
	for len(cards) > 0 {
		if len(cards) == 0 {
			break
		}
		card := cards[0]
		cards = cards[1:]
		log.Debugf("Checking card with id %d", card.id)
		log.Debugf("Number of cards remaining: %d", len(cards))
		numMatches := 0
		for _, num := range card.own {
			for _, winningNum := range card.winning {
				if num == winningNum {
					numMatches += 1
				}
			}
		}
		log.Debugf("Number of matches: %d", numMatches)
		if numMatches > 0 {
			for j := 1; j < numMatches+1; j++ {
				newId := card.id + j
				newCard := dict[newId][0].Copy()
				log.Debugf("Creating new card with id %d from %d", newId, card.id)
				dict[newCard.id] = append(dict[newCard.id], newCard)
				cards = append(cards, newCard)
			}
		}
		i += 1
	}

	for k := range dict {
		sum += len(dict[k])
	}

	return sum
}

func Run() {
	lines := utils.LoadText("./data/day4.txt")

	println()
	println("##############################")
	println("#            Day 4           #")
	println("##############################")
	println()
	println("Part 1")
	println(Part1(lines))

	println("Part 2")
	println(Part2(lines))
}
