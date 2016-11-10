package octal

import "errors"

const testVersion = 1

// ParseOctal ...
func ParseOctal(in string) (int64, error) {
	var res, octal int64
	octal = 1
	for i := len(in) - 1; i >= 0; i-- {
		if in[i] < '0' || in[i] > '7' {
			return 0, errors.New("invalid")
		}
		res += int64(in[i]-'0') * octal
		octal *= 8
	}
	return res, nil
}
