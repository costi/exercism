//package pangram finds if sentences are pangrams
// it doesn't have any external dependencies like unicode
// because it deals only with ascii
package pangram

// IsPangram returns true when sentence is a pangram
// meaning that it has all the words in the alphabet
// case insensitive
func IsPangram(s string) bool {
	var appear [26]bool
	for _, c := range s {
		lowercase := (c >= 'a' && c <= 'z')
		uppercase := (c >= 'A' && c <= 'Z')
		if lowercase || uppercase {
			if uppercase {
				c = c + 32
			}
			if !appear['z'-c] {
				appear['z'-c] = true
			}
		}
	}
	for _, v := range appear {
		if !v {
			return false
		}
	}
	return true
}
