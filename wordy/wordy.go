package wordy

import (
	"bufio"
	"strconv"
	"strings"
)

// Answer to life...
func Answer(q string) (int, bool) {
	c := bufio.NewScanner(strings.NewReader(q))
	c.Split(bufio.ScanWords)

	var res int
	var action string

	for c.Scan() {
		word := c.Text()
		if strings.HasSuffix(word, "?") {
			word = strings.TrimSuffix(word, "?")
			_, err := strconv.Atoi(word)
			if err != nil {
				return 0, false
			}
		}
		num, err := strconv.Atoi(word)
		if err != nil && word != "by" {
			action = word
		} else {
			if res == 0 {
				res += num
				continue
			}
			if word == "by" {
				continue
			}
			if res != 0 && action != "" {
				switch action {
				case "plus":
					res += num
				case "minus":
					res -= num
				case "multiplied":
					res *= num
				case "divided":
					res /= num
				default:
					return 0, false
				}
			}
		}
	}

	return res, true
}
