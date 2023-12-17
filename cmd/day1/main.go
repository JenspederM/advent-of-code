package main

import (
	"bufio"
	"os"

	"github.com/jenspederm/advent-of-code/internal"
)

func main() {
	file, err := os.Open("./data/day1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines = []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := internal.NewCalibrationValues(lines)

	println("Part 1")
	println(part1.Sum())

	println("Part 2")
	part2 := internal.NewCalibrationValues(lines)
	for i := range part2 {
		part2[i].ReplaceWordDigits()
	}
	println(part2.Sum())
}
