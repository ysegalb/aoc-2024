package main

import (
	"aoc2024/utils"
	"fmt"
)

func GetGuardLocations(contentFile string, detectLoop bool) int {
	lines := utils.ReadLines(contentFile)
	board := newBoard(lines)
	if detectLoop {
		return board.findLoopPositions()
	}
	return board.countUniquePositions()
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Cell struct {
	Content byte
	Visited bool
}

type Board struct {
	Cells          [][]Cell
	GuardX, GuardY int
	GuardDir       Direction
}

func newBoard(lines []string) *Board {
	board := &Board{
		Cells: make([][]Cell, len(lines)),
	}
	for i, line := range lines {
		board.Cells[i] = make([]Cell, len(line))
		for j, ch := range line {
			board.Cells[i][j].Content = byte(ch)
			if ch == '^' || ch == 'v' || ch == '<' || ch == '>' {
				board.GuardX, board.GuardY = i, j
				switch ch {
				case '^':
					board.GuardDir = UP
				case 'v':
					board.GuardDir = DOWN
				case '<':
					board.GuardDir = LEFT
				case '>':
					board.GuardDir = RIGHT
				}
			}
		}
	}
	return board
}

func (b *Board) countUniquePositions() int {
	visited := make(map[string]bool)
	for {
		pos := fmt.Sprintf("%d,%d", b.GuardX, b.GuardY)
		visited[pos] = true
		if err := b.moveGuard(); err != nil {
			break
		}
	}
	return len(visited)
}

func (b *Board) findLoopPositions() int {
	possibleObstructions := 0
	for i := range b.Cells {
		for j := range b.Cells[i] {
			if b.Cells[i][j].Content == '.' {
				originalBoard := b.clone()
				b.Cells[i][j].Content = '#'
				if b.detectLoop() {
					possibleObstructions++
				}
				*b = *originalBoard
			}
		}
	}
	return possibleObstructions
}

func (b *Board) detectLoop() bool {
	states := make(map[string]bool)
	for {
		state := b.getState()
		if states[state] {
			return true
		}
		states[state] = true
		if err := b.moveGuard(); err != nil {
			return false
		}
	}
}

func (b *Board) getState() string {
	return fmt.Sprintf("%d,%d,%d", b.GuardX, b.GuardY, b.GuardDir)
}

func (b *Board) moveGuard() error {
	newX, newY := b.GuardX, b.GuardY
	switch b.GuardDir {
	case UP:
		newX--
	case DOWN:
		newX++
	case LEFT:
		newY--
	case RIGHT:
		newY++
	}

	if !b.isWithinBounds(newX, newY) {
		return fmt.Errorf("out of bounds")
	}

	if b.Cells[newX][newY].Content == '#' {
		b.GuardDir = (b.GuardDir + 1) % 4
		return nil
	}

	b.GuardX, b.GuardY = newX, newY
	return nil
}

func (b *Board) isWithinBounds(x, y int) bool {
	return x >= 0 && x < len(b.Cells) && y >= 0 && y < len(b.Cells[0])
}

func (b *Board) clone() *Board {
	newBoard := &Board{
		Cells:    make([][]Cell, len(b.Cells)),
		GuardX:   b.GuardX,
		GuardY:   b.GuardY,
		GuardDir: b.GuardDir,
	}
	for i := range b.Cells {
		newBoard.Cells[i] = make([]Cell, len(b.Cells[i]))
		copy(newBoard.Cells[i], b.Cells[i])
	}
	return newBoard
}
