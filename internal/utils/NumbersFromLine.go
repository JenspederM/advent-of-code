package utils

import (
	"strconv"
	"strings"
)

func NumbersFromLine(line string, sep ...string) []int {
	seperator := " "
	if len(sep) > 0 {
		seperator = sep[0]
	}
	seq := []int{}
	for _, l := range strings.Split(strings.TrimSpace(line), seperator) {
		s, _ := strconv.Atoi(l)
		seq = append(seq, s)
	}
	return seq
}
