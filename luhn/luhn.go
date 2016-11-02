package luhn

//package main

import (
	"strconv"
	"unicode"
)

func luhnChecksum(input string) int {
	// sum input
	var total int
	odd := true
	for i := len(input) - 1; i >= 0; i-- {
		// byte value of rune - 48 gets us it's real number
		num := int(input[i]) - 48
		if odd {
			total += num
			odd = false
		} else {
			num *= 2
			if num < 10 {
				total += num
			} else {
				total += (num - 9)
			}
			odd = true
		}
	}
	return total % 10
}

func cleanString(input string) string {
	var cleaninput string
	for _, c := range input {
		if unicode.IsNumber(c) {
			cleaninput += string(c)
		}
	}
	return cleaninput
}

// AddCheck adds the checkum digit to the string
func AddCheck(input string) string {
	cleaninput := cleanString(input)
	// *10 or add "0"
	total := luhnChecksum(cleaninput + "0")
	checkDigit := total % 10
	if checkDigit != 0 {
		checkDigit = 10 - checkDigit
	}
	return input + strconv.Itoa(checkDigit)
}

// Valid uses Luhn algo to confirm string is valid
func Valid(input string) bool {
	// clean string
	cleaninput := cleanString(input)
	//	fmt.Printf("cleaninput = %+v\n", cleaninput)
	if len(cleaninput) == 0 {
		return false
	}

	total := luhnChecksum(cleaninput)

	// check if mod 10 == 0
	if total == 0 {
		return true
	}
	return false
}
