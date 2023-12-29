package main

import (
	"flag"
	"strings"

	"github.com/jenspederm/advent-of-code/internal/day1"
	"github.com/jenspederm/advent-of-code/internal/day2"
	"github.com/jenspederm/advent-of-code/internal/day3"
	"github.com/jenspederm/advent-of-code/internal/day4"
	"github.com/jenspederm/advent-of-code/internal/day5"
	"github.com/jenspederm/advent-of-code/internal/day6"
	"github.com/jenspederm/advent-of-code/internal/day7"
)

func main() {
	day := flag.String("day", "1", "Day to run")
	flag.Parse()

	days := map[string]func(){
		"1": day1.Run,
		"2": day2.Run,
		"3": day3.Run,
		"4": day4.Run,
		"5": day5.Run,
		"6": day6.Run,
		"7": day7.Run,
	}
	valid_days := []string{}
	for day := range days {
		valid_days = append(valid_days, day)
	}
	valid_msg := "Valid days are: " + strings.Join(valid_days, ", ")

	if *day == "" {
		panic("No day specified. " + valid_msg)
	}

	if _, ok := days[*day]; !ok {
		panic("Invalid day specified: " + *day + ". " + valid_msg)
	}

	println("Running day " + *day)
	days[*day]()
}
