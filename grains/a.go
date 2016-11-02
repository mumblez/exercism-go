package grains

import "errors"

const testVersion = 1

// Square returns 2 to the power of input
func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("Invalid range")
	}
	return 1 << uint64(input-1), nil
}

// Total returns the maximum 64 bit integer value
func Total() uint64 {
	return 1<<64 - 1
}
