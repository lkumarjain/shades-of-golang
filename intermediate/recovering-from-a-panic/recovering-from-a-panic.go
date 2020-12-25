package main

import "fmt"

// The recover function can be used to catch/intercept a panic.
// Calling recover will do the trick only when it's done in a deferred function.
// The call to recover works only if it's called directly in your deferred function.
// Please comment trap1 and trap2 respactively to validate alternative1 and alternative2
func main() {
	trap1()        // Prints: panic: Trap1 Panic !!
	alternative1() // Prints: recovered: Alternative1 Panic !!

	trap2()        // Prints: panic: Trap2 Panic !!
	alternative2() // Prints: recovered => Alternative2 Panic !!
}

func trap1() {
	recover() // This cal does not do anything a we are not calling this in defer
	panic("Trap1 Panic !!")
}

func alternative1() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	panic("Alternative1 Panic !!")
}

func doRecover() {
	fmt.Println("recovered =>", recover())
}

func trap2() {
	defer func() {
		doRecover() // panic is not recovered, aa call to recover is not direct
	}()

	panic("Trap2 Panic !!")
}

func alternative2() {
	defer doRecover() // panic is recovered
	panic("Alternative2 Panic !!")
}
