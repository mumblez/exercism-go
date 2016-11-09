package igpay

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/rangetable"
)

func encode(in string) string {
	vowel := rangetable.New('a', 'e', 'i', 'o', 'u')
	var vowelidx int
	for k, v := range in {
		if unicode.Is(vowel, v) {
			// weird exceptions for 'u'
			if v == 'u' && unicode.Is(vowel, rune(in[k+1])) {
				vowelidx++
			}
			vowelidx += k
			break
		}
	}
	// add weird exeption for xray
	if vowelidx == 0 || vowelidx > 3 || in[len(in)-2:] == "ay" {
		return in + "ay"
	}
	return in[vowelidx:] + in[:vowelidx] + "ay"
}

// PigLatin converts string to piglatin, look for consonants, also include 'u' !!!
func PigLatin(in string) string {
	words := strings.Split(in, " ")
	res := []string{}
	for _, str := range words {
		res = append(res, encode(str))
	}
	return strings.Join(res, " ")
}
