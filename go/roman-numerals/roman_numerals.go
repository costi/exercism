package romannumerals

import (
	"bytes"
	"errors"
	"strings"
)

type romanNumeral struct {
	letter string
	value  int
}

// ToRomanNumeral takes an arabic number as int and returns the roman numeral as string
// I, V, X, L, C, D, M.
func ToRomanNumeral(arabic int) (string, error) {
	if arabic < 0 {
		return "", errors.New("negative numbers do not exist in the Roman Numerals system")
	}
	if arabic == 0 {
		return "", errors.New("zero does not exist in the Roman Numerals system")
	}
	if arabic > 3000 {
		return "", errors.New("numbers over 3000 are impossible")
	}
	var numerals = []romanNumeral{
		romanNumeral{letter: "M", value: 1000},
		romanNumeral{letter: "CM", value: 900},
		romanNumeral{letter: "D", value: 500},
		romanNumeral{letter: "CD", value: 400},
		romanNumeral{letter: "C", value: 100},
		romanNumeral{letter: "XC", value: 90},
		romanNumeral{letter: "L", value: 50},
		romanNumeral{letter: "XL", value: 40},
		romanNumeral{letter: "X", value: 10},
		romanNumeral{letter: "IX", value: 9},
		romanNumeral{letter: "V", value: 5},
		romanNumeral{letter: "IV", value: 4},
		romanNumeral{letter: "I", value: 1},
	}
	var buffer bytes.Buffer
	var quotient, remainder, current int
	current = arabic
	for _, romanLetter := range numerals {
		quotient = current / romanLetter.value
		remainder = current % romanLetter.value
		if quotient <= 3 && quotient >= 1 {
			buffer.WriteString(strings.Repeat(romanLetter.letter, quotient))
		}

		current = remainder
	}

	return buffer.String(), nil
}
