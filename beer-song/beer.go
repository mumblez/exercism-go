package beer

import (
	"errors"
	"fmt"
	"strings"
)

const testVersion = 1

func bottle(v int) string {
	switch v {
	case 1:
		return "1 bottle"
	case 0:
		return "no more bottles"
	case -1:
		return "99 bottles"
	default:
		return fmt.Sprintf("%d bottles", v)
	}
}

func isValidVerse(v int) bool {
	if v < 0 || v > 99 {
		return false
	}
	return true
}

// Verse makes rainbows come out your mouth
func Verse(v int) (string, error) {
	if !isValidVerse(v) {
		return "", errors.New("invalid")
	}
	var l1, l2, take, takeitdown string
	ofbeer := "of beer"
	onthewall := "on the wall"
	capitaliseN := func(r rune) rune {
		if r == 'n' {
			return 'N'
		}
		return r
	}
	if v == 1 {
		take = "it"
	} else {
		take = "one"
	}
	if v == 0 {
		takeitdown = "Go to the store and buy some more"
	} else {
		takeitdown = "Take " + take + " down and pass it around"
	}
	l1 = fmt.Sprintf("%s %s %s, %s %[2]s.\n", strings.Map(capitaliseN, bottle(v)), ofbeer, onthewall, bottle(v))
	l2 = fmt.Sprintf("%s, %s %s %s.\n", takeitdown, bottle(v-1), ofbeer, onthewall)
	return l1 + l2, nil
}

// Verses ...
func Verses(start, stop int) (string, error) {
	if !isValidVerse(start) || !isValidVerse(stop) || start < stop {
		return "", errors.New("invalid")
	}
	var res string
	for i := start; i >= stop; i-- {
		v, _ := Verse(i)
		res += v
		res += "\n"
	}
	return res, nil
}

// Song ...
func Song() string {
	res, _ := Verses(99, 0)
	return res
}
