package utils

import (
	"regexp"
	"strconv"
	"strings"
)

// SplitLineIntoNumbers takes a string and a delimiter string and returns a slice of integers.
// The line is split into parts by the delimiter, and each part is converted to an integer.
// The resulting slice of integers is then returned.
func SplitLineIntoNumbers(line string, delimiter string) []int {
	re := regexp.MustCompile(delimiter)
	parts := re.Split(line, -1)
	result := []int{}
	for _, part := range parts {
		i, _ := strconv.Atoi(part)
		result = append(result, i)
	}
	return result
}

// GridFromSlice takes a slice of strings and returns a grid of strings.
// Each string in the input slice is split into individual characters,
// and the resulting slices of characters are used as the Rows of the
// 2D grid.
func GridFromSlice(input []string) (grid Grid[string]) {
	grid = NewGrid[string]()
	for _, line := range input {
		grid.AppendAll(strings.Split(line, ""))
	}
	return
}
