// Package queenattack calculates if queen's can attack
package queenattack

import (
	"fmt"
	"math"
	"strconv"
)

const testVersion = 1

func extractXY(s string) (x, y int, err error) {
	ltn := make(map[string]int, 8)
	ltn["a"] = 1
	ltn["b"] = 2
	ltn["c"] = 3
	ltn["d"] = 4
	ltn["e"] = 5
	ltn["f"] = 6
	ltn["g"] = 7
	ltn["h"] = 8

	x, ok := ltn[string(s[0])]
	if !ok {
		err = fmt.Errorf("Not a valid file / letter")
		return
	}
	y, err = strconv.Atoi(string(s[1]))
	if err != nil {
		err = fmt.Errorf("Not a valid number")
		return
	}
	if x < 0 || x > 8 || y < 0 || y > 8 {
		err = fmt.Errorf("Number out of range")
		return
	}
	return
}

// CanQueenAttack - can attack? valid?
func CanQueenAttack(player1, player2 string) (bool, error) {
	// validity checks
	if len(player1) != 2 || len(player2) != 2 {
		return false, fmt.Errorf("Not correct lengths")
	}
	if player1 == player2 {
		return false, fmt.Errorf("Same Square")
	}
	// convert letter to number
	p1x, p1y, err := extractXY(player1)
	if err != nil {
		return false, err
	}
	p2x, p2y, err := extractXY(player2)
	if err != nil {
		return false, err
	}
	// same row or column
	if p1x == p2x || p1y == p2y {
		return true, nil
	}
	// diagonal
	diffx := math.Abs(float64(p1x) - float64(p2x))
	diffy := math.Abs(float64(p1y) - float64(p2y))
	if diffx == diffy {
		return true, nil
	}
	return false, nil
}
