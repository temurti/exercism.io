// Package proverb generates proverb based on array with nouns
package proverb

import "fmt"

// Proverb function returns array with proverb
func Proverb(rhyme []string) []string {
	var output []string // array with proverb
	if len(rhyme) > 0 {
		for i := 0; i < len(rhyme)-1; i++ {
			output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
		output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}
	return output
}
