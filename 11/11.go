package main

import (
	"aoc2024"
	"fmt"
	"strconv"
	"strings"
)

func CountBlinkingStones(filename string, blinks int) int {
	numbers := strings.Split(aoc2024.ReadLines(filename)[0], " ")
	stones := make([]int, 0)
	for _, line := range numbers {
		stones = append(stones, aoc2024.MustAtoi(line))
	}
	sum := 0
	// Use the memoize technique. We store calculated values, so if we find the same value to process, we can return from the cache instead of calculating it again. We reduce enormously the time complexity of the recursion.
	cache := make(map[int][]int)
	for _, stone := range stones {
		sum += getCountAfterXBlinks(stone, cache, blinks)
	}
	return sum
}

func BlinkOnce(s int) []int {
	if s == 0 {
		return []int{1}
	}
	if len(fmt.Sprintf("%d", s))%2 == 0 {
		strStone := strconv.Itoa(s)
		half := len(strStone) / 2
		firstHalf, _ := strconv.Atoi(strStone[:half])
		secondHalf, _ := strconv.Atoi(strStone[half:])
		return []int{firstHalf, secondHalf}
	}
	return []int{s * 2024}
}

func getCountAfterXBlinks(stone int, cache map[int][]int, blinks int) int {
	// Check if we have the key loaded
	if _, ok := cache[stone]; ok {
		// Check if we have the blinks loaded
		if cache[stone][blinks-1] != 0 {
			return cache[stone][blinks-1]
		}
	} else {
		// We don't have the key loaded, so we create the key
		cache[stone] = make([]int, 75)
	}

	if blinks == 1 {
		// Store the value in the cache
		cache[stone][blinks-1] = len(BlinkOnce(stone))
		return len(BlinkOnce(stone))
	} else {
		// Store the value in the cache and calculate the next one
		sum := 0
		for _, stone := range BlinkOnce(stone) {
			sum += getCountAfterXBlinks(stone, cache, blinks-1)
		}
		cache[stone][blinks-1] = sum
		return sum
	}
}
