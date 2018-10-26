package luhn

import (
	"strconv"
	"strings"
)

// Valid returns if a string is Luhn checksum valid
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	if len(s) <= 1 {
		return false
	}
	l := len(s)

	// for even length numbers the doubled up digit indexes are 0, 2, 4...
	// for odd length numbers the doubled up digit indexes are 1, 3, 5...
	// for example: 059 has an odd length and we double up id 1
	// 1234 has even length and we double up id 2 and 0
	lEven := true
	if l%2 == 1 {
		lEven = false
	}
	sum := 0
	for i := 0; i < l; i++ {
		value, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}
		if (lEven && i%2 == 0) || (!lEven && i%2 == 1) {
			value = 2 * value
			if value > 9 {
				value = value - 9
			}
		}
		sum += value
	}

	if sum%10 == 0 {
		return true
	}
	return false
}
