package day3

import (
	"strconv"
	"strings"
)

type Number struct {
	_string    []string
	_positions [][]int
	number     int
	isValid    bool
}

func newNumber() Number {
	validNumber := Number{}
	validNumber._string = []string{}
	validNumber.number = 0
	validNumber.isValid = false
	return validNumber
}

func (v *Number) addCharToString(char string, i int, j int) bool {
	v._string = append(v._string, char)
	v._positions = append(v._positions, []int{i, j})
	number, err := strconv.Atoi(strings.Join(v._string, ""))
	if err != nil {
		return false
	}
	v.number = number
	return true
}
