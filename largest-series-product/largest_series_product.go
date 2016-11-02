package lsproduct

import "fmt"

const testVersion = 4

// LargestSeriesProduct returns largest sets multiplied
func LargestSeriesProduct(series string, span int) (int, error) {
	if span > len(series) || span < 0 {
		return -1, fmt.Errorf("Invalid span")
	}
	if span == 0 {
		return 1, nil
	}
	highestTotal := 0
	for n := 0; n <= (len(series) - span); n++ {
		counter := 0
		for i := 0; i < span; i++ {
			if series[i+n] < '0' || series[i+n] > '9' {
				return -1, fmt.Errorf("Invalid character found")
			}
			num := int(series[i+n] - '0')
			if i == 0 {
				counter += num
			} else {
				counter *= num
			}
		}
		if counter > highestTotal {
			highestTotal = counter
		}
	}
	return highestTotal, nil
}
