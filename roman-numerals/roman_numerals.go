package romannumerals

import "errors"

const testVersion = 3

// Runit holds the roman numeral for one / five (*10)
type Runit struct {
	five string
	one  string
}

// ator is a map of arabic units to roman units (one / five)
var ator = map[int]Runit{
	1:    {five: "V", one: "I"},
	10:   {five: "L", one: "X"},
	100:  {five: "D", one: "C"},
	1000: {five: "M", one: "M"},
}

func appendRN(letter string, result *string, qty int) {
	for i := 0; i < qty; i++ {
		*result += letter
	}
}

// ToRomanNumeral ...
func ToRomanNumeral(year int) (string, error) {
	if year > 3999 || year < 1 {
		return "", errors.New("Invalid year")
	}
	var result string
	for yearunit := 1000; yearunit >= 1; yearunit /= 10 {
		unit := year / yearunit
		if unit < 1 {
			continue
		}
		switch unit {
		case 1, 2, 3:
			appendRN(ator[yearunit].one, &result, unit)
		case 5, 6, 7, 8:
			appendRN(ator[yearunit].five, &result, 1)
			appendRN(ator[yearunit].one, &result, (unit - 5))
		case 4:
			appendRN(ator[yearunit].one, &result, 1)
			appendRN(ator[yearunit].five, &result, 1)
		case 9:
			appendRN(ator[yearunit].one, &result, 1)
			appendRN(ator[yearunit*10].one, &result, 1)
		}
		year -= (yearunit * unit)
	}
	return result, nil
}
