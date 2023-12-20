package utils

import "strconv"

func NumbersFromLine(line string) []int {
	numbers := []int{}
	lastCharWasNumber := false
	number := ""
	for i := range line {
		char := string(line[i])
		if _, err := strconv.Atoi(char); err == nil {
			number += char
			lastCharWasNumber = true
		} else {
			if lastCharWasNumber {
				num, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, num)
				number = ""
				lastCharWasNumber = false
			}
		}

		if i == len(line)-1 && lastCharWasNumber {
			num, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, num)
		}
	}
	return numbers
}
