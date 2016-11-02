// Package bob calculates bob's responses
package bob

import "strings"
import "unicode"

const testVersion = 2 // same as targetTestVersion

// Hey returns bob's response
func Hey(s string) (resp string) {
	s = strings.TrimSpace(s)

	shouting := true
	nothing := true
	noletters := true
	var question bool

	// Question if ends in ?
	if strings.HasSuffix(s, "?") {
		question = true
	}

	for _, v := range s {
		if unicode.IsLetter(v) {
			nothing = false
			noletters = false
		}
		if unicode.IsLower(v) {
			shouting = false
		}
		if unicode.IsNumber(v) {
			nothing = false
		}
	}

	switch {
	case shouting && !noletters:
		resp = "Whoa, chill out!"
	case question:
		resp = "Sure."
	case nothing:
		resp = "Fine. Be that way!"
	default:
		resp = "Whatever."
	}

	return

}
