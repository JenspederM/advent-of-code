package day10

import (
	"fmt"

	"github.com/jenspederm/advent-of-code/internal/utils"
)

type Point struct {
	I int
	J int
}

type Graph struct {
	Nodes map[Point]Node
	Start Point
	nrows int
	ncols int
}

type NodeType string

type Node struct {
	Value      NodeType
	Point      Point
	Neighbours []Point
}

type Polygon struct {
	Points []Point
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
			{point.I - 1, point.J},
			{point.I + 1, point.J},
		}
		return Node{Value: Vertical, Point: point, Neighbours: neighbours}
	case "-":
		neighbours := []Point{
			{point.I, point.J - 1},
			{point.I, point.J + 1},
		}
		return Node{Value: Horizontal, Point: point, Neighbours: neighbours}
	case "L":
		neighbours := []Point{
			{point.I, point.J + 1},
			{point.I - 1, point.J},
		}
		return Node{Value: NorthEast, Point: point, Neighbours: neighbours}
	case "J":
		neighbours := []Point{
			{point.I - 1, point.J},
			{point.I, point.J - 1},
		}
		return Node{Value: NorthWest, Point: point, Neighbours: neighbours}
	case "7":
		neighbours := []Point{
			{point.I, point.J - 1},
			{point.I + 1, point.J},
		}
		return Node{Value: SouthEast, Point: point, Neighbours: neighbours}
	case "F":
		neighbours := []Point{
			{point.I + 1, point.J},
			{point.I, point.J + 1},
		}
		return Node{Value: SouthWest, Point: point, Neighbours: neighbours}
	case "S":
		neighbours := []Point{
			{point.I - 1, point.J},
			{point.I + 1, point.J},
			{point.I, point.J - 1},
			{point.I, point.J + 1},
		}
		return Node{Value: Start, Point: point, Neighbours: neighbours}
	default:
		panic("Unknown tile type")
	}
}

func NewGraph(lines []string) Graph {
	g := Graph{nrows: len(lines), ncols: len(lines[0])}
	tiles := map[Point]Node{}
	start := Point{}
	for i := range lines {
		for j := range lines[i] {
			tile := NewNode(string(lines[i][j]), Point{i, j})
			tiles[Point{i, j}] = tile
			if tile.Value == Start {
				start = Point{i, j}
			}
		}
	}
	g.Start = start
	g.Nodes = tiles
	return g
}

func (p Point) Compare(other Point) bool {
	return p.I == other.I && p.J == other.J
}

func (g Graph) String() string {
	str := fmt.Sprintf("Start: %v\n\n", g.Start)
	for _, node := range g.Nodes {
		str += fmt.Sprintf("%v (%v) -> %v\n", node, node.Point, node.Neighbours)
	}
	return str
}

func (g Graph) Walk(prev Point, current Point, path []Point) []Point {
	currentNode := g.Nodes[current]
	if len(currentNode.Neighbours) == 0 || currentNode.Value == Ground || g.Start.Compare(current) {
		return path
	}
	for _, next := range currentNode.Neighbours {
		if prev.Compare(next) {
			continue
		}
		return g.Walk(current, next, append(path, current))
	}
	path = append(path, current)
	return g.Walk(current, currentNode.Neighbours[0], path)
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

func (g Graph) PrintMatrix(contained []Point) string {
	str := fmt.Sprintf("Start: %v\n\n", g.Start)

	for i := 0; i < g.nrows; i++ {
		for j := 0; j < g.ncols; j++ {
			isContained := false
			for _, point := range contained {
				if point.Compare(Point{i, j}) {
					isContained = true
				}
			}
			if isContained {
				str += "I"
			} else {
				str += fmt.Sprintf("%v", g.Nodes[Point{i, j}].Value)
			}
		}
		str += "\n"
	}
	//println(str)
	return str
}

func (p Polygon) Contains(point Point) bool {
	for _, pp := range p.Points {
		if pp.I == point.I && pp.J == point.J {
			return false
		}
	}
	// https://stackoverflow.com/questions/217578/how-can-i-determine-whether-a-2d-point-is-within-a-polygon
	inside := false
	for i := 0; i < len(p.Points); i++ {
		j := (i + 1) % len(p.Points)
		if ((p.Points[i].I > point.I) != (p.Points[j].I > point.I)) &&
			(point.J < (p.Points[j].J-p.Points[i].J)*(point.I-p.Points[i].I)/(p.Points[j].I-p.Points[i].I)+p.Points[i].J) {
			inside = !inside
		}
	}
	return inside
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
	g := NewGraph(lines)
	paths := g.GetAllPaths()
	longestPath := []Point{}

	for _, path := range paths {
		if len(path) > len(longestPath) {
			longestPath = path
		}
	}

	polygon := Polygon{Points: longestPath}
	contained := []Point{}
	for _, node := range g.Nodes {
		if polygon.Contains(node.Point) {
			contained = append(contained, node.Point)
		}
	}

	g.PrintMatrix(contained)
	return len(contained)
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
