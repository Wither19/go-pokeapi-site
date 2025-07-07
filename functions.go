package main

import (
	"fmt"
	"math/rand/v2"
	"unicode"
)

// Adds [length] amount of leading zeroes to [num], for display purposes.
func leadingZeroes(num int, length int) string {
	return fmt.Sprintf("%0*d", length, num)
}

// Capitalizes first unicode character in string [s]
func caps(s string) string {
	if len(s) == 0 {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// Random Integer function for use in Go templates, cuz Sprig's randInt suddenly borked.
func randomNumber(start int, end int) int {
	return rand.IntN(end) + start
}