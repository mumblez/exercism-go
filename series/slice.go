// Package slice returns strings of numbers
package slice

// All returns all substrings of s with length n
func All(n int, s string) []string {
	res := []string{}
	last := len(s) - n
	for i := 0; i <= last; i++ {
		res = append(res, s[i:n+i])
	}
	return res
}

// UnsafeFirst returns the first substring of s with length n
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First will return error if entered length is longer than string length
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		ok = false
		return
	}

	ok = true
	first = UnsafeFirst(n, s)
	return
}
