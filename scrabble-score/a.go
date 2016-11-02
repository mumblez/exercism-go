package scrabble

import "strings"

const (
	p1          = "AEIOULNRST"
	p2          = "DG"
	p3          = "BCMP"
	p4          = "FHVWY"
	p5          = "K"
	p8          = "JX"
	p10         = "QZ"
	testVersion = 4
)

// Score calculates scrabble points
func Score(input string) int {
	input = strings.ToUpper(input)
	var points int
	for _, v := range input {
		switch {
		case strings.Contains(p1, string(v)):
			points++
		case strings.Contains(p2, string(v)):
			points += 2
		case strings.Contains(p3, string(v)):
			points += 3
		case strings.Contains(p4, string(v)):
			points += 4
		case strings.Contains(p5, string(v)):
			points += 5
		case strings.Contains(p8, string(v)):
			points += 8
		case strings.Contains(p10, string(v)):
			points += 10
		}
	}
	return points
}
