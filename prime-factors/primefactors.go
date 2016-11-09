package prime

const testVersion = 2

// Factors calculates the factors that multiply to create the prime number
func Factors(in int64) (res []int64) {
	if in < 2 {
		return []int64{}
	}
	num := int64(2)
	for {
		if in%num == 0 {
			res = append(res, num)
			in /= num
			continue
		}
		if in == 1 {
			break
		}
		num++
	}
	return
}
