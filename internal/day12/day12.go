package day12

import (
	"strconv"
	"strings"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

func clearMap(m map[[3]int]int) {
	for k := range m {
		delete(m, k)
	}
}

func toIntSlice(s []string) []int {
	res := []int{}
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}
	return res
}

func count(lava []string, springs []int) int {
	currentState := map[[3]int]int{{0, 0, 0}: 1}
	nextState := map[[3]int]int{}

	for i := 0; i < len(lava); i++ {
		value := lava[i]
		for state, count := range currentState {
			idx, spring, exponent := state[0], state[1], state[2]
			switch {
			case (value == "#" || value == "?") && idx < len(springs) && exponent == 0:
				if value == "?" && spring == 0 {
					nextState[[3]int{idx, spring, exponent}] += count
				}
				spring++
				if spring == springs[idx] {
					idx++
					spring = 0
					exponent = 1
				}
				nextState[[3]int{idx, spring, exponent}] += count
			case (value == "." || value == "?") && spring == 0:
				exponent = 0
				nextState[[3]int{idx, spring, exponent}] += count
			}
		}
		currentState, nextState = nextState, currentState
		clearMap(nextState)
	}

	sum := 0
	for s, count := range currentState {
		if s[0] == len(springs) {
			sum += count
		}
	}

	return sum
}

func Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		_lava, _springs, _ := strings.Cut(line, " ")
		springs := toIntSlice(strings.Split(_springs, ","))
		sum += count(strings.Split(_lava, ""), springs)
	}
	return sum
}

func Part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		_lava, _springs, _ := strings.Cut(line, " ")
		newLava, newSprings := "", ""
		for i := 0; i < 5; i++ {
			newLava = newLava + _lava + "?"
			newSprings = newSprings + _springs + ","
		}
		newLava = strings.TrimSuffix(newLava, "?")
		newSprings = strings.TrimSuffix(newSprings, ",")
		sum += count(strings.Split(newLava, ""), toIntSlice(strings.Split(newSprings, ",")))
	}
	return sum
}

func Run() {
	lines := utils.LoadText("./data/day12.txt")

	println()
	println("##############################")
	println("#           Day 12           #")
	println("##############################")
	println()
	println("Part 1")
	println(Part1(lines))

	println("Part 2")
	println(Part2(lines))
}
