package utils

import (
	"errors"
	"flag"
	"os"
	"strings"
)

func GetDataPath(def string) (string, error) {
	dataPtr := flag.String("data", def, "Path to data file")
	flag.Parse()

	switch arg := strings.Split(os.Args[1], "=")[0]; arg {
	case "-data":
		println("Fetching data from: %s", *dataPtr)
		return *dataPtr, nil
	default:
		println(arg)
		return "", errors.New("unknown command")
	}
}
