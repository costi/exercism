// Package leap deals with leap years
package leap

// IsLeapYear returns true if a year is a leap year, and false if it's not
func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
