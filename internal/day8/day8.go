package day8

import (
	"fmt"
	"math"
	"strings"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

type Node struct {
	Left    string
	Right   string
	isStart bool
	isEnd   bool
}

type Tree map[string]Node

func (t Tree) String() string {
	s := ""
	for k, v := range t {
		s += k + " -> " + v.Left + ", " + v.Right + "\n"
	}
	return s
}

func NewNode(src string, dest string, isStart bool, isEnd ...bool) Node {
	if len(isEnd) > 0 {
		return Node{src, dest, isStart, isEnd[0]}
	}
	return Node{src, dest, isStart, false}
}

func (n *Node) String() string {
	return n.Left + ", " + n.Right
}

func (n *Node) SetEnd() {
	n.isEnd = true
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
		isStart := strings.HasSuffix(name, "A")
		isEnd := strings.HasSuffix(name, "Z")
		src = strings.Replace(src, "(", "", -1)
		src = strings.Replace(src, ")", "", -1)
		dest = strings.Replace(dest, "(", "", -1)
		dest = strings.Replace(dest, ")", "", -1)
		tree[name] = NewNode(src, dest, isStart, isEnd)
	}

	return tree
}

func (t Tree) Walk(directions string, start string) int {
	current := start
	steps := 0
	i := 0
	isEnd := false
	for !isEnd {
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
		isEnd = t[current].isEnd
	}
	return steps
}

func (t Tree) Walk2(directions string, start []string) int {
	current := start
	steps := 0
	i := 0
	isEnd := false
	found := map[string]bool{}
	for !isEnd {
		curi := i % len(directions)
		d := directions[curi]
		newCurrent, newSteps := t.TakeStep(current, d, false)
		steps += newSteps
		current = newCurrent
		for _, c := range current {
			currentIsEnd := t[c].isEnd
			if currentIsEnd {
				found[c] = true
			}
		}
		if len(found) == len(start) {
			isEnd = true
		}

		i++
		if i > 100000000 {
			panic("Too many steps")
		}
	}
	return steps
}

func (t Tree) TakeStep(current []string, direction byte, verbose ...bool) ([]string, int) {
	steps := 0
	if len(verbose) > 0 && verbose[0] {
		fmt.Printf("%v %v", current, direction)
	}
	newCurrent := []string{}
	for _, c := range current {
		steps++
		if direction == byte('L') {
			newCurrent = append(newCurrent, t[c].Left)
		} else {
			newCurrent = append(newCurrent, t[c].Right)
		}
	}
	return newCurrent, steps
}

func Part1(lines []string) int {
	directions := lines[0]
	tree := NewTree(lines[2:])
	return tree.Walk(directions, "AAA")
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

type Array []int

type ReducerFn func(prev int, next int) int

func (arr Array) Reduce(fn ReducerFn, i int) int {
	prev := i
	for _, v := range arr {
		prev = fn(prev, v)
	}
	return prev
}

func Part2(lines []string) int {
	directions := lines[0]
	tree := NewTree(lines[2:])
	startNodes := []string{}
	for k := range tree {
		if strings.HasSuffix(k, "A") {
			startNodes = append(startNodes, k)
		}
	}

	nStart := len(startNodes)
	nFound := 0
	count := 0
	values := Array{}
	for nFound < nStart {
		for _, dir := range directions {
			for i := 0; i < len(startNodes); i++ {
				if dir == 'L' {
					startNodes[i] = tree[startNodes[i]].Left
				} else {
					startNodes[i] = tree[startNodes[i]].Right
				}
			}
			count++
			for i, node := range startNodes {
				if strings.HasSuffix(node, "Z") {
					nFound += 1
					startNodes = append(startNodes[:i], startNodes[i+1:]...)
					values = append(values, count)
				}
			}
		}
	}

	minCommonMultiple := math.MaxInt32
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			minCommonMultiple = utils.Min(minCommonMultiple, lcm(values[i], values[j]))
		}
	}

	return values.Reduce(func(prev int, next int) int {
		return lcm(prev, next)
	}, 1)
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
