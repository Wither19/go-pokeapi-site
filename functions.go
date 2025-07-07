package main

import (
	"fmt"
	"math/rand/v2"
	"unicode"
)

func leadingZeroes(num int, length int) string {
	return fmt.Sprintf("%0*d", length, num)
}

func caps(s string) string {
	if len(s) == 0 {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func randomNumber(s int, e int) int {
	return rand.IntN(e) + s
}