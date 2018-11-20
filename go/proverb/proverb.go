// Package proverb generates proverbs
package proverb

import (
	"fmt"
)

const (
	stanza = "For want of a %s the %s was lost."
	last   = "And all for the want of a %s."
)

// Proverb takes a list of words
// returns a list of proverbs like "For want of a <word0> the <word1> was lost."
func Proverb(rhyme []string) []string {
	length := len(rhyme)

	if length == 0 {
		return []string{}
	}

	proverb := make([]string, length)

	// we go to length-1 because we add rhyme[i+1] to the string on each loop
	for i := 0; i < length-1; i++ {
		proverb[i] = fmt.Sprintf(stanza, rhyme[i], rhyme[i+1])
	}

	proverb[length-1] = fmt.Sprintf(last, rhyme[0])

	return proverb
}
