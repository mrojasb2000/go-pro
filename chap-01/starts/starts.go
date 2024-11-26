package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Generator seed ramdom numbers
	rand.NewSource(time.Now().UnixNano())
	// Generate a random number between 0 and then add 1 to get a number between 1 and 5
	r := rand.Intn(5) + 1
	// Use the repeat to create a string with the number of starts we need
	starts := strings.Repeat("*", r)
	fmt.Println(starts)
}
