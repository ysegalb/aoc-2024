package aoc2024

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func SplitLineIntoNumbers(line string) []int {
	re := regexp.MustCompile(`\s+`)
	parts := re.Split(line, -1)
	result := []int{}
	for _, part := range parts {
		i, _ := strconv.Atoi(part)
		result = append(result, i)
	}
	return result
}
