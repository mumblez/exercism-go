// Package raindrops looks for factors of 3,5,7 and outputs it's corresponding word
package raindrops

import "fmt"

const (
	Pling       = 3
	Plang       = 5
	Plong       = 7
	testVersion = 2
)

// Convert looks for factors in the number
func Convert(a int) (res string) {

	var num int
	if num = a % Pling; num == 0 {
		res += "Pling"
	}
	if num = a % Plang; num == 0 {
		res += "Plang"
	}
	if num = a % Plong; num == 0 {
		res += "Plong"
	}

	if res == "" {
		res = fmt.Sprintf("%d", a)
	}
	return
}
