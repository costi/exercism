// Package hamming calculates the hamming distance in DNA
package hamming

import "errors"

// Distance returns the hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("The strings are of different size")
	}
	length := len(a)
	var distance int
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
