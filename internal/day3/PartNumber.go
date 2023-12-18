package day3

import (
	"strconv"
	"strings"
)

type PartNumber struct {
	_string    []string
	_positions [][]int
	number     int
	isValid    bool
}

func newPartNumber() PartNumber {
	validNumber := PartNumber{}
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
