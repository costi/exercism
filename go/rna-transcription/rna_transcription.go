// Package strand deals with DNA and RNA strands
package strand

import "bytes"

// ToRNA takes in a DNA string and returns its RNA
func ToRNA(dna string) string {
	d2r := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
	var rna bytes.Buffer
	for _, letter := range dna {
		rna.WriteRune(d2r[letter])
	}

	return rna.String()
}
