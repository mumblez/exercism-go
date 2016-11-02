package wordcount

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Frequency type
type Frequency map[string]int

// WordCount counts frequence of words in phrase
func WordCount(phrase string) Frequency {
	var cleanPhrase string
	for _, v := range phrase {
		if unicode.IsSpace(v) || unicode.IsLetter(v) || unicode.IsNumber(v) {
			cleanPhrase += string(unicode.ToLower(v))
		}
	}
	fm := Frequency{}
	for _, v := range strings.Fields(cleanPhrase) {
		fm[v]++
	}
	return fm
}
