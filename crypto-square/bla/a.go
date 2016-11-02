package cryptosquare

import (
	"math"
	"strings"
)

const testVersion = 2

func normalize(in string) (out string) {
	for _, c := range in {
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
			out += string(c)
		}
		if c >= 'A' && c <= 'Z' {
			out += string(c + 32)
		}
	}
	return
}

// Encode ..
func Encode(plain string) string {
	norm := normalize(plain)
	normLen := len(norm)
	sideLen := int(math.Ceil(math.Sqrt(float64(normLen))))

	encodedLines := make([]string, sideLen)
	for i := 0; i < normLen; i++ {
		encodedLines[i%sideLen] += string(norm[i])
	}

	return strings.Join(encodedLines, " ")
}
