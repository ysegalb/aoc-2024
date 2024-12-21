package main

import (
	"aoc2024/utils"
)

const EmptyFrequency = "."

type Cell struct {
	Frequency string
	Antinode  bool
}

type Board struct {
	Cells       [][]Cell
	Frequencies map[string][]utils.Point
}

func (b *Board) CountAntinodes() int {
	count := 0
	for _, row := range b.Cells {
		for _, cell := range row {
			if cell.Antinode {
				count++
			}
		}
	}
	return count
}

func GetTotalAntinodeCount(contentFile string, withHarmonics bool) int {
	lines := utils.ReadLines(contentFile)
	board := processFrequencies(lines, withHarmonics)
	board.calculateAntinodes(withHarmonics)
	return board.CountAntinodes()
}

func processFrequencies(lines []string, withHarmonics bool) Board {
	board := Board{
		Cells:       make([][]Cell, len(lines)),
		Frequencies: make(map[string][]utils.Point),
	}
	for i, line := range lines {
		board.Cells[i] = make([]Cell, len(line))
		for j, ch := range line {
			if EmptyFrequency != string(ch) {
				if _, ok := board.Frequencies[string(ch)]; !ok {
					board.Frequencies[string(ch)] = make([]utils.Point, 0)
				}
				board.Frequencies[string(ch)] = append(board.Frequencies[string(ch)], utils.Point{X: i, Y: j})
			}
			board.Cells[i][j].Frequency = string(ch)
			board.Cells[i][j].Antinode = withHarmonics && string(ch) != EmptyFrequency
		}
	}
	return board
}

func (b *Board) calculateAntinodes(withHarmonics bool) {
	for _, positions := range b.Frequencies {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				dx, dy := positions[j].X-positions[i].X, positions[j].Y-positions[i].Y
				x1, x2 := positions[i].X, positions[j].X
				y1, y2 := positions[i].Y, positions[j].Y
				inboundI, inboundJ := true, true
				for inboundI || inboundJ {
					x1 -= dx
					y1 -= dy
					x2 += dx
					y2 += dy
					inboundI = b.SetAntinodeAt(x1, y1) && withHarmonics
					inboundJ = b.SetAntinodeAt(x2, y2) && withHarmonics
				}
			}
		}
	}
}

func (b *Board) SetAntinodeAt(x, y int) bool {
	if x >= 0 && x < len(b.Cells) && y >= 0 && y < len(b.Cells[0]) {
		b.Cells[x][y].Antinode = true
		return true
	} else {
		return false
	}
}
