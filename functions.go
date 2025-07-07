package main

import (
	"fmt"
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