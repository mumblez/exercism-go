// Package hamming calculates mutations in DNS strands.
package hamming

import (
	"fmt"
)

const testVersion = 5

// Distance calculates the number of mutations given a pair of same length strings.
func Distance(a, b string) (diff int, err error) {
	// compare length and return error if different
	if len(a) != len(b) {
		diff = -1
		err = fmt.Errorf("Lengths don't match\n")
		return
	}

	// range over each letter and increment diff if other string does not match
	for k, v := range a {
		if byte(v) != b[k] {
			diff++
		}
	}
	return
}
