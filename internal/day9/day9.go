package day9

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

func Diff(numbers []int) int {
	diff := make([]int, len(numbers)-1)
	for i := 0; i < len(numbers)-1; i++ {
		diff[i] = numbers[i+1] - numbers[i]
	}

	if !slices.ContainsFunc(diff, func(x int) bool { return x != 0 }) {
		return numbers[len(numbers)-1]
	}
	//fmt.Printf("Dropped %d from %v -> %v\n", numbers[0], numbers, diff)
	return numbers[len(numbers)-1] + Diff(diff)
}

func Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		seq := []int{}
		for _, l := range strings.Split(strings.TrimSpace(line), " ") {
			s, _ := strconv.Atoi(l)
			seq = append(seq, s)
		}
		sum += Diff(seq)
	}
	return sum
}

func Part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		seq := utils.NumbersFromLine(line)
		slices.Reverse(seq)
		sum += Diff(seq)
	}
	return sum
}

func Run() {
	input, _ := os.ReadFile("./data/day9.txt")
	lines := []string{}
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		lines = append(lines, strings.TrimSpace(line))
	}

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
