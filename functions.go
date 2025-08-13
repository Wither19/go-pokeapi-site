package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"strconv"
)

// Adds [length] amount of leading zeroes to [num], for display purposes.
func leadingZeroes(num int, length int) string {
	return fmt.Sprintf("%0*d", length, num)
}

// Random Integer function for use in Go templates, cuz Sprig's randInt suddenly borked.
func randomNumber(start int, end int) int {
	return rand.IntN(end) + start
}

func parseInt(string string) int {
	num, err := strconv.ParseInt(string, 0, 0)
	if err != nil {
		log.Fatalln("Could not convert number:", err)
	}

	return int(num)
}
