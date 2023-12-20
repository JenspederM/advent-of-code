package day3

import (
	"strconv"
)

type Gear struct {
	partNumbers []PartNumber
	position    []int
	ratio       int
}

func newGear(i int, j int) Gear {
	gear := Gear{}
	gear.partNumbers = []PartNumber{}
	gear.position = []int{i, j}
	gear.ratio = 0
	return gear
}

func (g Gear) ToString() string {
	return "Gear(" + strconv.Itoa(g.position[0]) + ", " + strconv.Itoa(g.position[1]) + ")"
}

func (g Gear) getAllAdjacent(matrix StringMatrix) [][]int {
	allAdjacent := [][]int{}
	posAdjacent := matrix.getAdjacent(g.position)
	for _, adjacent := range posAdjacent {
		partOfSelf := false
		if adjacent[0] == g.position[0] && adjacent[1] == g.position[1] {
			partOfSelf = true
			break
		}
		if !partOfSelf {
			allAdjacent = append(allAdjacent, adjacent)
		}
	}
	return allAdjacent
}

func (g Gear) CalculateRatio(matrix StringMatrix, numbers []PartNumber) int {
	allAdjacent := g.getAllAdjacent(matrix)
	nearbyNumbers := []PartNumber{}

	for _, number := range numbers {
		for _, adjacent := range allAdjacent {
			i := adjacent[0]
			j := adjacent[1]
			found := false
			for _, pos := range number._positions {
				if i == pos[0] && j == pos[1] {
					nearbyNumbers = append(nearbyNumbers, number)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}

	count := len(nearbyNumbers)

	if count == 2 {
		g.ratio = nearbyNumbers[0].number * nearbyNumbers[1].number
		//println("Valid Gear at", g.ToString(), "has", count, "adjacent numbers", g.ratio)
		return g.ratio
	}

	return -1

}

func (g Gear) Ratio() int {
	if len(g.partNumbers) != 2 {
		return 0
	}

	return g.partNumbers[0].number * g.partNumbers[1].number
}
