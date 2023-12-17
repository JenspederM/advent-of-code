package day1

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

type calibrationValue struct {
	Line  string
	Value int
}

type match struct {
	index int
	key   string
	value string
}

type CalibrationValues []calibrationValue

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func NewCalibrationValues(lines []string) CalibrationValues {
	calibration_values := make(CalibrationValues, 0, len(lines))

	for _, line := range lines {
		calibration_values = append(calibration_values, newCalibrationValue(line))
	}

	return calibration_values
}

func (c CalibrationValues) Sum() int {
	sum := 0

	for _, calibration_value := range c {
		calibration_value.SumFirstAndLastDigit()
		sum += calibration_value.Value
	}

	return sum
}

func (c *calibrationValue) SumFirstAndLastDigit() {
	var nbrs = []string{}

	for i := 0; i < len(c.Line); i++ {
		value := string(c.Line[i])
		if _, err := strconv.Atoi(value); err == nil {
			nbrs = append(nbrs, value)
		}
	}

	value, err := strconv.Atoi(nbrs[0] + nbrs[len(nbrs)-1])
	if err != nil {
		panic(err)
	}
	c.Value = value
}

func (c *calibrationValue) ReplaceWordDigits() {
	matches := []match{}

	for word, digit := range digitMap {
		for i := 0; i < len(c.Line); i++ {
			if i+len(word) > len(c.Line) {
				break
			}
			if c.Line[i:i+len(word)] == word {
				matches = append(matches, match{i, word, digit})
			}
		}
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].index < matches[j].index
	})

	if len(matches) == 0 {
		return
	}

	first_match := matches[0]
	last_match := matches[len(matches)-1]
	new_line := strings.Replace(c.Line, first_match.key, first_match.value, 1)
	new_line = strings.Replace(new_line, last_match.key, last_match.value, 1)
	c.Line = new_line
}

func Part1(lines []string) int {
	calibration_values := NewCalibrationValues(lines)
	return calibration_values.Sum()
}

func Part2(lines []string) int {
	calibration_values := NewCalibrationValues(lines)
	for i := range calibration_values {
		calibration_values[i].ReplaceWordDigits()
	}
	return calibration_values.Sum()
}

func Run() {
	lines := utils.LoadText("./data/day1.txt")

	println()
	println("##############################")
	println("#            Day 1           #")
	println("##############################")
	println()
	println("Part 1")
	println(Part1(lines))

	println("Part 2")
	println(Part2(lines))
}

func newCalibrationValue(line string) calibrationValue {
	return calibrationValue{line, 0}
}
