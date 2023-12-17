package main

import (
	"flag"
	"strings"

	"github.com/jenspederm/advent-of-code/internal/day1"
)

func main() {
	day := flag.String("day", "1", "Day to run")
	flag.Parse()

	days := map[string]func(){
		"1": day1.Run,
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
