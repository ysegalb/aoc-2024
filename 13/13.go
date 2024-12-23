package main

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
	"strings"
)

type Button struct {
	X     int
	Y     int
	Price int
}

type Machine struct {
	A             Button
	B             Button
	PriceLocation utils.Point
}

func (m Machine) print() {
	fmt.Printf("A: %d,%d\nB: %d,%d\nPrice: %d,%d\n", m.A.X, m.A.Y, m.B.X, m.B.Y, m.PriceLocation.X, m.PriceLocation.Y)
}

func GetTotalTokensForPrices(filename string, correctedCoordinates bool) (totalPrice int) {
	lines := utils.ReadLines(filename)
	machines, _ := parseMachineData(lines, correctedCoordinates)
	for _, machine := range machines {
		totalPrice += GetPriceCostForMachine(machine)
	}
	return totalPrice
}

func parseMachineData(lines []string, correctedCoordinates bool) ([]Machine, error) {
	var machines []Machine
	var machine Machine

	correction := 0
	if correctedCoordinates {
		correction = 10000000000000
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "Button A:") {
			machine = Machine{} // Reset machine for new entry
			re := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
			fields := re.FindAllStringSubmatch(line, -1)
			machine.A.X = utils.MustAtoi(fields[0][1])
			machine.A.Y = utils.MustAtoi(fields[0][2])
			machine.A.Price = 3
		} else if strings.HasPrefix(line, "Button B:") {
			re := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
			fields := re.FindAllStringSubmatch(line, -1)
			machine.B.X = utils.MustAtoi(fields[0][1])
			machine.B.Y = utils.MustAtoi(fields[0][2])
			machine.B.Price = 1
		} else if strings.HasPrefix(line, "Prize:") {
			re := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
			fields := re.FindAllStringSubmatch(line, -1)
			machine.PriceLocation.X = utils.MustAtoi(fields[0][1]) + correction
			machine.PriceLocation.Y = utils.MustAtoi(fields[0][2]) + correction
			machines = append(machines, machine) // Store machine in slice
		}
	}

	return machines, nil
}

func GetPriceCostForMachine(m Machine) int {
	pushesA, pushesB := m.solve()

	return m.A.Price*pushesA + m.B.Price*pushesB
}

func (m Machine) solve() (pushesA, pushesB int) {
	// We need to find if there is a combination of button pushes that solves an equation system:
	// a·x + b·y = c
	// d·x + e·y = f

	a := m.A.X
	b := m.B.X
	c := m.PriceLocation.X
	d := m.A.Y
	e := m.B.Y
	f := m.PriceLocation.Y

	// Using Cramer method
	det := (a * e) - (b * d)
	numeratorX := c*e - b*f
	numeratorY := a*f - c*d
	pushesA = numeratorX / det
	pushesB = numeratorY / det
	if numeratorX%det != 0 || numeratorY%det != 0 {
		// If either pushA or pushB are decimal numbers, the solution shouldn't be counted towards the total tokens, so no pushes are returned
		return 0, 0
	}
	return pushesA, pushesB
}
