package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// NameInUse a map of names in use
var NameInUse = map[string]bool{}

// Robot type with just name field
type Robot struct {
	name string
}

// Name returns it's current name
func (r *Robot) Name() string {
	if r.name != "" {
		return r.name
	}
	r.name = genName()
	return r.name
}

// Reset generates a new name
func (r *Robot) Reset() {
	r.name = genName()
}

// genName generates a new random name
func genName() string {
	// create random seed for our random number generator
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	var tempName string

	// add 65 as that's the start of the byte for 'A' (25 including 0 is the alphabet size)
	for {
		letter1 := rng.Intn(25) + 65
		letter2 := rng.Intn(25) + 65
		digits := rng.Intn(999)
		tempName = string(letter1) + string(letter2) + fmt.Sprintf("%03d", digits)

		if !NameInUse[tempName] {
			NameInUse[tempName] = true
			break
		}
	}
	return tempName
}
