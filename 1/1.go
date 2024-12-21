package main

import (
	"aoc2024/utils"
	"math"
	"sort"
)

func GetLocationDistances(locationsFile string) int {
	aList := []int{}
	bList := []int{}
	locations := utils.ReadLines(locationsFile)
	for _, location := range locations {
		nums := utils.SplitLineIntoNumbers(location, `\s+`)
		aList = append(aList, nums[0])
		bList = append(bList, nums[1])
	}

	sort.Ints(aList)
	sort.Ints(bList)

	result := 0
	for i := 0; i < len(aList); i++ {
		result += int(math.Abs(float64(aList[i]) - float64(bList[i])))
	}

	return result
}

func GetSimilarityScore(locationsFile string) int {
	aList := []int{}
	bList := []int{}
	locations := utils.ReadLines(locationsFile)
	for _, location := range locations {
		nums := utils.SplitLineIntoNumbers(location, `\s+`)
		aList = append(aList, nums[0])
		bList = append(bList, nums[1])
	}

	sort.Ints(aList)
	sort.Ints(bList)

	similarity := 0
	for _, a := range aList {
		for _, b := range bList {
			if a == b {
				similarity += a
			}
		}
	}

	return similarity
}
