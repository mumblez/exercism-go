package prime

// Nth returns nth prime number
func Nth(pn int) (int, bool) {
	if pn < 1 {
		return 0, false
	}
	var prime bool
	for num, pnsFound := 2, 0; ; num++ {
		prime = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				prime = false
				break
			}
		}
		if prime || num == 2 {
			pnsFound++
		}
		if pnsFound == pn {
			return num, true
		}
	}
}
