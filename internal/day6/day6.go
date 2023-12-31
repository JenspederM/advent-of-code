package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

func SumForTime(time int, distance int) int {
	sum := 0
	for x := 0; x < time+1; x++ {
		dx := x * (time - x)
		if dx > distance {
			sum += 1
		}
	}
	return sum
}

func ParseInput(lines []string) ([]int, []int) {
	times := utils.NumbersFromLine(strings.Split(lines[0], ":")[1])
	distances := utils.NumbersFromLine(strings.Split(lines[1], ":")[1])
	return times, distances
}

func ConcatListToInt(list []int) int {
	str := ""
	for i := range list {
		str += strings.TrimSpace(fmt.Sprint(list[i]))
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		panic("Paniced when parsing " + str + " " + err.Error())
	}
	return n
}

func Part1(lines []string) int {
	times, distances := ParseInput(lines)
	sum := 1
	for i := range times {
		sum *= SumForTime(times[i], distances[i])
	}
	return sum
}

func Part2(lines []string) int {
	times, distances := ParseInput(lines)
	fullTimes := ConcatListToInt(times)
	fullDistances := ConcatListToInt(distances)
	return SumForTime(fullTimes, fullDistances)
}

func Run() {
	lines := utils.LoadText("./data/day6.txt")

	println()
	println("##############################")
	println("#            Day 6           #")
	println("##############################")
	println()
	println("Part 1")
	println(Part1(lines))

	println("Part 2")
	println(Part2(lines))
}
