package main

import (
	"aoc2024"
	"fmt"
	"regexp"
	"strconv"
)

type Operation struct {
	a int
	b int
}

func (o *Operation) execute() int {
	return o.a * o.b
}

func (o *Operation) String() string {
	return fmt.Sprintf("mul[%d,%d]", o.a, o.b)
}

func parseLines(lines []string, instructionsEnabled bool) []Operation {
	operations := []Operation{}
	addOperation := true
	for _, line := range lines {
		re := regexp.MustCompile(`(do\(\))|(don't\(\))|(mul\((\d{1,3}),(\d{1,3})\))|\s+`)
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			if match == "do()" {
				addOperation = true
			} else if match == "don't()" {
				addOperation = false
			} else if match != "" {
				if addOperation || !instructionsEnabled {
					operations = append(operations, processOperation(match))
				}
			}
		}
	}
	return operations
}

func processOperation(match string) Operation {
	operation := Operation{}
	mulRe, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})`)
	for _, mul := range mulRe.FindAllStringSubmatch(match, -1) {
		a, _ := strconv.Atoi(mul[1])
		b, _ := strconv.Atoi(mul[2])
		operation = Operation{a: a, b: b}
	}
	return operation
}

func GetMultiplicationAddedTotal(fileName string, instructionsEnabled bool) int {
	lines := aoc2024.ReadLines(fileName)
	operations := parseLines(lines, instructionsEnabled)

	result := 0
	for _, operation := range operations {
		result += operation.execute()
	}

	return result
}
