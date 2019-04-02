package strand

import "strings"

// ToRNA - converts DNA to RNA
func ToRNA(dna string) string {
	return strings.NewReplacer(
		"G", "C",
		"C", "G",
		"T", "A",
		"A", "U",
	).Replace(dna)
}
