package main

import (
	"flag"

	"github.com/jenspederm/advent-of-code/internal/day1"
	"github.com/jenspederm/advent-of-code/internal/day10"
	"github.com/jenspederm/advent-of-code/internal/day11"
	"github.com/jenspederm/advent-of-code/internal/day2"
	"github.com/jenspederm/advent-of-code/internal/day3"
	"github.com/jenspederm/advent-of-code/internal/day4"
	"github.com/jenspederm/advent-of-code/internal/day5"
	"github.com/jenspederm/advent-of-code/internal/day6"
	"github.com/jenspederm/advent-of-code/internal/day7"
	"github.com/jenspederm/advent-of-code/internal/day8"
	"github.com/jenspederm/advent-of-code/internal/day9"
)

func main() {
	day := flag.String("day", "", "Day to run")
	flag.Parse()

	days := map[string]func(){
		"1":  day1.Run,
		"2":  day2.Run,
		"3":  day3.Run,
		"4":  day4.Run,
		"5":  day5.Run,
		"6":  day6.Run,
		"7":  day7.Run,
		"8":  day8.Run,
		"9":  day9.Run,
		"10": day10.Run,
		"11": day11.Run,
	}

	if *day == "" {
		for _, day := range days {
			day()
		}
	}

	println("Running day " + *day)
	days[*day]()
}
