package main

import (
	"aoc2024"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func GetPageOrdering(contentFile string, pickCorrectOnes bool) int {
	lines := aoc2024.ReadLines(contentFile)

	rules, prints := processLines(lines)

	result := 0
	for _, printRequest := range prints {
		if pickCorrectOnes {
			if printRequest.isValid(rules) {
				result += printRequest.getMiddleValue()
				continue
			}
		} else {
			if !printRequest.isValid(rules) {
				printRequest.sortAccordingToRules(rules)
				result += printRequest.getMiddleValue()
				continue
			}
		}
	}
	return result
}

type Print struct {
	values []int
}

func (p *Print) addValue(value string) {
	intVal, _ := strconv.Atoi(strings.Replace(value, ",", "", -1))
	p.values = append(p.values, intVal)
}

func (p *Print) addValues(values []string) {
	for _, value := range values {
		p.addValue(value)
	}
}

func (p *Print) isValid(rules map[int][]int) bool {
	prevPages := make([]int, 0)
	for _, value := range p.values {
		prevPages = append(prevPages, value)
		if rules[value] != nil {
			for _, postPage := range rules[value] {
				if slices.Contains(prevPages, postPage) {
					return false
				}
			}
		}
	}
	return true
}

func (p *Print) sortAccordingToRules(rules map[int][]int) {
	newValues := make([]int, 0)

VALUES:
	for _, value := range p.values {
		for j, newValue := range newValues {
			if slices.Contains(rules[value], newValue) {
				newValues = append(newValues[:j], append([]int{value}, newValues[j:]...)...)
				continue VALUES
			}
		}
		newValues = append(newValues, value)
	}
	p.values = newValues
}

func (p *Print) getMiddleValue() int {
	return p.values[len(p.values)/2]
}

func processLines(lines []string) (rules map[int][]int, prints []Print) {
	rules = make(map[int][]int)
	prints = make([]Print, 0)
	printNum := 0

	ruleRe := regexp.MustCompile(`(\d+)\|(\d+)`)
	printRe := regexp.MustCompile(`(\d+),?`)

	for _, line := range lines {
		if ruleRe.MatchString(line) {
			matches := ruleRe.FindAllStringSubmatch(line, -1)
			a, _ := strconv.Atoi(matches[0][1])
			b, _ := strconv.Atoi(matches[0][2])
			rules[a] = append(rules[a], b)
		} else if printRe.MatchString(line) {
			values := printRe.FindAllString(line, -1)
			newPrint := Print{}
			newPrint.addValues(values)
			prints = append(prints, newPrint)
			printNum++
		}
	}
	return rules, prints
}
