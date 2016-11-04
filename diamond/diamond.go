package diamond

import (
	"errors"
	"strings"
)

const testVersion = 1

// Gen creates our diamond
func Gen(c byte) (result string, err error) {
	if c < 'A' || c > 'Z' {
		err = errors.New("Invalid character")
	}
	rlen := int(((c - 'A' + 1) * 2) - 1)
	block := make([]string, rlen)

	// add character in the correct positions
	// once we cross halfway point the formula for letter and row change
	for k := range block {
		letter := 'A' + ((rlen / 2) - (k - (rlen / 2)))
		if k < (rlen+1)/2 {
			letter = 'A' + k%((rlen+1)/2)
		}
		var row string
		for i := 0; i < rlen; i++ {
			rowpos := (rlen - 1) - i
			if i < (rlen+1)/2 {
				rowpos = i % ((rlen + 1) / 2)
			}
			letterpos := (rlen / 2) - (letter - 'A')
			if rowpos == letterpos {
				row += string(letter)
				continue
			}
			row += "-"
		}
		block[k] = row
	}

	result = strings.Join(block, "\n")
	return
}
