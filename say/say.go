package say

// Unit to map a number with it's string
type Unit map[uint64]string

var numreal = Unit{
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
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

var big = []string{"thousand", "million", "billion", "trillion", "quadrillion", "quintillion"}

func section(num uint64) (res string) {
	unit := num % 10
	ten := (num / 10) % 10
	hundred := (num / 100) % 10
	if hundred > 0 {
		res += numreal[hundred] + " hundred"
	}
	if unit == 0 && ten == 0 {
		return
	}
	if (unit > 0 || ten > 0) && hundred > 0 {
		res += " "
	}
	if ten == 1 {
		res += teen[(ten*10)+unit]
		return
	}
	res += tens[ten]
	if unit == 0 {
		return
	}
	if ten > 1 && unit > 0 {
		res += "-"
	}
	res += numreal[unit]
	return
}

// Say input number as english words
func Say(num uint64) string {
	var result string
	if num < 1 {
		return "zero"
	}
	var numsections []uint64

	// break numbers up into thousands (3 digits per element)
	for {
		if num <= 0 {
			break
		}
		numsections = append(numsections, num%uint64(1000))
		num /= 1000
	}

	// say 3 digit number (if last set, combine ten and unit number with dash, e.g twenty-three)
	for i := len(numsections) - 1; i >= 0; i-- {
		result += section(numsections[i])
		if i == 0 {
			break
		}
		if numsections[i] > 0 {
			result += " " + big[i-1]
		}
		if numsections[i-1] > 0 {
			result += " "
		}
	}
	return result
}

// func Macsay(num uint64) {
//  sentence := Say(num)
// 	saybin, err := exec.LookPath("say")
// 	if err != nil {
// 		fmt.Printf("err = %+v, say binary could not be found!\n", err)
// 		os.Exit(1)
// 	}
// 	err = exec.Command(saybin, sentence).Run()
// 	if err != nil {
// 		fmt.Printf("err = %+v\n", err)
// 	}
// }

// func main() {
// 	fmt.Println(say(1234567890))
// }
