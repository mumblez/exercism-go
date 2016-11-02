package palindrome

import "fmt"

type Product struct {
	Product        int
	Factorizations [][2]int
}

const testVersion = 1

// isPalindrome checks each digit by multiplying, dividing and modding by 10 each iteration
func isPalindrome(x int) bool {
	div := 1
	for x/div >= 10 {
		div *= 10
	}
	for x != 0 {
		l := x / div
		r := x % 10
		if l != r {
			return false
		}
		x = (x % div) / 10
		div /= 100
	}
	return true
}

// isPairExist helper function checks if pair already exists (both orders redundant)
func (p *Product) isPairExist(x, y int) bool {
	for _, v := range p.Factorizations {
		if (x == v[0] && y == v[1]) || (x == v[1] && y == v[0]) {
			return true
		}
	}
	return false
}

// appendPair helper function checks if pair exists and if not then append
func (p *Product) appendPair(x, y int) {
	if p.isPairExist(x, y) {
		return
	}
	p.Factorizations = append(p.Factorizations, [2]int{x, y})
}

// copyProduct helper function rebuilds the Factorization [][2]int as a copy will not suffice
func (p *Product) copyProduct(pmax *Product) {
	p.Product = pmax.Product
	p.Factorizations = p.Factorizations[:0]
	for _, v := range pmax.Factorizations {
		p.Factorizations = append(p.Factorizations, v)
	}
}

// Products returns smallest and largest palindrome found
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = fmt.Errorf("fmin > fmax")
		return
	}
	for x := fmin; x <= fmax; x++ {
		for y := fmin; y <= fmax; y++ {
			res := x * y
			if isPalindrome(res) {
				switch {
				case pmax.Product == 0:
					pmax.Product = res
					pmax.appendPair(x, y)
				case res > pmax.Product:
					if pmin.Product == 0 {
						pmin.copyProduct(&pmax)
					}
					pmax.Product = res
					pmax.Factorizations = pmax.Factorizations[:0]
					pmax.appendPair(x, y)
				case res == pmin.Product:
					pmin.appendPair(x, y)
				case res == pmax.Product:
					pmax.appendPair(x, y)
				}
			}
		}
	}
	if pmin.Product == 0 && pmax.Product == 0 {
		err = fmt.Errorf("No palindromes")
		return
	}
	return
}
