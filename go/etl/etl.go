package etl

import "strings"

// Transform returns a mapping from letters to points
// Example: a:1, b:2, c:1...
func Transform(pointsToLetters map[int][]string) map[string]int {
	letterToPoints := make(map[string]int)
	for points, letters := range pointsToLetters {
		for _, letter := range letters {
			letterToPoints[strings.ToLower(letter)] = points
		}
	}
	return letterToPoints
}
