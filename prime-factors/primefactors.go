package prime

const testVersion = 2

// Factors calculates the factors that multiply to create the prime number
func Factors(in int64) []int64 {
	res := []int64{}
	for i := int64(2); in > 1; {
		if in%i == 0 {
			res = append(res, i)
			in /= i
			continue
		}
		i++
	}
	return res
}
