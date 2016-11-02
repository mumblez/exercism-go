package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
)

type Unit map[uint64]string

var num = Unit{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var teen = Unit{
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var tens = Unit{
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

var big = Unit{
	1e3:  "thousand",
	1e6:  "million",
	1e9:  "billion",
	1e12: "trillion",
	1e15: "quadrillion",
	1e18: "quintillion",
}

func Macsay(s string) {
	saybin, err := exec.LookPath("say")
	if err != nil {
		fmt.Printf("err = %+v, say binary could not be found!\n", err)
		os.Exit(1)
	}
	err = exec.Command(saybin, s).Run()
	if err != nil {
		fmt.Printf("err = %+v\n", err)
	}
}

// Say input number as english words
func say(num uint64) string {
	var result string

	return result

}

func main() {
	fmt.Println(big[100])
	fmt.Println(big[1e2])
	fmt.Println(big[1000])
	fmt.Printf("100000 exponent = %v\n", math.Ilogb(3000))
	fmt.Printf("123400 exponent = %v\n", math.Exp(5))
}
