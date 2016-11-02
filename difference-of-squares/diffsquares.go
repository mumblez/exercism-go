package diffsquares

// SquareOfSums returns square of sums
func SquareOfSums(n int) int {
	tally := 0
	for i := n; i != 0; i-- {
		tally += i
	}
	return tally * tally
}

// SumOfSquares returns sum of squares
func SumOfSquares(n int) int {
	tally := 0
	for i := n; i != 0; i-- {
		tally += i * i
	}
	return tally
}

// Difference returns the difference of SquareOfSums - SumOfSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
