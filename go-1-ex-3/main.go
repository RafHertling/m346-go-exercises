package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Seed the random number generator with current time
	rand.Seed(time.Now().UnixNano())

	// Generate random number between 1 and 6
	eyes := rand.Intn(6) + 1

	// Get current time
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Output dice roll to stdout (normal program output)
	fmt.Fprintln(os.Stdout, eyes)

	// Output timestamp to stderr (logging information)
	fmt.Fprintln(os.Stderr, "Rolled at:", currentTime)
}

// To redirect the outputs to different files, run the program as follows:
// ./main > eyes.txt 2> dice.log
