package ocr

import "strings"

var numbers = map[string]string{
	" _ | ||_|": "0",
	"     |  |": "1",
	" _  _||_ ": "2",
	" _  _| _|": "3",
	"   |_|  |": "4",
	" _ |_  _|": "5",
	" _ |_ |_|": "6",
	" _   |  |": "7",
	" _ |_||_|": "8",
	" _ |_| _|": "9",
}

func recognizeDigit(s string) string {
	for k, v := range numbers {
		if s == k {
			return v
		}
	}
	return "?"
}

// Recognize ...
func Recognize(input string) []string {
	var digit, ret []string
	var lines [][]string
	for k, v := range strings.Split(input, "\n") {
		if k%4 == 0 {
			continue
		}
		for x, y := range v {
			if k%4 == 1 && x%3 == 0 {
				digit = append(digit, string(y))
				continue
			}
			digit[x/3] += string(y)
		}
		if k%4 == 3 {
			lines = append(lines, digit)
			digit = []string{}
		}
	}
	for k, line := range lines {
		for kk, v := range line {
			if kk == 0 {
				ret = append(ret, recognizeDigit(v))
				continue
			}
			ret[k] += recognizeDigit(v)
		}
	}
	return ret
}
