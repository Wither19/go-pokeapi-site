package main

import (
	"fmt"
	"math/rand/v2"
)

// Adds [length] amount of leading zeroes to [num], for display purposes.
func leadingZeroes(num int, length int) string {
	return fmt.Sprintf("%0*d", length, num)
}

// Random Integer function for use in Go templates, cuz Sprig's randInt suddenly borked.
func randomNumber(start int, end int) int {
	return rand.IntN(end) + start
}