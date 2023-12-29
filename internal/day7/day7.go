package day7

import (
	"github.com/jenspederm/advent-of-code/internal/utils"
)

func Part1(lines []string) int {
	c := NewCamelCards(lines)
	return c.Sum(false)
}

func Part2(lines []string) int {
	c := NewCamelCards(lines, true)
	return c.Sum(true)
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
