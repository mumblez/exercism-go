package binary

import "errors"

const testVersion = 2

// ParseBinary parses a string of binary and outputs it's decimal value
func ParseBinary(input string) (int, error) {
	var result int
	for i, c := range input {
		switch c {
		case '1':
			result += 1 << uint64(len(input)-i-1)
		case '0':
			continue
		default:
			return 0, errors.New("Invalid input")
		}
	}
	return result, nil
}
