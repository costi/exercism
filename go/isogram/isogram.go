package isogram

import (
	"unicode"
)

// IsIsogram returns true if a word has only non-repeating characters
func IsIsogram(word string) bool {
	counts := make(map[rune]int)
	for _, c := range word {
		if (c != ' ') && (c != '-') {
			c = unicode.ToLower(c)
			_, ok := counts[c]
			if ok {
				counts[c]++
				if counts[c] > 1 {
					return false
				}
			} else {
				counts[c] = 1
			}
		}
	}
	return true
}
