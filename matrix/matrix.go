package matrix

import (
	"errors"
	"strconv"
	"strings"
)

const testVersion = 1

// ErrInvalid ...
var ErrInvalid = errors.New("Invalid Input")

// Matrix ...
type Matrix [][]int

// New ...
func New(in string) (*Matrix, error) {
	m := new(Matrix)
	var clen int
	for _, row := range strings.Split(in, "\n") {
		if clen == 0 {
			clen = len(strings.Fields(row))
		} else if clen != len(strings.Fields(row)) {
			return nil, ErrInvalid
		}
		var r []int
		for _, col := range strings.Fields(row) {
			num, err := strconv.Atoi(col)
			if err != nil {
				return nil, ErrInvalid
			}
			r = append(r, num)
		}
		*m = append(*m, r)
	}
	return m, nil

}

// Rows ...
func (m *Matrix) Rows() Matrix {
	var r Matrix

	return r

}

// Cols ...
func (m *Matrix) Cols() Matrix {
	var c Matrix

	return c

}

// Set ...
func (m *Matrix) Set(row, col, val int) bool {
	return false
}
