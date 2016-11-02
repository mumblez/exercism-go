package summultiples

// MultipleSummer returns a faction that does the work
func MultipleSummer(multi ...int) func(int) int {
	MultiFunc := func(limit int) (sum int) {
		if len(multi) == 0 {
			return 0
		}
		// for every number from 1 up to the limit
		// check if it is a multiple of elements in the multi slice
		// if it's a multiple (mod = 0 / division leaves no remainders)
		// we accumalate in sum
		for i := 1; i < limit; i++ {
			for _, v := range multi {
				if i%v == 0 {
					sum += i
					break
				}
			}

		}
		return
	}
	return MultiFunc
}
