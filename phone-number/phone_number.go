package phonenumber

import "errors"

const testVersion = 1

// Number checks if phone number is valid
func Number(s string) (res string, err error) {
	// clean input
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			res += string(s[i])
		}
	}
	if len(res) == 10 {
		return
	}
	if len(res) < 10 || len(res) > 11 || res[0] != '1' {
		err = errors.New("invalid number of digits")
	}
	res = res[1:]
	return
}

// AreaCode ...
func AreaCode(s string) (res string, err error) {
	cleanInput, err := Number(s)
	if err != nil {
		return
	}
	res += cleanInput[:3]
	return
}

// Format ...
func Format(s string) (res string, err error) {
	cleanInput, err := Number(s)
	if err != nil {
		return
	}
	area, err := AreaCode(cleanInput)
	if err != nil {
		return
	}
	res += "("
	res += area
	res += ") "
	res += cleanInput[3:6]
	res += "-"
	res += cleanInput[6:]
	return
}
