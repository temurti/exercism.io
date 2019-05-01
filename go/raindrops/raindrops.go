//Package raindrops - Convert a number to a string, the contents of which depend on the number's factors.
package raindrops

import "fmt"

// Convert function converts a number to a string, the contents of which depend on the number's factors
func Convert(number int) string {
	var output string
	pfs := Factor(number)
	for _, a := range pfs {
		if a == 3 {
			output += "Pling"
		} else if a == 5 {
			output += "Plang"
		} else if a == 7 {
			output += "Plong"
		}
	}
	if output == "" {
		return fmt.Sprintf("%v", number)
	}
	return output
}

// Factor function - Get all prime factors of a given number
func Factor(number int) []int {
	pfs := []int{1}
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			pfs = append(pfs, i)
		}
	}
	pfs = append(pfs, number)
	return pfs
}
