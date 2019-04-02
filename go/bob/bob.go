package bob

import (
	"regexp"
)

// Hey
func Hey(remark string) string {
	reply := "Whatever."
	matched1, _ := regexp.MatchString(`^[ .,\-A-Z0-9!]*$`, remark)
	matched2, _ := regexp.MatchString(`^[ .,\-A-Z0-9!]*\?$`, remark)
	matched3, _ := regexp.MatchString(`\?$`, remark)
	matched4, _ := regexp.MatchString(`^\s*$`, remark)

	if matched1 {
		reply = "Whoa, chill out!"
	} else if matched2 {
		reply = "Calm down, I know what I'm doing!"
	} else if matched3 {
		reply = "Sure."
	} else if matched4 {
		reply = "Fine. Be that way!"
	}

	return reply
}
