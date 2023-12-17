package main

import (
	"io/fs"
	"os"
	"strings"
)

func main() {
	day := os.Args[1]

	println("Creating day " + day)

	println("Creating folder structure")
	os.MkdirAll(strings.Join([]string{"./internal", "day" + day}, "/"), fs.ModePerm)

	dayFileExists, err := os.Stat(strings.Join([]string{"./internal", "day" + day, "day" + day + ".go"}, "/"))
	if err == nil && !dayFileExists.IsDir() {
		println("day" + day + ".go already exists")
	} else {
		println("Creating day" + day + ".go")
		dayFile, err := os.Create(strings.Join([]string{"./internal", "day" + day, "day" + day + ".go"}, "/"))
		if err != nil {
			panic(err)
		}
		defer dayFile.Close()

		dayFile.WriteString(strings.Join([]string{"package day" + day, ""}, "\n"))
	}

	dayTestFileExists, err := os.Stat(strings.Join([]string{"./internal", "day" + day, "day" + day + "_test.go"}, "/"))
	if err == nil && !dayTestFileExists.IsDir() {
		println("day" + day + "_test.go already exists")
	} else {
		println("Creating day" + day + "_test.go")
		dayTestFile, err := os.Create(strings.Join([]string{"./internal", "day" + day, "day" + day + "_test.go"}, "/"))
		if err != nil {
			panic(err)
		}
		defer dayTestFile.Close()
		dayTestFile.WriteString(strings.Join([]string{"package day" + day + "_test"}, "\n"))
	}
}
