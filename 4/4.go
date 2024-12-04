package main

import "aoc2024"

func containsWord(x int, y int, dx int, dy int, searchFor string, lines Lines) bool {
	word := ""
	for i := 0; i < len(searchFor); i++ {
		newX, newY := x+i*dx, y+i*dy
		if newX < 0 || newX >= lines.width() || newY < 0 || newY >= lines.height() {
			return false
		}
		word += lines.getWord(newX, newY)
	}
	return word == searchFor || word == reverse(searchFor)
}

func countXMASOccurrences(lines Lines) int {
	const searchFor = "XMAS"
	result := 0
	// Check all positions and directions
	for y := 0; y < lines.height(); y++ {
		for x := 0; x < lines.width(); x++ {
			directions := [][2]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}}
			for _, dir := range directions {
				if containsWord(x, y, dir[0], dir[1], searchFor, lines) {
					result++
				}
			}
		}
	}

	return result
}

func countCrossMASShapes(lines Lines) int {
	const searchFor = "MAS"
	result := 0
	for y := 1; y < lines.height()-1; y++ {
		for x := 1; x < lines.width()-1; x++ {
			// Check top-left to bottom-right diagonal
			topLeft := containsWord(x-1, y-1, 1, 1, searchFor, lines)
			// Check top-right to bottom-left diagonal
			topRight := containsWord(x+1, y-1, -1, 1, searchFor, lines)

			if topLeft && topRight {
				result++
			}
		}
	}

	return result
}

func GetXmasOccurrenceCount(contentFile string, shapeSearch bool) int {
	lines := Lines{aoc2024.ReadLines(contentFile)}
	if shapeSearch {
		return countCrossMASShapes(lines)
	} else {
		return countXMASOccurrences(lines)
	}
}
