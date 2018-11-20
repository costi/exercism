// Package raindrops makes raindrop-speak from  numbers
package raindrops

import (
	"bytes"
	"strconv"
)

const (
	threeSpeak = "Pling"
	fiveSpeak  = "Plang"
	sevenSpeak = "Plong"
)

// Convert takes in a number and returns a string in raindrop-speak
func Convert(number int) string {
	var buffer bytes.Buffer
	if number%3 == 0 {
		buffer.WriteString(threeSpeak)
	}

	if number%5 == 0 {
		buffer.WriteString(fiveSpeak)
	}

	if number%7 == 0 {
		buffer.WriteString(sevenSpeak)
	}

	if buffer.Len() == 0 {
		return strconv.Itoa(number)
	}

	return buffer.String()
}
