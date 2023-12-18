package day3

import (
	"github.com/jenspederm/advent-of-code/internal/utils"
)

func Part1(lines []string) int {
	s := StringMatrix{}.FromLines(lines)
	sum := 0
	numbers := s.getNumbers()

	for _, number := range numbers {
		isValid := s.CheckIsValidNumber(number)

		if isValid {
			sum += number.number
		}

		//println(number.number, isValid)
	}

	return sum
}

func Part2(lines []string) int {
	return 0
}

func Run() {
	lines := utils.LoadText("./data/day3.txt")

	println()
	println("##############################")
	println("#            Day 3           #")
	println("##############################")
	println()
	println("Part 1")
	println(Part1(lines))

	println("Part 2")
	println(Part2(lines))
}
