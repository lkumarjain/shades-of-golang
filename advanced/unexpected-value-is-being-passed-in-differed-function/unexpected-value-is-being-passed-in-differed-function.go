package main

import "fmt"

// Arguments for a deferred function call are evaluated at the time of
// defer statement is evaluation and not at the time of function execution.
// You should make a use of pointer parameters to overcome this as
// pointer is saved at the time of defer statement is evaluation.
func main() {
	trap()        // prints: Trap result => 40 (not ok if you expected 200)
	alternative() // prints: Alternative result => 200
}

func trap() {
	var i int = 20
	// Copy of i is passed with Value 20
	defer fmt.Println("Trap result =>", func() int { return i * 2 }())
	i = 100 // Does not impact copy of the value passed in defer function
}

func alternative() {
	i := 20
	// Address of i is passed
	defer func(in *int) { fmt.Println("Alternative result =>", *in*2) }(&i)
	i = 100 // Changing this value also changes value in differ
}

// Output:
// Trap result => 40
// Alternative result => 200
