// Package pascal returns pascals triangles
package pascal

// Triangle calculates pascals triangles row by row
func Triangle(n int) (t [][]int) {
	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		for a := 0; a <= i; a++ {
			// beginning and end
			if a == 0 || a == i {
				row[a] = 1
			} else {
				// add numbers from previous row directly above (a) and above -1 (a-1)
				row[a] = t[i-1][a-1] + t[i-1][a]
			}
		}
		t = append(t, row)
	}
	return
}
