package main

import (
	"aoc2024/utils"
)

type Report struct {
	Levels []int
}

func (r *Report) isSafe(withDampener ...bool) bool {
	if withDampener[0] {
		// Test exhaustively removing one level at a time
		for i := 0; i < len(r.Levels); i++ {
			newLevels := make([]int, 0)
			newLevels = append(newLevels, r.Levels[:i]...)
			newLevels = append(newLevels, r.Levels[i+1:]...)

			if safeReport(newLevels) {
				return true
			}
		}
	}
	return safeReport(r.Levels)
}
func safeReport(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	incremental := levels[1] > levels[0]

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff == 0 || utils.Abs(diff) > 3 {
			return false
		}
		if incremental && levels[i] < levels[i-1] {
			return false
		}
		if !incremental && levels[i] > levels[i-1] {
			return false
		}
	}
	return true
}

func getReports(contentFile string) []Report {
	lines := utils.ReadLines(contentFile)

	reports := []Report{}
	for _, line := range lines {
		reports = append(reports, Report{Levels: utils.SplitLineIntoNumbers(line, `\s+`)})
	}

	return reports
}

func GetSafeReports(contentFile string, dampener ...bool) int {
	reports := getReports(contentFile)

	safeReports := 0
	for _, report := range reports {
		if report.isSafe(dampener...) {
			safeReports++
		}
	}

	return safeReports
}
