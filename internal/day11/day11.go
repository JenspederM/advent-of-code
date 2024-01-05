package day11

import (
	"math"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

type Point struct {
	row, col int
}

func CalculateDistance(p1, p2 Point) int {
	return int(math.Abs(float64(p1.row-p2.row)) + math.Abs(float64(p1.col-p2.col)))
}

func SumDistances(galaxies []Point) int {
	sum := 0

	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += CalculateDistance(galaxies[i], galaxies[j])
		}
	}

	return sum
}

func GetMinMax(m map[int]bool) (int, int) {
	min := 0
	max := 0
	for k := range m {
		if k < min {
			min = k
		}
		if k > max {
			max = k
		}
	}
	return min, max
}

func Expand(lines []string, nBetween int) []Point {
	galaxies := []Point{}
	colsWithGalaxies := map[int]bool{}
	rowsWithGalaxies := map[int]bool{}
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				colsWithGalaxies[j] = true
				rowsWithGalaxies[i] = true
				galaxies = append(galaxies, Point{i, j})
			}
		}
	}

	minRow, maxRow := GetMinMax(rowsWithGalaxies)
	minCol, maxCol := GetMinMax(colsWithGalaxies)
	rowsWithoutGalaxies := []int{}
	colsWithoutGalaxies := []int{}
	for i := minRow; i <= maxRow; i++ {
		if !rowsWithGalaxies[i] {
			rowsWithoutGalaxies = append(rowsWithoutGalaxies, i)
		}
	}

	for j := minCol; j <= maxCol; j++ {
		if !colsWithGalaxies[j] {
			colsWithoutGalaxies = append(colsWithoutGalaxies, j)
		}
	}

	expanded := []Point{}

	for _, galaxy := range galaxies {
		exp := galaxy
		for _, row := range rowsWithoutGalaxies {
			if row > galaxy.row {
				break
			}
			exp.row += nBetween - 1
		}
		for _, col := range colsWithoutGalaxies {
			if col > galaxy.col {
				break
			}
			exp.col += nBetween - 1
		}
		expanded = append(expanded, exp)
	}

	return expanded
}

func Part1(lines []string) int {
	expanded := Expand(lines, 2)
	return SumDistances(expanded)
}

func Part2(lines []string, nBetween ...int) int {
	if len(nBetween) > 0 {
		expanded := Expand(lines, nBetween[0])
		return SumDistances(expanded)
	}
	expanded := Expand(lines, 1_000_000)
	return SumDistances(expanded)
}

func Run() {
	lines := utils.LoadText("./data/day11.txt")

	println()
	println("##############################")
	println("#           Day 11           #")
	println("##############################")
	println()
	println("Part 1")
	println(Part1(lines))

	println("Part 2")
	println(Part2(lines))
}
