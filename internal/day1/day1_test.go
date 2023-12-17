package day1_test

import (
	"testing"

	"github.com/jenspederm/advent-of-code/internal/day1"
)

func TestDay1(t *testing.T) {
	lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	line_answers := []int{
		29,
		83,
		13,
		24,
		42,
		14,
		76,
	}
	values := day1.NewCalibrationValues(lines)

	if len(lines) != len(line_answers) {
		t.Errorf("Expected %d, got %d", len(lines), len(line_answers))
	}

	for i := range values {
		values[i].ReplaceWordDigits()
		values[i].SumFirstAndLastDigit()
		if values[i].Value != line_answers[i] {
			t.Errorf("Expected %d, got %d", line_answers[i], values[i].Value)
		} else {
			t.Logf("Expected %d, got %d", line_answers[i], values[i].Value)
		}
	}
}
