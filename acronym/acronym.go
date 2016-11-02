// Package acronym produces acronyms from strings
package acronym

import "unicode"

const testVersion = 1

// abbreviate produces an acronym
func abbreviate(s string) (res string) {
	skip := false
	for k, v := range s {
		if skip {
			skip = false
			continue
		}
		if k == 0 {
			res += string(v)
			continue
		}
		if string(v) == ":" {
			return
		}
		if string(v) == "-" || unicode.IsSpace(v) {
			nextLetter := rune(s[k+1])
			nextLetter = unicode.ToUpper(nextLetter)
			res += string(nextLetter)
			skip = true
			continue
		}
		if unicode.IsUpper(v) {
			res += string(v)
		}
	}

	return
}
