package utils

import (
	"errors"
	"flag"
	"os"
)

func ParseArgs(def string) (string, error) {
	dataPtr := flag.String("data", def, "Path to data file")
	flag.Parse()

	args := os.Args[1:]

	for _, arg := range args {
		if arg == "-data" {
			println("Fetching data from: %s", *dataPtr)
			return *dataPtr, nil
		}
	}

	return "", errors.New("no data path specified")
}
