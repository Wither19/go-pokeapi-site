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

func parseInt(string string) int {
	num, err := strconv.ParseInt(string, 0, 0)
	if err != nil {
		log.Fatalln("Could not convert number:", err)
	}

	return int(num)
}
