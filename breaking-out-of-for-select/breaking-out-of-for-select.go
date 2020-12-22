package main

import (
	"fmt"
	"time"
)

// In Go, select works similar to switch statement only difference
// is case statements should have a channel.
// The break statement in the select moves out of select and not for loop.
// If you want to move out of for loop, one option is return
// but this may not be a solution every time another alternative is to use label in break.
func main() {
	trap()
	solution()
}

func trap() {
	for i := 0; i < 3; i++ {
		select {
		case <-time.After(time.Nanosecond):
			fmt.Printf(" Breaking out of %v iteration;", i)
			break // Breaking out of select statement and not for loop
		}
	}

	fmt.Println(" Moved out of loop !!")
	// Prints:  Breaking out of 0 iteration; Breaking out of 1 iteration; Breaking out of 2 iteration; Moved out of loop !!
}

func solution() {
loop: // Creating a label
	for i := 0; i < 5; i++ {
		select {
		case <-time.After(time.Nanosecond):
			fmt.Printf(" Breaking out of %v iteration;", i)
			break loop // Breaking out on label
		}
	}
	fmt.Println(" Moved out of loop !!")

	// Prints:  Breaking out of 0 iteration; Moved out of loop !!
}

// Output of Program:
// Breaking out of 0 iteration; Breaking out of 1 iteration; Breaking out of 2 iteration; Moved out of loop !!
// Breaking out of 0 iteration; Moved out of loop !!
