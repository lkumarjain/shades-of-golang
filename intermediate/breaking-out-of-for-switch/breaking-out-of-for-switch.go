package main

import "fmt"

// The break statement in the switch moves out of switch and not for loop.
// If you want to move out of for loop, one option is return but
// this may not be a solution every time another alternative is to
// use label in break statement (similar to goto statements)
func main() {
	trap()
	solution()
}

func trap() {
	for i := 0; i < 3; i++ {
		switch {
		case true:
			fmt.Printf(" Breaking out of %v iteration;", i)
			break // Breaking out of switch statement and not for loop
		}
	}

	fmt.Println(" Moved out of loop !!")
	// Prints:  Breaking out of 0 iteration; Breaking out of 1 iteration; Breaking out of 2 iteration; Moved out of loop !!
}

func solution() {
loop: // Creating a label
	for i := 0; i < 5; i++ {
		switch {
		case true:
			fmt.Printf(" Breaking out of %v iteration;", i)
			break loop // Breaking on label
		}
	}
	fmt.Println(" Moved out of loop !!")

	// Prints:  Breaking out of 0 iteration; Moved out of loop !!
}

// Output of a program:
// Breaking out of 0 iteration; Breaking out of 1 iteration; Breaking out of 2 iteration; Moved out of loop !!
// Breaking out of 0 iteration; Moved out of loop !!
