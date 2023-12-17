package day2

import (
	"strconv"
	"strings"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

var GameSettings = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type game struct {
	Name      string
	Number    int
	Rounds    []round
	IsTooHigh bool
}

type round struct {
	Picks []pick
}

type pick struct {
	Count int
	Color string
}

func newPick(count int, color string) pick {
	return pick{Count: count, Color: color}
}

func newRound(picks []pick) round {
	round := round{}
	round.Picks = picks
	return round
}

func newGame(number int, rounds []round, isTooHigh bool) game {
	game := game{}
	game.Name = "Game " + strconv.Itoa(number)
	game.Number = number
	game.Rounds = rounds
	game.IsTooHigh = isTooHigh
	return game
}

func GamesFromLines(lines []string, settings map[string]int) []game {
	games := make([]game, 0, len(lines))

	for _, line := range lines {
		data := strings.Split(line, ":")
		if len(data) != 2 {
			panic("Expected 2, got " + strconv.Itoa(len(data)))
		}
		game := strings.TrimSpace(data[0])
		if game == "" {
			panic("Expected game, got " + game)
		}
		game_number, err := strconv.Atoi(game[5:])
		isTooHigh := false
		if err != nil {
			panic(err)
		}
		rounds := strings.Split(data[1], ";")
		if len(rounds) < 1 {
			panic("Expected at least 1 round, got " + strconv.Itoa(len(rounds)))
		}
		newRounds := make([]round, 0, len(rounds))
		for _, round := range rounds {
			picks := strings.Split(round, ",")
			newPicks := make([]pick, 0, len(picks))
			if len(picks) < 1 {
				panic("Expected at least 1 pick, got " + strconv.Itoa(len(picks)))
			}
			for _, pick := range picks {
				pick = strings.TrimSpace(pick)
				if pick == "" {
					panic("Expected pick, got " + pick)
				}
				parts := strings.Split(pick, " ")
				if len(parts) != 2 {
					panic("Expected 2 parts, got " + strconv.Itoa(len(parts)))
				}
				count, err := strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}
				color := strings.TrimSpace(parts[1])
				if settings[color] < count {
					isTooHigh = true
				}
				newPicks = append(newPicks, newPick(count, color))
			}
			newRounds = append(newRounds, newRound(newPicks))
		}
		games = append(games, newGame(game_number, newRounds, isTooHigh))
		isTooHigh = false
	}
	return games
}

func Part1(lines []string) int {
	games := GamesFromLines(lines, GameSettings)

	sum := 0

	for _, game := range games {
		if !game.IsTooHigh {
			sum += game.Number
		}
	}

	return sum
}

func Part2(lines []string) int {
	games := GamesFromLines(lines, GameSettings)
	verbose := false
	total_sum := 0
	for _, game := range games {
		minimums := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, round := range game.Rounds {
			for _, pick := range round.Picks {
				if minimums[pick.Color] < pick.Count {
					minimums[pick.Color] = pick.Count
				}
			}
		}

		sum := 1
		for color, count := range minimums {
			if verbose {
				println(game.Name, color, count, "=>", sum, "*=", count, "=>", sum*count)
			}
			if count > 0 {
				sum *= count
			}
		}
		total_sum += sum
	}
	return total_sum
}

func Run() {
	lines := utils.LoadText("./data/day2.txt")

	println()
	println("##############################")
	println("#            Day 2           #")
	println("##############################")
	println()
	println("Part 1")
	println(Part1(lines))

	println("Part 2")
	println(Part2(lines))
}
