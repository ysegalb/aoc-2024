package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	Height int
	North  *Node
	East   *Node
	South  *Node
	West   *Node
}

var VisitedNodes = make([]*Node, 0)
var Trails = make([][]*Node, 0)

func GetTrailheadScore(contentFile string, wholeSector bool) int {
	nodes, err := ReadNodesFromFile(contentFile)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	if wholeSector {
		return countTrails(nodes)
	} else {
		return searchTrails(nodes)
	}
}

func searchTrails(nodes []Node) int {
	totalScore := 0
	for _, node := range nodes {
		if node.Height == 0 {
			node.traverse()
			fmt.Printf("Visited nodes: %v\n", unique(VisitedNodes))
			totalScore += len(unique(VisitedNodes))
			VisitedNodes = make([]*Node, 0)
		}
	}
	return totalScore
}

func countTrails(nodes []Node) int {
	totalScore := 0
	for _, node := range nodes {
		if node.Height == 0 {
			Trails = make([][]*Node, 0)
			node.traverseTrail()
			totalScore += len(Trails)
		}
	}
	return totalScore
}

func (n *Node) traverse() {
	if n.Height == 9 {
		VisitedNodes = append(VisitedNodes, n)
		return
	}

	for _, dir := range []*Node{n.North, n.East, n.South, n.West} {
		if dir != nil && dir.Height == n.Height+1 {
			dir.traverse()
		}
	}

	return
}

func (n *Node) traverseTrail() {
	if n.Height == 9 {
		trail := make([]*Node, 0)
		for _, node := range VisitedNodes {
			trail = append(trail, node)
		}
		trail = append(trail, n)
		Trails = append(Trails, trail)
		return
	}

	VisitedNodes = append(VisitedNodes, n)
	for _, dir := range []*Node{n.North, n.East, n.South, n.West} {
		if dir != nil && dir.Height == n.Height+1 {
			dir.traverseTrail()
		}
	}
	VisitedNodes = VisitedNodes[:len(VisitedNodes)-1]

	return
}

func unique(slice []*Node) []Node {
	uniqueMap := make(map[string]bool)
	result := make([]Node, 0, len(slice))

	for _, item := range slice {
		if _, exists := uniqueMap[item.String()]; !exists {
			uniqueMap[item.String()] = true
			result = append(result, *item)
		}
	}

	return result
}

func (n *Node) String() string {
	return fmt.Sprintf("%p", n)
}

func ReadNodesFromFile(filename string) ([]Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var nodes [][]Node
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]Node, len(line))
		for i, char := range line {
			height, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			row[i] = Node{Height: height}
		}
		nodes = append(nodes, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Set North, East, South, and West pointers
	setSurroundingNodes(nodes)

	// Flatten the 2D slice into a 1D slice
	var flatNodes []Node
	for _, row := range nodes {
		flatNodes = append(flatNodes, row...)
	}

	return flatNodes, nil
}

func setSurroundingNodes(nodes [][]Node) {
	for i, row := range nodes {
		for j, node := range row {
			if i > 0 {
				node.North = &nodes[i-1][j]
			}
			if j < len(row)-1 {
				node.East = &nodes[i][j+1]
			}
			if i < len(nodes)-1 {
				node.South = &nodes[i+1][j]
			}
			if j > 0 {
				node.West = &nodes[i][j-1]
			}
			nodes[i][j] = node
		}
	}
}
