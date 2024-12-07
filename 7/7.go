package main

import (
	"aoc2024"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetValidEquationTotalSum(contentFile string, concatenate bool) int {
	lines := aoc2024.ReadLines(contentFile)
	result := 0
	for _, line := range lines {
		equation := parseEquation(line)
		if equation.hasValidOperatorCombination(concatenate) {
			result += equation.Result
		}
	}
	return result
}

type Equation struct {
	Result   int
	Operands []int
}

func parseEquation(line string) *Equation {
	splitLine := strings.Split(line, ": ")
	re := regexp.MustCompile(`(\d+)+`)
	matches := re.FindAllString(splitLine[1], -1)
	equation := &Equation{
		Result:   aoc2024.MustAtoi(splitLine[0]),
		Operands: make([]int, len(matches)),
	}
	for i, s := range matches {
		equation.Operands[i] = aoc2024.MustAtoi(s)
	}
	return equation
}

func (e *Equation) hasValidOperatorCombination(concatenate bool) bool {
	if len(e.Operands) == 0 {
		return false
	}
	return evaluateCombination(e.Operands, e.Result, e.Operands[0], 1, concatenate)
}

func evaluateCombination(operands []int, targetValue int, result int, idx int, concatenate bool) bool {
	if idx == len(operands) {
		return result == targetValue
	}
	if evaluateCombination(operands, targetValue, result+operands[idx], idx+1, concatenate) {
		return true
	}
	if evaluateCombination(operands, targetValue, result*operands[idx], idx+1, concatenate) {
		return true
	}
	if concatenate {
		concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", result, operands[idx]))
		return evaluateCombination(operands, targetValue, concatenated, idx+1, concatenate)
	}
	return false
}
