// Package leap informs if the year is a leap year or not.
package leap

// IsLeapYear receives an int indicating a year and returns a boolean that tells if this is a leap year or not
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	return true
}
