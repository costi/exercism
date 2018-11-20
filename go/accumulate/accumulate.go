// Package accumulate maps arrays
package accumulate

// Accumulate is actually map. It accepts a array of strings and runs the converter function
// on each of the strings. Returns the converted array.
func Accumulate(wordList []string, converter func(string) string) []string {
	var convertedList = make([]string, len(wordList))
	for i, word := range wordList {
		convertedList[i] = converter(word)
	}
	return convertedList
}
