package day3

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type PartNumber struct {
	_string    []string
	_positions [][]int
	Uid        string
	number     int
	isValid    bool
}

func newPartNumber() PartNumber {
	validNumber := PartNumber{}
	validNumber.Uid = uuid.New().String()
	validNumber._positions = [][]int{}
	validNumber._string = []string{}
	validNumber.number = 0
	validNumber.isValid = false
	return validNumber
}

func (v *PartNumber) addCharToString(char string, i int, j int) bool {
	v._string = append(v._string, char)
	v._positions = append(v._positions, []int{i, j})
	number, err := strconv.Atoi(strings.Join(v._string, ""))
	if err != nil {
		return false
	}
	v.number = number
	return true
}

func (n PartNumber) getAllAdjacent(matrix StringMatrix) [][]int {
	allAdjacent := [][]int{}
	for _, position := range n._positions {
		posAdjacent := matrix.getAdjacent(position)
		for _, adjacent := range posAdjacent {
			partOfSelf := false
			for _, pos := range n._positions {
				if adjacent[0] == pos[0] && adjacent[1] == pos[1] {
					partOfSelf = true
					break
				}
			}
			if !partOfSelf {
				allAdjacent = append(allAdjacent, adjacent)
			}
		}
	}
	return allAdjacent
}

func (n PartNumber) CheckIsValid(matrix StringMatrix) bool {
	allAdjacent := n.getAllAdjacent(matrix)
	isValid := false
	for _, adjacent := range allAdjacent {
		i := adjacent[0]
		j := adjacent[1]
		if matrix.matrix[i][j] != "." {
			isValid = true
			break
		}
	}
	n.isValid = isValid
	return isValid
}
