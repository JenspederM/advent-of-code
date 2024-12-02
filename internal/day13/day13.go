package day13

import (
	"strings"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

func GetParts(lines []string) [][]string {
	parts := [][]string{}
	part := []string{}
	for _, line := range lines {
		if line == "" {
			parts = append(parts, part)
			part = []string{}
		} else {
			part = append(part, line)
		}
	}
	parts = append(parts, part)
	return parts
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IsPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

type Matrix struct {
	Rows int
	Cols int
	Data [][]string
}

func NewMatrix(lines []string) Matrix {
	m := Matrix{}
	m.Rows = len(lines)
	m.Cols = len(lines[0])
	m.Data = [][]string{}
	for _, line := range lines {
		m.Data = append(m.Data, strings.Split(line, ""))
	}
	return m
}

func (m Matrix) Print() {
	println(m.Rows, m.Cols)
	for _, row := range m.Data {
		println(strings.Join(row, ""))
	}
}

func (m Matrix) Transpose() Matrix {
	t := Matrix{}
	t.Rows = m.Cols
	t.Cols = m.Rows
	t.Data = [][]string{}
	for i := 0; i < t.Rows; i++ {
		row := []string{}
		for j := 0; j < t.Cols; j++ {
			row = append(row, m.Data[j][i])
		}
		t.Data = append(t.Data, row)
	}
	return t
}

func Part1(lines []string) int {
	sum := 0
	parts := GetParts(lines)
	column_matches := map[int]int{}
	row_matches := map[int]int{}
	for _, part := range parts {
		m := NewMatrix(part)
		mt := m.Transpose()
		isMatch := false
		for row := 0; row < m.Rows; row++ {
			for col := 0; col < m.Cols; col++ {
				if IsPalindrome(strings.Join(m.Data[row][col:], "")) {
					column_matches[len(m.Data[row][col:])/2+1]++
					isMatch = true
					break
				} else if IsPalindrome(strings.Join(mt.Data[col][row:], "")) {
					row_matches[len(mt.Data[col][row:])/2+1]++
					isMatch = true
					break
				}
			}
			if isMatch {
				break
			}
		}
	}

	for k := range column_matches {
		println("col", k)
		sum += k
	}

	for k := range row_matches {
		println("row", k)
		sum += k * 100
	}

	return sum
}

func Part2(lines []string) int {
	sum := 0
	return sum
}

func Run() {
	lines := utils.LoadText("./data/day13.txt")

	println()
	println("##############################")
	println("#           Day 13           #")
	println("##############################")
	println()
	println("Part 1")
	println(Part1(lines))

	println("Part 2")
	println(Part2(lines))
}
