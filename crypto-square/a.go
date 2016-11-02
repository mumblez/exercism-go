package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const (
	testVersion = 2
)

// Assemble ...
func assemble(r, c int, s string) string {
	// assemble new string
	var finalString string
	var startingpos int
	var charpos int
	for i := 0; i < r; i++ {
		for l := 0; l < c; l++ {
			charpos = (l * r) + startingpos
			if charpos > (len(s) - 1) {
				break
			}
			finalString += string(s[charpos])
		}
		if i != (r - 1) {
			finalString += " "
		}
		startingpos++
	}
	return finalString
}

// Encode super secret message
func Encode(s string) string {
	// if input blank or less then 2 charaters
	if len(s) < 2 {
		return s
	}
	// lowercase and extract only letters and numbers
	s = strings.ToLower(s)
	var newString string
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			newString += string(v)
		}
	}

	// calculate best rectangle size
	sq := math.Sqrt(float64(len(newString)))
	r := int(math.Ceil(sq))
	c := int(math.Floor(sq))
	if c == 1 || sq-math.Floor(sq) > .5 {
		c++
	}

	return assemble(r, c, newString)
}
