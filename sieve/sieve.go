package sieve

import "sort"

// Sieve returns primes within submitted limit
func Sieve(limit int) []int {
	// create and fill map
	nums := make(map[int]bool, limit)
	for i := 2; i <= limit; i++ {
		if i%2 == 0 && i != 2 {
			nums[i] = false
			continue
		}
		nums[i] = true
	}
	// mark false multiples
	for i := 3; i <= (limit / 2); i++ {
		for n := 2; n <= (limit / i); n++ {
			nums[n*i] = false
		}
	}
	// gather our precious primes
	var primes []int
	for num, valid := range nums {
		if valid {
			primes = append(primes, num)
		}
	}
	sort.Ints(primes)
	return primes
}
