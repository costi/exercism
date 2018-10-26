// Package reverse reverses stuff, like strings
package reverse

// String reverses a string. Returns a string
func String(s string) string {
	// cast s to an array of runes
	r := []rune(s)
	// reverse the array in place
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		// swap in place because go is cool
		r[i], r[j] = r[j], r[i]
	}
	// cast to string
	return string(r)
}
