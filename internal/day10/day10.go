package day10

import (
	"fmt"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

type Point struct {
	X int
	Y int
}

type Graph struct {
	Nodes map[Point]Node
	Start Point
}

type NodeType string

type Node struct {
	Value      NodeType
	Point      Point
	Neighbours []Point
}

const (
	Ground     NodeType = "."
	Vertical   NodeType = "|"
	Horizontal NodeType = "-"
	NorthEast  NodeType = "L"
	NorthWest  NodeType = "J"
	SouthEast  NodeType = "7"
	SouthWest  NodeType = "F"
	Start      NodeType = "S"
)

func NewNode(value string, point Point) Node {
	switch value {
	case ".":
		return Node{Value: Ground, Point: point}
	case "|":
		neighbours := []Point{
			{point.X, point.Y - 1},
			{point.X, point.Y + 1},
		}
		return Node{Value: Vertical, Point: point, Neighbours: neighbours}
	case "-":
		neighbours := []Point{
			{point.X - 1, point.Y},
			{point.X + 1, point.Y},
		}
		return Node{Value: Horizontal, Point: point, Neighbours: neighbours}
	case "L":
		neighbours := []Point{
			{point.X + 1, point.Y},
			{point.X, point.Y - 1},
		}
		return Node{Value: NorthEast, Point: point, Neighbours: neighbours}
	case "J":
		neighbours := []Point{
			{point.X, point.Y - 1},
			{point.X - 1, point.Y},
		}
		return Node{Value: NorthWest, Point: point, Neighbours: neighbours}
	case "7":
		neighbours := []Point{
			{point.X - 1, point.Y},
			{point.X, point.Y + 1},
		}
		return Node{Value: SouthEast, Point: point, Neighbours: neighbours}
	case "F":
		neighbours := []Point{
			{point.X, point.Y + 1},
			{point.X + 1, point.Y},
		}
		return Node{Value: SouthWest, Point: point, Neighbours: neighbours}
	case "S":
		neighbours := []Point{
			{point.X, point.Y - 1},
			{point.X, point.Y + 1},
			{point.X - 1, point.Y},
			{point.X + 1, point.Y},
		}
		return Node{Value: Start, Point: point, Neighbours: neighbours}
	default:
		panic("Unknown tile type")
	}
}

func NewGraph(lines []string) Graph {
	g := Graph{}
	tiles := map[Point]Node{}
	start := Point{}
	for i := range lines {
		for j := range lines[i] {
			tile := NewNode(string(lines[i][j]), Point{j, i})
			if tile.Value != Ground {
				tiles[Point{j, i}] = tile
			}
			if tile.Value == Start {
				start = Point{j, i}
			}
		}
	}
	g.Start = start
	g.Nodes = tiles
	return g
}

func (m Graph) String() string {
	str := fmt.Sprintf("Start: %v\n\n", m.Start)
	for _, node := range m.Nodes {
		str += fmt.Sprintf("%v (%v) -> %v\n", node, node.Point, node.Neighbours)
	}
	return str
}

func (g Graph) Walk(prev Point, current Point, path []Point) []Point {
	if len(g.Nodes[current].Neighbours) == 0 || g.Start.X == current.X && g.Start.Y == current.Y {
		return path
	}
	for _, n := range g.Nodes[current].Neighbours {
		if n.X == prev.X && n.Y == prev.Y {
			continue
		}
		return g.Walk(current, n, append(path, current))
	}
	path = append(path, current)
	return g.Walk(current, g.Nodes[current].Neighbours[0], path)
}

func (g Graph) GetAllPaths() [][]Point {
	paths := [][]Point{}
	start := g.Start
	for _, n := range g.Nodes[start].Neighbours {
		path := g.Walk(start, n, []Point{start})
		paths = append(paths, path)
	}
	return paths
}

func Part1(lines []string) int {
	sum := 0
	g := NewGraph(lines)
	paths := g.GetAllPaths()

	for _, path := range paths {
		if len(path) > sum {
			sum = len(path)
		}
	}

	return sum / 2
}

func Part2(lines []string) int {
	sum := 0
	return sum
}

func Run() {
	lines := utils.LoadText("./data/day10.txt")

	println()
	println("##############################")
	println("#           Day 10           #")
	println("##############################")
	println()
	println("Part 1")
	println(Part1(lines))

	println("Part 2")
	println(Part2(lines))
}
