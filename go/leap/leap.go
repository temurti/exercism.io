// Package leap verifies if given a year is a leap year
package leap

// IsLeapYear verifies is that a leap year
func IsLeapYear(year int) bool {
	var isLeapYear bool
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		isLeapYear = true
	}
	return isLeapYear
}
