// Package leap calculates if the given year is a leap year.
package leap

const testVersion = 2

// IsLeapYear returns true if year = leap year else false.
func IsLeapYear(year int) bool {
	if year%4 > 0 {
		return false
	}
	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}
		return false
	}
	return true
}
