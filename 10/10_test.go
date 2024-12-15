package main

import "testing"

func TestTrailHeadsScore(t *testing.T) {
	testCases := []struct {
		caseName       string
		contentFile    string
		wholeSector    bool
		expectedResult int
	}{
		{"example", "example.txt", false, 36},
		{"puzzle", "puzzle.txt", false, 709}, // x < 232
		{"example part 2", "example.txt", true, 81},
		{"puzzle part 2", "puzzle.txt", true, 1326},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			result := GetTrailheadScore(tc.contentFile, tc.wholeSector)
			if result != tc.expectedResult {
				t.Errorf("Expected for %s: %d, got %d", tc.caseName, tc.expectedResult, result)
			}
		})
	}
}
