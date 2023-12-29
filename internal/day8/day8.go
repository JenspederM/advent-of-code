package day8

import (
	"strings"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

type Node struct {
	Left  string
	Right string
}

type Tree map[string]Node

func (t Tree) String() string {
	s := ""
	for k, v := range t {
		s += k + " -> " + v.Left + ", " + v.Right + "\n"
	}
	return s
}

func NewTree(data []string) Tree {
	tree := Tree{}
	for _, line := range data {
		values := strings.Split(line, " = ")
		name := values[0]
		value := values[1]
		splits := strings.Split(value, ", ")
		src := splits[0]
		dest := splits[1]
		src = strings.Replace(src, "(", "", -1)
		src = strings.Replace(src, ")", "", -1)
		dest = strings.Replace(dest, "(", "", -1)
		dest = strings.Replace(dest, ")", "", -1)
		tree[name] = Node{src, dest}
	}
	return tree
}

func (t Tree) Walk(directions string) int {
	current := "AAA"
	steps := 0
	i := 0
	for current != "ZZZ" {
		d := directions[i]
		steps++
		i++
		if d == 'L' {
			current = t[current].Left
		} else {
			current = t[current].Right
		}
		if i >= len(directions) {
			i = 0
		}
	}
	return steps
}

func Part1(lines []string) int {
	directions := lines[0]
	tree := NewTree(lines[2:])

	// println(directions)
	// println(tree.String())

	return tree.Walk(directions)
}

func Part2(lines []string) int {
	return 0
}

func Run() {
	lines := utils.LoadText("./data/day8.txt")

	println()
	println("##############################")
	println("#            Day 8           #")
	println("##############################")
	println()
	println("Part 1")
	println(Part1(lines))

	println("Part 2")
	println(Part2(lines))
}
