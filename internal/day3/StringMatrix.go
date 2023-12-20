package day3

import (
	"strconv"
	"strings"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

type StringMatrix struct {
	matrix [][]string
	nrows  int
	ncols  int
}

func (s StringMatrix) FromLines(lines []string) StringMatrix {
	matrix := [][]string{}
	for _, line := range lines {
		row := []string{}
		for _, char := range line {
			row = append(row, string(char))
		}
		matrix = append(matrix, row)
	}
	return StringMatrix{matrix: matrix, nrows: len(matrix), ncols: len(matrix[0])}
}

func (s StringMatrix) getAdjacent(position []int) [][]int {
	nrows := s.nrows
	ncols := s.ncols

	return [][]int{
		{utils.Max(position[0]-1, 0), position[1]},                             // up
		{utils.Min(position[0]+1, nrows-1), position[1]},                       // down
		{position[0], utils.Max(position[1]-1, 0)},                             // left
		{position[0], utils.Min(position[1]+1, ncols-1)},                       // right
		{utils.Max(position[0]-1, 0), utils.Max(position[1]-1, 0)},             // up-left
		{utils.Max(position[0]-1, 0), utils.Min(position[1]+1, ncols-1)},       // up-right
		{utils.Min(position[0]+1, nrows-1), utils.Max(position[1]-1, 0)},       // down-left
		{utils.Min(position[0]+1, nrows-1), utils.Min(position[1]+1, ncols-1)}, // down-right
	}
}

func (s StringMatrix) Print(params ...string) {
	replace := false
	original := ""
	replacement := ""

	if len(params) == 2 {
		replace = true
		original = params[0]
		replacement = params[1]
	}

	for _, line := range s.matrix {
		for _, char := range line {
			if replace {
				char = strings.Replace(char, original, replacement, -1)
			}
			print(char)
		}
		println()
	}
}

func (s StringMatrix) getNumbers() []PartNumber {
	numbers := []PartNumber{}
	lastCharWasNumber := false

	for i, line := range s.matrix {
		number := newPartNumber()

		for j := 0; j < len(line); j++ {
			char := string(line[j])
			if _, err := strconv.Atoi(char); err == nil {
				number.addCharToString(char, i, j)
				lastCharWasNumber = true
			} else {
				if lastCharWasNumber {
					numbers = append(numbers, number)
					number = newPartNumber()
					lastCharWasNumber = false
				}
			}

			if j == len(line)-1 && lastCharWasNumber {
				numbers = append(numbers, number)
				number = newPartNumber()
				lastCharWasNumber = false
			}
		}

	}
	return numbers
}

func (s StringMatrix) GetGears() []Gear {
	gears := []Gear{}
	for i, line := range s.matrix {
		for j := 0; j < len(line); j++ {
			char := string(line[j])
			if char == "*" {
				gears = append(gears, newGear(i, j))
			}
		}
	}
	return gears
}
