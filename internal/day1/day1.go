package day1

import (
	"sort"
	"strconv"
	"strings"
)

type CalibrationValue struct {
	line  string
	value int
}

type CalibrationValues []CalibrationValue

func NewCalibrationValue(line string) CalibrationValue {
	return CalibrationValue{line, 0}
}

func NewCalibrationValues(lines []string) CalibrationValues {
	calibration_values := make(CalibrationValues, 0, len(lines))

	for _, line := range lines {
		calibration_values = append(calibration_values, NewCalibrationValue(line))
	}

	return calibration_values
}

func (c CalibrationValues) Sum() int {
	sum := 0

	for _, calibration_value := range c {
		calibration_value.SumFirstAndLastDigit()
		sum += calibration_value.value
	}

	return sum
}

func (c *CalibrationValue) SumFirstAndLastDigit() {
	var nbrs = []string{}

	for i := 0; i < len(c.line); i++ {
		value := string(c.line[i])
		if _, err := strconv.Atoi(value); err == nil {
			nbrs = append(nbrs, value)
		}
	}

	value, err := strconv.Atoi(nbrs[0] + nbrs[len(nbrs)-1])
	if err != nil {
		panic(err)
	}
	c.value = value
}

type Match struct {
	index int
	key   string
	value string
}

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

func (c *CalibrationValue) ReplaceWordDigits() {
	matches := []Match{}

	for word, digit := range digitMap {
		for i := 0; i < len(c.line); i++ {
			if i+len(word) > len(c.line) {
				break
			}
			if c.line[i:i+len(word)] == word {
				matches = append(matches, Match{i, word, digit})
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
	new_line := strings.Replace(c.line, first_match.key, first_match.value, 1)
	new_line = strings.Replace(new_line, last_match.key, last_match.value, 1)
	c.line = new_line
}
