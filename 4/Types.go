package main

type Lines struct {
	lines []string
}

func (l *Lines) height() int {
	return len(l.lines)
}

func (l *Lines) width() int {
	return len(l.lines[0])
}

func (l *Lines) getWord(x, y int) string {
	return string(l.lines[y][x])
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
