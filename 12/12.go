package main

import (
	"aoc2024/utils"
)

type Polygon struct {
	Area      int
	Perimeter int
	Sides     int
}

func CountFencesPrice(filename string, withDiscount bool) int {
	input := utils.GridFromSlice(utils.ReadLines(filename))
	return GetFencePrice(input, withDiscount)
}

func getSameKindSurroundingNeighbors(input utils.Grid[string], current utils.Point) []utils.Point {
	neighbors := make([]utils.Point, 0)

	// left
	if current.X > 0 && input.SamePoints(current.Left(), current) {
		neighbors = append(neighbors, current.Left())
	}

	// right
	if current.X < input.Width()-1 && input.SamePoints(current.Right(), current) {
		neighbors = append(neighbors, current.Right())
	}

	// up
	if current.Y > 0 && input.SamePoints(current.Up(), current) {
		neighbors = append(neighbors, current.Up())
	}

	// down
	if current.Y < input.Height()-1 && input.SamePoints(current.Down(), current) {
		neighbors = append(neighbors, current.Down())
	}
	return neighbors
}

// checkSurroundings checks if there are any corners on the current coordinates. It checks for both, outside and inside corners
func checkSurroundings(input utils.Grid[string], current utils.Point) int {
	return countIfIsGridEdge(input, current) +
		countTopLeft(input, current) +
		countTopRight(input, current) +
		countBottomLeft(input, current) +
		countBottomRight(input, current)
}

func countBottomRight(input utils.Grid[string], p utils.Point) (count int) {
	// outside corner
	if (p.X < input.Width()-1 && p.Y < input.Height()-1 && !input.SamePoints(p.Right(), p) && !input.SamePoints(p.Down(), p)) ||
		(p.X < input.Width()-1 && p.Y == input.Height()-1 && !input.SamePoints(p.Right(), p)) ||
		(p.X == input.Width()-1 && p.Y < input.Height()-1 && !input.SamePoints(p.Down(), p)) {
		count += 1
	}

	// inside corner
	if p.X > 0 && p.Y > 0 && input.SamePoints(p.Left(), p) && input.SamePoints(p.Up(), p) && !input.SamePoints(p.UpLeft(), p) {
		count += 1
	}
	return count
}

func countBottomLeft(input utils.Grid[string], p utils.Point) (count int) {
	// outside corner
	if (p.X > 0 && p.Y < input.Height()-1 && !input.SamePoints(p.Left(), p) && !input.SamePoints(p.Down(), p)) ||
		(p.X > 0 && p.Y == input.Height()-1 && !input.SamePoints(p.Left(), p)) ||
		(p.X == 0 && p.Y < input.Height()-1 && !input.SamePoints(p.Down(), p)) {
		count += 1
	}

	// inside corner
	if p.X < input.Width()-1 && p.Y > 0 && input.SamePoints(p.Right(), p) && input.SamePoints(p.Up(), p) && !input.SamePoints(p.UpRight(), p) {
		count += 1
	}
	return count
}

func countTopRight(input utils.Grid[string], p utils.Point) (count int) {
	// outside corner
	if (p.X < input.Width()-1 && p.Y > 0 && !input.SamePoints(p.Right(), p) && !input.SamePoints(p.Up(), p)) ||
		(p.X < input.Width()-1 && p.Y == 0 && !input.SamePoints(p.Right(), p)) ||
		(p.X == input.Width()-1 && p.Y > 0 && !input.SamePoints(p.Up(), p)) {
		count += 1
	}

	// inside corner
	if p.X > 0 && p.Y < input.Height()-1 && input.SamePoints(p.Left(), p) && input.SamePoints(p.Down(), p) && !input.SamePoints(p.DownLeft(), p) {
		count += 1
	}
	return count
}

func countTopLeft(input utils.Grid[string], p utils.Point) (count int) {
	// outside corner
	if (p.X > 0 && p.Y > 0 && !input.SamePoints(p.Left(), p) && !input.SamePoints(p.Up(), p)) ||
		(p.X > 0 && p.Y == 0 && !input.SamePoints(p.Left(), p)) ||
		(p.X == 0 && p.Y > 0 && !input.SamePoints(p.Up(), p)) {
		count += 1
	}

	// inside corner
	if p.X < input.Width()-1 && p.Y < input.Height()-1 && input.SamePoints(p.Right(), p) && input.SamePoints(p.Down(), p) && !input.SamePoints(p.DownRight(), p) {
		count += 1
	}
	return count
}

func countIfIsGridEdge(input utils.Grid[string], p utils.Point) (count int) {
	if p.X == 0 && p.Y == 0 {
		count += 1
	}

	if p.X == 0 && p.Y == input.Height()-1 {
		count += 1
	}

	if p.X == input.Width()-1 && p.Y == input.Height()-1 {
		count += 1
	}

	if p.X == input.Width()-1 && p.Y == 0 {
		count += 1
	}
	return count
}

func GetFencePrice(input utils.Grid[string], withDiscount bool) (price int) {
	visitedCoordinates := utils.NewCache[utils.Point, struct{}]()

	for j, row := range input.All() {
		for i, _ := range row {
			if !visitedCoordinates.Exists(utils.Point{X: i, Y: j}) {
				next := []utils.Point{{X: i, Y: j}}
				shape := Polygon{}
				for len(next) != 0 {
					newShape, traverseNext := findAllFlowersInGarden(input, next[0], shape, visitedCoordinates)
					shape = newShape
					next = append(next, traverseNext...)
					next = next[1:]
				}
				if withDiscount {
					price += shape.Area * shape.Sides
				} else {
					price += shape.Area * shape.Perimeter
				}
			}
		}
	}
	return price
}

func findAllFlowersInGarden(input utils.Grid[string], current utils.Point, shape Polygon, visited utils.Cache[utils.Point, struct{}]) (Polygon, []utils.Point) {
	if visited.Exists(current) {
		return shape, []utils.Point{}
	}

	checkNext := getSameKindSurroundingNeighbors(input, current)

	// A lonely flower. All neighbors are other kind. Direct result Area 1, Perimeter 4, Sides 4
	if len(checkNext) == 0 {
		visited[current] = struct{}{}
		return Polygon{Area: 1, Perimeter: 4, Sides: 4}, []utils.Point{}
	}

	shape.Perimeter += 4 - len(checkNext)
	shape.Area += 1
	visited[current] = struct{}{}
	shape.Sides += checkSurroundings(input, current)

	return shape, checkNext
}
