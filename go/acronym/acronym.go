// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym creates acronyms from phrases
package acronym

import "strings"

// Abbreviate creates receives a string and return its acronym
// Ex: NSA from National Security Agency
func Abbreviate(s string) string {
	var foundSpace = true
	initials := func(r rune) rune {
		switch {
		case r == ' ' || r == '-':
			foundSpace = true
			return -1
		case foundSpace:
			foundSpace = false
			return r
		default:
			return -1
		}
	}

	return strings.ToUpper(strings.Map(initials, s))
}
