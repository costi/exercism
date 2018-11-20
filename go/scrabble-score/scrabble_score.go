// Package scrabble calculates scrabble scores
package scrabble

// Score returns the scrabble score (int) of a string
func Score(word string) int {
	var score = 0

	for _, letter := range word {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't', 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score += 1
		case 'd', 'g', 'D', 'G':
			score += 2
		case 'b', 'c', 'm', 'p', 'B', 'C', 'M', 'P':
			score += 3
		case 'f', 'h', 'v', 'w', 'y', 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'k', 'K':
			score += 5
		case 'j', 'x', 'J', 'X':
			score += 8
		case 'q', 'z', 'Q', 'Z':
			score += 10
		}
	}
	return score
}
